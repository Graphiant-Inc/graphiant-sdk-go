.PHONY: build test lint tidy generate clean help

OPENAPI_SPEC     ?= api/openapi.yaml
OPENAPI_JAR      ?= $(shell which openapi-generator-cli 2>/dev/null || echo openapi-generator-cli)
PACKAGE_NAME     ?= graphiant_sdk

## build: Compile all packages
build:
	go build ./...

## test: Run tests with race detector and coverage
test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

## lint: Run golangci-lint (skips generated files via .golangci.yml)
lint:
	golangci-lint run --timeout=5m

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
