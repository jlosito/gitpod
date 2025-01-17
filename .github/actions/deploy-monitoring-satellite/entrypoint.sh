#!/usr/bin/env bash

set -euo pipefail

export HOME=/home/gitpod
export PREVIEW_ENV_DEV_SA_KEY_PATH="$HOME/.config/gcloud/preview-environment-dev-sa.json"

echo "${INPUT_SA_KEY}" > "${PREVIEW_ENV_DEV_SA_KEY_PATH}"
gcloud auth activate-service-account --key-file "${PREVIEW_ENV_DEV_SA_KEY_PATH}"

# Hack alert: We're building previewctl here until we decide how to properly distribute internal tools
# Also, LEEWAY_WORKSPACE_ROOT is set to /workspace/gitpod in our dev image, but that's not the path GH actions use
# shellcheck disable=SC2155
export LEEWAY_WORKSPACE_ROOT="$(pwd)"
leeway run dev/preview/previewctl:install --dont-test

echo "previewctl get-credentials"
previewctl get-credentials --gcp-service-account "${PREVIEW_ENV_DEV_SA_KEY_PATH}"

echo "previewctl install-context"
previewctl install-context --log-level debug --timeout 10m --gcp-service-account "${PREVIEW_ENV_DEV_SA_KEY_PATH}"

echo "leeway run dev/preview:deploy-monitoring-satellite"
leeway run dev/preview:deploy-monitoring-satellite

{
    echo '<p>Monitoring satellite has been installed in your preview environment.</p>'
    echo '<ul>'
    echo '<li><b>📚 Documentation</b> - See our <a href="https://www.notion.so/gitpod/f2938b2bcb0c4c8c99afe1d2b872380e" target="_blank">internal documentation</a> on how to use it.</li>'
    echo '</ul>'
} >> "${GITHUB_STEP_SUMMARY}"
