// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: usage/v1/billing.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Sessions  []*BilledSession       `protobuf:"bytes,3,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *UpdateInvoicesRequest) Reset() {
	*x = UpdateInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoicesRequest) ProtoMessage() {}

func (x *UpdateInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoicesRequest.ProtoReflect.Descriptor instead.
func (*UpdateInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateInvoicesRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *UpdateInvoicesRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *UpdateInvoicesRequest) GetSessions() []*BilledSession {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type UpdateInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInvoicesResponse) Reset() {
	*x = UpdateInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usage_v1_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoicesResponse) ProtoMessage() {}

func (x *UpdateInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usage_v1_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoicesResponse.ProtoReflect.Descriptor instead.
func (*UpdateInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_usage_v1_billing_proto_rawDescGZIP(), []int{1}
}

var File_usage_v1_billing_proto protoreflect.FileDescriptor

var file_usage_v1_billing_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x67, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_usage_v1_billing_proto_rawDescOnce sync.Once
	file_usage_v1_billing_proto_rawDescData = file_usage_v1_billing_proto_rawDesc
)

func file_usage_v1_billing_proto_rawDescGZIP() []byte {
	file_usage_v1_billing_proto_rawDescOnce.Do(func() {
		file_usage_v1_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_usage_v1_billing_proto_rawDescData)
	})
	return file_usage_v1_billing_proto_rawDescData
}

var file_usage_v1_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_usage_v1_billing_proto_goTypes = []interface{}{
	(*UpdateInvoicesRequest)(nil),  // 0: usage.v1.UpdateInvoicesRequest
	(*UpdateInvoicesResponse)(nil), // 1: usage.v1.UpdateInvoicesResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*BilledSession)(nil),          // 3: usage.v1.BilledSession
}
var file_usage_v1_billing_proto_depIdxs = []int32{
	2, // 0: usage.v1.UpdateInvoicesRequest.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: usage.v1.UpdateInvoicesRequest.end_time:type_name -> google.protobuf.Timestamp
	3, // 2: usage.v1.UpdateInvoicesRequest.sessions:type_name -> usage.v1.BilledSession
	0, // 3: usage.v1.BillingService.UpdateInvoices:input_type -> usage.v1.UpdateInvoicesRequest
	1, // 4: usage.v1.BillingService.UpdateInvoices:output_type -> usage.v1.UpdateInvoicesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_usage_v1_billing_proto_init() }
func file_usage_v1_billing_proto_init() {
	if File_usage_v1_billing_proto != nil {
		return
	}
	file_usage_v1_usage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_usage_v1_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInvoicesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_usage_v1_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInvoicesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_usage_v1_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usage_v1_billing_proto_goTypes,
		DependencyIndexes: file_usage_v1_billing_proto_depIdxs,
		MessageInfos:      file_usage_v1_billing_proto_msgTypes,
	}.Build()
	File_usage_v1_billing_proto = out.File
	file_usage_v1_billing_proto_rawDesc = nil
	file_usage_v1_billing_proto_goTypes = nil
	file_usage_v1_billing_proto_depIdxs = nil
}