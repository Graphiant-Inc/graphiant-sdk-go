# Security Policy

## Supported Versions

We actively support the following versions with security updates:

| Version | Supported          | Notes                    |
| ------- | ------------------ | ------------------------ |
| 25.11.x | :white_check_mark: | Current stable release    |
| 25.10.x | :white_check_mark: | Previous release         |
| < 25.10 | :x:                | No longer supported      |

**Note:** We recommend always using the latest version to ensure you have the most recent security patches.

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability, please follow these steps:

### How to Report

1. **Do NOT** open a public GitHub issue for security vulnerabilities
2. Email security details to: **security@graphiant.com**
3. Include the following information:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if available)
   - Your contact information

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days
- **Resolution**: Depends on severity (see below)

### Severity Levels

| Severity | Response Time | Description                                    |
|----------|---------------|------------------------------------------------|
| Critical | 24-48 hours    | Remote code execution, authentication bypass   |
| High     | 7 days         | Privilege escalation, data exposure           |
| Medium   | 30 days        | Information disclosure, denial of service     |
| Low      | 90 days        | Best practice violations, minor issues         |

### What to Expect

- **Acknowledgment**: You will receive an acknowledgment email within 48 hours
- **Updates**: Regular updates on the status of the vulnerability
- **Credit**: With your permission, we will credit you in security advisories
- **Disclosure**: We will coordinate public disclosure after a fix is available

## Security Best Practices

### Credential Management

- **Never commit secrets**: Never commit API keys, tokens, passwords, or other sensitive information to the repository
- **Use environment variables**: Store credentials in environment variables
  ```go
  username := os.Getenv("GRAPHIANT_USERNAME")
  password := os.Getenv("GRAPHIANT_PASSWORD")
  host := os.Getenv("GRAPHIANT_HOST")
  ```
- **Use secure storage**: For production applications, use secure secret management systems
- **Rotate credentials**: Regularly rotate API keys and passwords

### Code Security

- **Input Validation**: Always validate and sanitize user inputs
- **Error Handling**: Don't expose sensitive information in error messages
- **Dependency Management**: Keep dependencies up to date
  ```bash
  go list -u -m all  # Check for updates
  go get -u ./...     # Update dependencies
  ```
- **Security Scanning**: Use tools like `gosec` for security analysis
  ```bash
  go install github.com/securego/gosec/v2/cmd/gosec@latest
  gosec ./...
  ```

### Dependency Management

- **Regular Updates**: Regularly update dependencies to patch security vulnerabilities
- **Vulnerability Scanning**: Use `govulncheck` to scan for known vulnerabilities
  ```bash
  go install golang.org/x/vuln/cmd/govulncheck@latest
  govulncheck ./...
  ```
- **Minimal Dependencies**: Only include necessary dependencies
- **Version Pinning**: Use specific versions in `go.mod` for production

### CI/CD Security

- **GitHub Actions Secrets**: Use GitHub Secrets for sensitive data (never hardcode)
- **Branch Protection**: All changes require review and signed commits
- **Code Scanning**: Automated security scanning in CI/CD pipelines
- **Dependency Scanning**: Automated dependency vulnerability scanning

### Repository Security

- **CODEOWNERS**: Code owners are automatically requested for review
- **Branch Protection**: Main branch is protected with required reviews
- **Signed Commits**: All commits must be verified with GPG signatures
- **Access Control**: Repository access is restricted to authorized team members

### Go-Specific Security

- **Race Conditions**: Use `go test -race` to detect race conditions
- **Memory Safety**: Leverage Go's memory safety features
- **Type Safety**: Use strong typing to prevent type-related vulnerabilities
- **Error Handling**: Always handle errors explicitly

### Environment Variables

When using environment variables for credentials:

```go
package main

import (
    "os"
    "log"
)

func main() {
    username := os.Getenv("GRAPHIANT_USERNAME")
    if username == "" {
        log.Fatal("GRAPHIANT_USERNAME environment variable is required")
    }
    
    password := os.Getenv("GRAPHIANT_PASSWORD")
    if password == "" {
        log.Fatal("GRAPHIANT_PASSWORD environment variable is required")
    }
    
    host := os.Getenv("GRAPHIANT_HOST")
    if host == "" {
        host = "https://portal.graphiant.com" // default
    }
}
```

### Testing Security

- **Test with invalid inputs**: Test error handling and edge cases
- **Test authentication**: Verify authentication and authorization work correctly
- **Review test coverage**: Ensure security-critical paths are tested

### Documentation Security

- **No secrets in examples**: Never include real credentials in documentation or examples
- **Security warnings**: Document security considerations for sensitive operations
- **Best practices**: Include security best practices in documentation

## Security Checklist for Contributors

Before submitting a pull request, ensure:

- [ ] No secrets or credentials are committed
- [ ] Input validation is implemented
- [ ] Error messages don't expose sensitive information
- [ ] Dependencies are up to date
- [ ] Tests cover security-critical paths
- [ ] Code follows Go security best practices
- [ ] No hardcoded credentials or API keys
- [ ] Environment variables are used for configuration

## Additional Resources

- [Go Security Best Practices](https://go.dev/doc/security/best-practices)
- [OWASP Go Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Go_Language_Cheat_Sheet.html)
- [Go Vulnerability Database](https://pkg.go.dev/vuln)

## Contact

For security concerns, please contact: **security@graphiant.com**

---

**Last Updated**: 2025-12-13
