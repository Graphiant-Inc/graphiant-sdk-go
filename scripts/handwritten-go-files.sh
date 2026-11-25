#!/usr/bin/env bash
# List hand-maintained Go files: those without a "Code generated" header, plus
# files pinned in .openapi-generator-ignore (e.g. client.go, configuration.go),
# which carry the generator header but are maintained by hand.
#
# Generated files are left exactly as OpenAPI Generator writes them, so
# `make fmt` / `make fmt-check` only touch the files listed here.
set -euo pipefail

cd "$(dirname "$0")/.."

git ls-files '*.go' | while read -r f; do
  if ! head -12 "$f" | grep -q 'Code generated' || grep -qxF "$f" .openapi-generator-ignore; then
    echo "$f"
  fi
done
