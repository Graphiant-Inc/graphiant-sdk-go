.PHONY: build test vet fmt fmt-check lint vulncheck tidy generate clean help

# Hand-maintained Go files (see scripts/handwritten-go-files.sh). Generated
# files are left as OpenAPI Generator writes them.
HANDWRITTEN_GO = $(shell bash scripts/handwritten-go-files.sh)

OPENAPI_SPEC     ?= api/openapi.yaml
OPENAPI_JAR      ?= $(shell which openapi-generator-cli 2>/dev/null || echo openapi-generator-cli)
PACKAGE_NAME     ?= graphiant_sdk

## build: Compile all packages
build:
	go build ./...

## test: Run tests with race detector and coverage
test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

## vet: Run go vet
vet:
	go vet ./...

## fmt: Format hand-maintained Go files
fmt:
	gofmt -w $(HANDWRITTEN_GO)

## fmt-check: Fail if any hand-maintained Go file is not gofmt-formatted
fmt-check:
	@unformatted="$$(gofmt -l $(HANDWRITTEN_GO))"; \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not gofmt-formatted (run 'make fmt'):"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

## lint: Run golangci-lint (same config as CI; skips generated files via .golangci.yml)
lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest run --timeout=5m ./...

## vulncheck: Report known vulnerabilities in reachable code (Go toolchain and modules)
vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## tidy: Tidy and verify go.mod / go.sum
tidy:
	go mod tidy
	go mod verify

## generate: Regenerate the SDK from the OpenAPI spec
generate:
	@bash scripts/generate.sh

## clean: Remove generated model and doc files (keeps hand-written core files)
clean:
	@echo "Removing generated model files..."
	find . -maxdepth 1 -name "model_*.go" -delete
	@echo "Removing generated api_default.go..."
	rm -f api_default.go
	@echo "Removing generated docs..."
	find docs/ -name "*.md" -delete 2>/dev/null || true
	@echo "Done."

## help: Show this help
help:
	@grep -E '^## ' Makefile | sed 's/## /  /'

.DEFAULT_GOAL := build
