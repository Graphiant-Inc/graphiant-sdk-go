#!/usr/bin/env bash
# Regenerate the graphiant-sdk-go from the OpenAPI specification.
#
# Prerequisites:
#   - Java 11+ on PATH
#   - openapi-generator-cli JAR (download from https://openapi-generator.tech/docs/installation)
#     or install via Homebrew: brew install openapi-generator
#
# Usage:
#   bash scripts/generate.sh                   # uses api/openapi.yaml
#   OPENAPI_SPEC=path/to/spec.yaml bash scripts/generate.sh

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OPENAPI_SPEC="${OPENAPI_SPEC:-${REPO_ROOT}/api/openapi.yaml}"
# To use the versioned JSON bundle instead:
#   OPENAPI_SPEC=api/graphiant_api_docs_v26.5.0.json bash scripts/generate.sh
PACKAGE_NAME="${PACKAGE_NAME:-graphiant_sdk}"

# Resolve generator: honour explicit $GENERATOR, then try common install locations.
if [ -n "${GENERATOR:-}" ]; then
  : # use what was passed
elif command -v openapi-generator &>/dev/null; then
  GENERATOR="openapi-generator"          # Homebrew: brew install openapi-generator
elif command -v openapi-generator-cli &>/dev/null; then
  GENERATOR="openapi-generator-cli"      # npm: npm i -g @openapitools/openapi-generator-cli
else
  echo "❌ openapi-generator not found. Install via one of:"
  echo "   brew install openapi-generator          # macOS Homebrew"
  echo "   npm install -g @openapitools/openapi-generator-cli"
  echo "   https://openapi-generator.tech/docs/installation"
  exit 1
fi

if [ ! -f "${OPENAPI_SPEC}" ]; then
  echo "❌ OpenAPI spec not found: ${OPENAPI_SPEC}"
  exit 1
fi

echo "🔄 Generating SDK from ${OPENAPI_SPEC}..."

# Read SDK version from version.go so UserAgent and generated docs stay correct.
SDK_VERSION=$(grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' "${REPO_ROOT}/version.go" | head -1 | tr -d 'v')

"${GENERATOR}" generate \
  --input-spec "${OPENAPI_SPEC}" \
  --generator-name go \
  --output "${REPO_ROOT}" \
  --git-user-id Graphiant-Inc \
  --git-repo-id graphiant-sdk-go \
  --additional-properties="packageName=${PACKAGE_NAME},withGoMod=false,packageVersion=${SDK_VERSION}"

echo "🔧 Tidying module..."
cd "${REPO_ROOT}"
go mod tidy

echo "✅ SDK generation complete."
echo "   Files listed in .openapi-generator-ignore are NOT overwritten, including:"
echo "   - auth_env.go, api_custom.go, version.go  (hand-written SDK helpers)"
echo "   - client.go  (pinned to preserve the callAPI() security fix — CWE-532)"
echo "   Review the diff before committing."
