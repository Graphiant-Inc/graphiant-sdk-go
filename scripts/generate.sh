#!/usr/bin/env bash
# Regenerate the graphiant-sdk-go from the OpenAPI specification.
#
# Prerequisites:
#   - Java 11+ on PATH
#   - OpenAPI Generator >= 7.23.0 (Homebrew: brew install openapi-generator)
#
# Usage:
#   bash scripts/generate.sh                         # auto-selects JSON bundle, then YAML
#   OPENAPI_SPEC=api/openapi.yaml bash scripts/generate.sh
#
# Note: api/openapi.yaml exceeds SnakeYAML's default 3 MB code-point limit.
# Use the JSON bundle (api/graphiant_api_docs_*.json) as the default input, or
# pass JAVA_OPTS to raise the limit when you explicitly want the YAML:
#   JAVA_OPTS="-Dsnakeyaml.codepoints.max.all=99999999" \
#   OPENAPI_SPEC=api/openapi.yaml bash scripts/generate.sh

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

# Resolve default spec: prefer versioned JSON bundle (no SnakeYAML size limit);
# fall back to api/openapi.yaml only when no JSON bundle is present.
if [ -z "${OPENAPI_SPEC:-}" ]; then
  JSON_BUNDLE="$(ls "${REPO_ROOT}"/api/graphiant_api_docs_*.json 2>/dev/null | sort | tail -1)"
  if [ -n "${JSON_BUNDLE}" ]; then
    OPENAPI_SPEC="${JSON_BUNDLE}"
  elif [ -f "${REPO_ROOT}/api/openapi.yaml" ]; then
    OPENAPI_SPEC="${REPO_ROOT}/api/openapi.yaml"
    # Raise SnakeYAML limit for large YAML specs.
    export JAVA_OPTS="${JAVA_OPTS:+${JAVA_OPTS} }-Dsnakeyaml.codepoints.max.all=99999999"
  fi
fi

# If caller passed a YAML spec explicitly, ensure the size limit is raised.
if [[ "${OPENAPI_SPEC:-}" == *.yaml || "${OPENAPI_SPEC:-}" == *.yml ]]; then
  export JAVA_OPTS="${JAVA_OPTS:+${JAVA_OPTS} }-Dsnakeyaml.codepoints.max.all=99999999"
fi

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

MIN_OPENAPI_GENERATOR_VERSION="${MIN_OPENAPI_GENERATOR_VERSION:-7.23.0}"
GENERATOR_VERSION_RAW="$("${GENERATOR}" version 2>/dev/null || true)"
GENERATOR_VERSION="$(printf '%s' "${GENERATOR_VERSION_RAW}" | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1 || true)"

if [ -z "${GENERATOR_VERSION}" ]; then
  echo "⚠️  Unable to parse ${GENERATOR} version from: ${GENERATOR_VERSION_RAW}"
  echo "   Expected OpenAPI Generator >= ${MIN_OPENAPI_GENERATOR_VERSION}."
  echo "   If generation fails unexpectedly, upgrade your generator and retry."
elif [ "$(printf '%s\n%s\n' "${MIN_OPENAPI_GENERATOR_VERSION}" "${GENERATOR_VERSION}" | sort -V | head -1)" != "${MIN_OPENAPI_GENERATOR_VERSION}" ]; then
  echo "❌ OpenAPI Generator ${GENERATOR_VERSION} is too old."
  echo "   Required version: >= ${MIN_OPENAPI_GENERATOR_VERSION}"
  echo "   Upgrade example: brew upgrade openapi-generator"
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
