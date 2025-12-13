# GitHub Actions Workflows

This directory contains GitHub Actions workflows for the Graphiant SDK Go package.

## Workflows

### `lint.yml` - Code Quality Checks
Comprehensive linting workflow that runs multiple code quality checks in parallel:
- **golangci-lint** - Comprehensive Go linter (runs multiple linters)
- **go vet** - Go static analysis tool

### `test.yml` - Testing
Runs all tests for the SDK:
- **`test` job** - Matrix job testing against Go 1.21, 1.22, 1.23:
  - Unit tests with `go test`
  - Race condition detection (`-race`)
  - Code coverage reporting
  - Coverage reports uploaded to Codecov (Go 1.23 only)
  - **Credential Support**: Optionally reads `GRAPHIANT_HOST`, `GRAPHIANT_USERNAME`, and `GRAPHIANT_PASSWORD` from GitHub secrets/variables
  - Tests that require credentials will automatically skip if credentials are not configured

### `build.yml` - Build Module
Builds and verifies the Go module:
- Verifies module integrity (`go mod verify`)
- Ensures dependencies are tidy (`go mod tidy`)
- Builds all packages (`go build ./...`)
- Verifies build succeeds

### `release.yml` - Release and Tag
Creates releases for the Go module:
- Manual trigger only (workflow_dispatch)
- Restricted to repository admins and maintainers only
- Creates git tags for releases
- Creates GitHub releases
- **Note**: Go modules are published via git tags. Once a tag is pushed, the module is automatically available via `go get`

## Setup

### Required Secrets

No secrets are required for the workflows. Go modules are published via git tags, which are automatically available through the Go module proxy.

### Optional Secrets/Variables

The `test.yml` workflow supports optional credentials for running integration tests:

- **`GRAPHIANT_HOST`** (optional) - Graphiant API host (e.g., `https://api.graphiant.com`)
- **`GRAPHIANT_USERNAME`** (optional) - Graphiant API username
- **`GRAPHIANT_PASSWORD`** (optional) - Graphiant API password

**Configuration:**
- Set as **secrets** (recommended for sensitive data) or **variables** (for non-sensitive defaults)
- Secrets take precedence over variables
- If not configured, tests that require credentials will automatically skip
- Configure in: Repository Settings → Secrets and variables → Actions

**Note:** The workflow reads from both `secrets` and `vars`, with `secrets` taking precedence. This allows you to set default values in variables while keeping sensitive credentials in secrets.

### Workflow Triggers

- **Pull Requests**: All workflows run on PRs to ensure code quality
- **Push to main/develop**: CI workflows run on pushes to main branches
- **Scheduled**: Test workflow runs nightly at 2 AM UTC
- **Manual**: Release workflow must be manually triggered via workflow_dispatch (restricted to admins and maintainers)

## Usage

### Running Workflows Locally

While you can't run GitHub Actions locally, you can run the same commands:

```bash
# Linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run

# Static analysis
go vet ./...

# Testing
# Without credentials (tests requiring credentials will skip)
go test -v -race -coverprofile=coverage.out ./...

# With credentials (for integration tests)
export GRAPHIANT_HOST="https://api.graphiant.com"
export GRAPHIANT_USERNAME="your_username"
export GRAPHIANT_PASSWORD="your_password"
go test -v -race -coverprofile=coverage.out ./...

# Building
go mod download
go mod verify
go build ./...
```

### Publishing a Release

1. Ensure your code is ready and all tests pass
2. Manually trigger the release workflow via `workflow_dispatch` from the GitHub Actions tab
3. Provide the release version (e.g., `v1.2.3` or `1.2.3` - the workflow will add 'v' prefix if missing)
4. Choose whether to create a git tag (default: true)
5. The workflow will:
   - Verify the module
   - Build and test
   - Create a git tag
   - Create a GitHub release

**Note**: Once a git tag is created and pushed, the Go module is automatically available via:
```bash
go get github.com/Graphiant-Inc/graphiant-sdk-go@v1.2.3
```

## Workflow Status

You can view workflow status in the **Actions** tab of your GitHub repository.
