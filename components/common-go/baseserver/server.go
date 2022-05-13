// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package baseserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	common_grpc "github.com/gitpod-io/gitpod/common-go/grpc"
	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/common-go/pprof"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func New(name string, opts ...Option) (*Server, error) {
	options, err := evaluateOptions(defaultOptions(), opts...)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	server := &Server{
		Name:    name,
		options: options,
	}

	err = server.initializeDebug()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize debug server: %w", err)
	}

	server.httpMux = http.NewServeMux()
	server.http = &http.Server{Handler: server.httpMux}

	err = server.initializeGRPC()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gRPC server: %w", err)
	}

	return server, nil
}

// Server is a packaged server with batteries included. It is designed to be standard across components where it makes sense.
// Server implements graceful shutdown making it suitable for usage in integration tests. See server_test.go.
//
// Server is composed of the following:
// 	* Debug server which serves observability and debug endpoints
//		- /metrics for Prometheus metrics
//		- /pprof for Golang profiler
//		- /ready for kubernetes readiness check
//		- /live for kubernetes liveness check
//	* (optional) gRPC server with standard interceptors and configuration
//		- Started when baseserver is configured WithGRPCPort (port is non-negative)
//		- Use Server.GRPC() to get access to the underlying grpc.Server and register services
//	* (optional) HTTP server
//		- Currently does not come with any standard HTTP middlewares
//		- Started when baseserver is configured WithHTTPPort (port is non-negative)
// 		- Use Server.HTTPMux() to get access to the root handler and register your endpoints
type Server struct {
	// Name is the name of this server, used for logging context
	Name string

	options *options

	// debug is an HTTP server for debug endpoints - metrics, pprof, readiness & liveness.
	debug         *http.Server
	debugListener net.Listener

	// http is an http Server, only used when port is specified in cfg
	http         *http.Server
	httpMux      *http.ServeMux
	httpListener net.Listener

	// grpc is a grpc Server, only used when port is specified in cfg
	grpc         *grpc.Server
	grpcListener net.Listener

	// listening indicates the server is serving. When closed, the server is in the process of graceful termination.
	listening chan struct{}
	closeOnce sync.Once
}

func serveHTTP(cfg *ServerConfiguration, srv *http.Server, l net.Listener) (err error) {
	if cfg.TLS == nil {
		err = srv.Serve(l)
	} else {
		err = srv.ServeTLS(l, cfg.TLS.Cert, cfg.TLS.Key)
	}
	return
}

func (s *Server) ListenAndServe() error {
	var err error

	s.listening = make(chan struct{})
	defer func() {
		err := s.Close()
		if err != nil {
			s.Logger().WithError(err).Errorf("cannot close gracefully")
		}
	}()

	if srv := s.options.config.Services.Debug; srv != nil {
		s.debugListener, err = net.Listen("tcp", srv.Address)
		if err != nil {
			return fmt.Errorf("failed to start debug server: %w", err)
		}
		s.debug.Addr = srv.Address

		go func() {
			err := serveHTTP(srv, s.debug, s.debugListener)
			if err != nil {
				s.Logger().WithError(err).Errorf("debug server encountered an error - closing remaining servers.")
				s.Close()
			}
		}()
	}

	if srv := s.options.config.Services.HTTP; srv != nil {
		s.httpListener, err = net.Listen("tcp", srv.Address)
		if err != nil {
			return fmt.Errorf("failed to start HTTP server: %w", err)
		}
		s.http.Addr = srv.Address

		go func() {
			err := serveHTTP(srv, s.http, s.httpListener)
			if err != nil {
				s.Logger().WithError(err).Errorf("HTTP server encountered an error - closing remaining servers.")
				s.Close()
			}
		}()
	}

	if srv := s.options.config.Services.GRPC; srv != nil {
		s.grpcListener, err = net.Listen("tcp", srv.Address)
		if err != nil {
			return fmt.Errorf("failed to start gRPC server: %w", err)
		}

		go func() {
			err := s.grpc.Serve(s.grpcListener)
			if err != nil {
				s.Logger().WithError(err).Errorf("gRPC server encountered an error - closing remaining servers.")
				s.Close()
			}
		}()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Await operating system signals, or server errors.
	select {
	case sig := <-signals:
		s.Logger().Infof("Received system signal %s, closing server.", sig.String())
		return nil
	}
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.options.closeTimeout)
	defer cancel()

	var err error
	s.closeOnce.Do(func() {
		err = s.close(ctx)
	})
	return err
}

func (s *Server) Logger() *logrus.Entry {
	return s.options.logger
}

func (s *Server) HTTPMux() *http.ServeMux {
	return s.httpMux
}

