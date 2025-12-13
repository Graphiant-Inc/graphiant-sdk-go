# Contributing to Graphiant SDK Go

Thank you for your interest in contributing!

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork:**
   ```bash
   git clone https://github.com/Graphiant-Inc/graphiant-sdk-go.git
   cd graphiant-sdk-go
   ```
3. **Set up development environment:**
   ```bash
   # Ensure Go 1.21+ is installed
   go version
   
   # Download dependencies
   go mod download
   ```

## Development Workflow

1. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** and ensure they pass local checks:
   ```bash
   # Format code
   gofmt -s -w .
   
   # Run linting
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   golangci-lint run
   
   # Run static analysis
   go vet ./...
   
   # Run tests
   go test -v -race -coverprofile=coverage.out ./...
   ```

3. **Verify module integrity:**
   ```bash
   go mod verify
   go mod tidy
   ```

4. **Build the project:**
   ```bash
   go build ./...
   ```

5. **Commit with clear messages:**
   ```bash
   git commit -m "Add: description of changes"
   ```
   
   **Note**: All commits must be signed with GPG. See [Branch Protection Requirements](#branch-protection-requirements) below.

6. **Push and create a pull request**

## Linting Tools

The project uses multiple linting tools to ensure code quality:

| Tool | Purpose | Target | CI/CD |
|------|---------|--------|-------|
| `golangci-lint` | Comprehensive Go linter | All `.go` files | Yes (lint stage) |
| `gofmt` | Go code formatting | All `.go` files | Yes (lint stage) |
| `go vet` | Go static analysis | All `.go` files | Yes (lint stage) |

**Note:** All linting tools run automatically in CI/CD on every pull request and push to main/develop branches.

## Testing

### Running Tests Locally

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detection
go test -race ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Structure

- `test/` directory contains all test files
- Tests use the `testify` framework for assertions
- Tests are automatically run in CI/CD across Go 1.21, 1.22, and 1.23

## Code Standards

### Go Code
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` for formatting (enforced in CI)
- Follow Go naming conventions
- Include godoc comments for exported functions and types
- Keep functions focused and small
- Handle errors explicitly (don't ignore errors)

### Example Code Structure

```go
// Package graphiant_sdk provides a Go SDK for Graphiant APIs.
package graphiant_sdk

import (
    "context"
    "fmt"
)

// Client represents a Graphiant API client.
type Client struct {
    // fields
}

// NewClient creates a new Graphiant API client.
// It returns an error if the configuration is invalid.
func NewClient(config *Configuration) (*Client, error) {
    // implementation
}
```

## Pull Request Checklist

- [ ] Code follows Go style guidelines
- [ ] Code is formatted with `gofmt`
- [ ] All tests pass locally
- [ ] Linting passes (`golangci-lint`, `go vet`)
- [ ] Module is verified (`go mod verify`)
- [ ] Dependencies are tidy (`go mod tidy`)
- [ ] Commit messages are clear
- [ ] Commits are signed with GPG (required)
- [ ] Branch is rebased (no merge commits allowed)
- [ ] All CI/CD checks pass (lint, test, build)

## Branch Protection Requirements

This repository has branch protection rules that must be satisfied before a pull request can be merged:

### Required Approvals
- **SRE Team Approval**: All pull requests require approval from `@Graphiant-Inc/sre`
- **Code Owners**: Additional approvals may be required based on CODEOWNERS file

### Merge Requirements
- **Merge Method**: Only **squash merge** or **rebase merge** are allowed (standard merge is disabled)
- **No Merge Commits**: Your branch must not contain merge commits
  - Use `git rebase` instead of `git merge` when updating your branch
  - Example: `git pull --rebase origin main` or `git rebase origin/main`

### Commit Requirements
- **Signed Commits**: All commits must be verified with GPG signatures
  - Set up GPG signing: https://docs.github.com/en/authentication/managing-commit-signature-verification
  - Configure Git: `git config --global commit.gpgsign true`
  - Verify your commits are signed: `git log --show-signature`

### CI/CD Requirements
- **All Checks Must Pass**: All lint, test, and build workflows must pass
- **No Skipped Tests**: All tests must run and pass

## Troubleshooting

### GPG Signing Issues

If you encounter "gpg failed to sign the data":
1. Ensure GPG is installed and configured
2. Set `GPG_TTY`: `export GPG_TTY=$(tty)`
3. Verify your key: `gpg --list-secret-keys --keyid-format LONG`
4. Configure Git: `git config --global user.signingkey YOUR_KEY_ID`

If you encounter "The email in this signature doesn't match the committer email":
1. Ensure your commit author email matches your GPG key email
2. Amend the commit: `GIT_COMMITTER_EMAIL="your-email@example.com" git commit --amend --author="Your Name <your-email@example.com>" --no-edit -S`

### Merge Conflicts

If you have merge conflicts:
1. Fetch latest: `git fetch origin main`
2. Rebase your branch: `git rebase origin/main`
3. Resolve conflicts and continue: `git rebase --continue`
4. Force push: `git push --force-with-lease origin your-branch`

## Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [GitHub Actions Workflows](.github/workflows/README.md)