func (s *Server) GRPC() *grpc.Server {
	return s.grpc
}

func (s *Server) MetricsRegistry() *prometheus.Registry {
	return s.options.metricsRegistry
}

func (s *Server) close(ctx context.Context) error {
	if s.listening == nil {
		return fmt.Errorf("server is not running, invalid close operation")
	}

	if s.isClosing() {
		s.Logger().Debug("Server is already closing.")
		return nil
	}

	s.Logger().Info("Received graceful shutdown request.")
	close(s.listening)

	if s.grpc != nil {
		s.grpc.GracefulStop()
		// s.grpc.GracefulStop() also closes the underlying net.Listener, we just release the reference.
		s.grpcListener = nil
		s.Logger().Info("GRPC server terminated.")
	}

	if s.http != nil {
		err := s.http.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("failed to close http server: %w", err)
		}
		// s.http.Shutdown() also closes the underlying net.Listener, we just release the reference.
		s.httpListener = nil
		s.Logger().Info("HTTP server terminated.")
	}

	// Always terminate debug server last, we want to keep it running for as long as possible
	if s.debug != nil {
		err := s.debug.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("failed to close debug server: %w", err)
		}
		// s.http.Shutdown() also closes the underlying net.Listener, we just release the reference.
		s.debugListener = nil
		s.Logger().Info("Debug server terminated.")
	}

	return nil
}

func (s *Server) isClosing() bool {
	select {
	case <-s.listening:
		// listening channel is closed, we're in graceful shutdown mode
		return true
	default:
		return false
	}
}

func (s *Server) initializeDebug() error {
	logger := s.Logger().WithField("protocol", "debug")

	mux := http.NewServeMux()

	mux.HandleFunc("/ready", s.options.healthHandler.ReadyEndpoint)
	logger.Debug("Serving readiness handler on /ready")

	mux.HandleFunc("/live", s.options.healthHandler.LiveEndpoint)
	logger.Debug("Serving liveliness handler on /live")

	mux.Handle("/metrics", promhttp.InstrumentMetricHandler(
		s.options.metricsRegistry, promhttp.HandlerFor(s.options.metricsRegistry, promhttp.HandlerOpts{}),
	))
	s.Logger().WithField("protocol", "http").Debug("Serving metrics on /metrics")

	mux.Handle(pprof.Path, pprof.Handler())
	logger.Debug("Serving profiler on /debug/pprof")

	s.debug = &http.Server{
		Handler: mux,
	}

	return nil
}

func (s *Server) initializeGRPC() error {
	common_grpc.SetupLogging()

	grpcMetrics := grpc_prometheus.NewServerMetrics()
	grpcMetrics.EnableHandlingTimeHistogram()
	if err := s.MetricsRegistry().Register(grpcMetrics); err != nil {
		return fmt.Errorf("failed to register grpc metrics: %w", err)
	}

	unary := []grpc.UnaryServerInterceptor{
		grpc_logrus.UnaryServerInterceptor(s.Logger()),
		grpcMetrics.UnaryServerInterceptor(),
	}
	stream := []grpc.StreamServerInterceptor{
		grpc_logrus.StreamServerInterceptor(s.Logger()),
		grpcMetrics.StreamServerInterceptor(),
	}

	opts := common_grpc.ServerOptionsWithInterceptors(stream, unary)
	if cfg := s.options.config.Services.GRPC; cfg != nil && cfg.TLS != nil {
		tlsConfig, err := common_grpc.ClientAuthTLSConfig(
			cfg.TLS.CA, cfg.TLS.Cert, cfg.TLS.Key,
			common_grpc.WithSetClientCAs(true),
			common_grpc.WithServerName(s.Name),
		)
		if err != nil {
			log.WithError(err).Fatal("cannot load ws-manager certs")
		}

		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	s.grpc = grpc.NewServer(opts...)

	// Register health service by default
	grpc_health_v1.RegisterHealthServer(s.grpc, s.options.grpcHealthCheck)

	return nil
}

func httpAddress(cfg *ServerConfiguration, l net.Listener) string {
	if l == nil {
		return ""
	}
	protocol := "http"
	if cfg != nil && cfg.TLS != nil {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s", protocol, l.Addr().String())
}

func (s *Server) DebugAddress() string {
	return httpAddress(s.options.config.Services.Debug, s.debugListener)
}
func (s *Server) HTTPAddress() string {
	return httpAddress(s.options.config.Services.HTTP, s.httpListener)
}
func (s *Server) ReadinessAddress() string {
	return s.DebugAddress()
}
func (s *Server) GRPCAddress() string { return s.options.config.Services.GRPC.GetAddress() }
