# Graphiant SDK Go

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/Graphiant-Inc/graphiant-sdk-go.svg)](https://pkg.go.dev/github.com/Graphiant-Inc/graphiant-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Graphiant-Inc/graphiant-sdk-go)](https://goreportcard.com/report/github.com/Graphiant-Inc/graphiant-sdk-go)
[![Documentation](https://img.shields.io/badge/docs-latest-brightgreen.svg)](https://docs.graphiant.com/docs/graphiant-sdk-go)
[![CI/CD](https://github.com/Graphiant-Inc/graphiant-sdk-go/actions/workflows/test.yml/badge.svg)](https://github.com/Graphiant-Inc/graphiant-sdk-go/actions)

A Go client for [Graphiant Network-as-a-Service (NaaS)](https://www.graphiant.com), generated from the Graphiant OpenAPI specification.

More product context: [Graphiant Docs](https://docs.graphiant.com).

## Table of contents

| Section | Contents |
|--------|----------|
| [Documentation & links](#documentation--links) | Official guides, API reference, pkg.go.dev |
| [Features](#features) | What the SDK provides |
| [Quick start](#quick-start) | Install, password login, and **`GRAPHIANT_ACCESS_TOKEN`** |
| [Advanced usage](#advanced-usage) | Patterns and error handling |
| [Convenient wrapper functions](#convenient-wrapper-functions) | `api_custom.go` helpers |
| [Development](#development) | Build, test, code generation |
| [API reference (overview)](#api-reference) | Core types and common endpoints |
| [Security](#security) | Auth and environment variables |
| [Contributing](#contributing) | PR workflow |
| [License](#license) | MIT |
| [Version history](#version-history) | Changelog |
| [Support](#support) | Links and contact |
| [Related projects](#related-projects) | Other SDKs |

## Documentation & links

| Resource | Link |
|----------|------|
| **SDK guide** | [Graphiant SDK Go](https://docs.graphiant.com/docs/graphiant-sdk-go) |
| **Automation** | [Graphiant Automation](https://docs.graphiant.com/docs/automation) |
| **REST API** | [Graphiant Portal REST API](https://docs.graphiant.com/docs/graphiant-portal-rest-api) |
| **Method index (repo)** | [DefaultAPI.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/docs/DefaultAPI.md) |
| **Package** | [pkg.go.dev](https://pkg.go.dev/github.com/Graphiant-Inc/graphiant-sdk-go) |
| **Changelog** | [CHANGELOG.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/CHANGELOG.md) |

## Features

- **Full REST coverage** — Generated client for Graphiant API endpoints.
- **Typed models** — OpenAPI-generated structs and accessors.
- **Bearer auth** — Username/password login or a token you supply.
- **Environment variables** — Same **`GRAPHIANT_ACCESS_TOKEN`** / **`GRAPHIANT_USERNAME`** + **`GRAPHIANT_PASSWORD`** as [graphiant-sdk-python](https://github.com/Graphiant-Inc/graphiant-sdk-python): prefer token from env, else **`POST /v1/auth/login`** via **`AuthorizationBearerFromEnvOrLogin`** (also **`AuthorizationBearerFromEnv`**, **`AccessTokenFromEnv`**, **`ConfigureHostFromEnv`**) in `auth_env.go`.
- **Helpers** — Optional wrappers in `api_custom.go` (e.g. device config polling).

## Quick start

### 1. Install

```bash
go get github.com/Graphiant-Inc/graphiant-sdk-go
```

### 2. Basic usage

The default API base URL is `https://api.graphiant.com`. Use that host (not the portal URL) for `Configuration.Host`.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Graphiant-Inc/graphiant-sdk-go"
)

func main() {
	config := graphiant_sdk.NewConfiguration()
	config.Host = "https://api.graphiant.com" // or your tenant API host
	
	// Create API client
	client := graphiant_sdk.NewAPIClient(config)
	
	// Authentication request
	authReq := graphiant_sdk.NewV1AuthLoginPostRequestWithDefaults()
	authReq.SetUsername("your-username")
	authReq.SetPassword("your-password")
	
	// Get authentication token
	_, httpRes, err := client.DefaultAPI.
		V1AuthLoginPost(context.Background()).
		V1AuthLoginPostRequest(*authReq).
		Execute()
	
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	defer httpRes.Body.Close()
	
	// Parse authentication response
	var authResult struct {
		Auth        bool   `json:"auth"`
		AccountType string `json:"accountType"`
		Token       string `json:"token"`
	}
	
	if err := json.NewDecoder(httpRes.Body).Decode(&authResult); err != nil {
		log.Fatalf("Failed to decode auth response: %v", err)
	}
	
	if !authResult.Auth {
		log.Fatal("Authentication failed")
	}
	
	bearerToken := "Bearer " + authResult.Token
	
	// Get edge devices summary
	resp, _, err := client.DefaultAPI.
		V1EdgesSummaryGet(context.Background()).
		Authorization(bearerToken).
		Execute()
	
	if err != nil {
		log.Fatalf("Failed to get edge summary: %v", err)
	}
	
	// Print edge devices
	edges := resp.GetEdgesSummary()
	fmt.Printf("Found %d edge devices:\n", len(edges))
	for _, edge := range edges {
		fmt.Printf("- Device ID: %d, Hostname: %s, Status: %s\n",
			edge.GetDeviceId(), edge.GetHostname(), edge.GetStatus())
	}
}
```

### 3. Token or username/password (same pattern as Python SDK)

1. If **`GRAPHIANT_ACCESS_TOKEN`** is set (e.g. after **`graphiant login`** and **`source ~/.graphiant/env.sh`**), that raw JWT is used—no login call.
2. If the token is **not** set, **`AuthorizationBearerFromEnvOrLogin`** uses **`GRAPHIANT_USERNAME`** and **`GRAPHIANT_PASSWORD`** with **`POST /v1/auth/login`**.

Optionally set **`GRAPHIANT_API_HOST`** (or **`GRAPHIANT_HOST`**) and call **`ConfigureHostFromEnv(config)`** after **`NewConfiguration()`**.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Graphiant-Inc/graphiant-sdk-go"
)

func main() {
	cfg := graphiant_sdk.NewConfiguration()
	graphiant_sdk.ConfigureHostFromEnv(cfg)
	client := graphiant_sdk.NewAPIClient(cfg)

	authz, err := graphiant_sdk.AuthorizationBearerFromEnvOrLogin(context.Background(), client)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}

	resp, _, err := client.DefaultAPI.
		V1EdgesSummaryGet(context.Background()).
		Authorization(authz).
		Execute()
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	edges := resp.GetEdgesSummary()
	fmt.Fprintf(os.Stdout, "devices: %d\n", len(edges))
}
```

```bash
# Option A — bearer only (same as Python)
export GRAPHIANT_ACCESS_TOKEN="your_jwt"

# Option B — password login when token is unset
export GRAPHIANT_USERNAME="your_username"
export GRAPHIANT_PASSWORD="your_password"
```

Do **not** set a second bearer via context API keys when using `Authorization(...)` on the request, or some gateways may reject duplicate `Authorization` headers.

## 🔧 Advanced Usage

### Device Configuration Management

```go
// Verify device portal status before configuration
func verifyDevicePortalStatus(client *graphiant_sdk.APIClient, bearerToken string, deviceID int64) error {
	resp, _, err := client.DefaultAPI.
		V1EdgesSummaryGet(context.Background()).
		Authorization(bearerToken).
		Execute()
	
	if err != nil {
		return fmt.Errorf("failed to get edges summary: %v", err)
	}
	
	for _, edge := range resp.GetEdgesSummary() {
		if edge.GetDeviceId() == deviceID {
			if edge.GetPortalStatus() == "Ready" && edge.GetTtConnCount() == 2 {
				return nil
			} else {
				return fmt.Errorf("device %d not ready. Status: %s, TT Connections: %d",
					deviceID, edge.GetPortalStatus(), edge.GetTtConnCount())
			}
		}
	}
	return fmt.Errorf("device %d not found", deviceID)
}

// Configure device interfaces
func configureDeviceInterfaces(client *graphiant_sdk.APIClient, bearerToken string, deviceID int64) error {
	// Define circuits
	circuits := map[string]graphiant_sdk.ManaV2CircuitConfig{
		"c-gigabitethernet5-0-0": {
			Name:              graphiant_sdk.PtrString("c-gigabitethernet5-0-0"),
			Description:       graphiant_sdk.PtrString("c-gigabitethernet5-0-0"),
			LinkUpSpeedMbps:   graphiant_sdk.PtrInt32(50),
			LinkDownSpeedMbps: graphiant_sdk.PtrInt32(100),
			ConnectionType:    graphiant_sdk.PtrString("internet_dia"),
			Label:             graphiant_sdk.PtrString("internet_dia_4"),
			QosProfile:        graphiant_sdk.PtrString("gold25"),
			QosProfileType:    graphiant_sdk.PtrString("balanced"),
			DiaEnabled:        graphiant_sdk.PtrBool(false),
			LastResort:        graphiant_sdk.PtrBool(false),
		},
	}
	
	// Define interfaces
	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		"GigabitEthernet5/0/0": {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				AdminStatus:       graphiant_sdk.PtrBool(true),
				MaxTransmissionUnit: graphiant_sdk.PtrInt32(1500),
				Circuit:           graphiant_sdk.PtrString("c-gigabitethernet5-0-0"),
				Description:       graphiant_sdk.PtrString("wan_1"),
				Alias:             graphiant_sdk.PtrString("primary_wan"),
				Ipv4: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Dhcp: &graphiant_sdk.ManaV2InterfaceDhcpConfig{
						DhcpClient: graphiant_sdk.PtrBool(true),
					},
				},
				Ipv6: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Dhcp: &graphiant_sdk.ManaV2InterfaceDhcpConfig{
						DhcpClient: graphiant_sdk.PtrBool(true),
					},
				},
			},
		},
	}
	
	// Create configuration request
	configRequest := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configRequest.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configRequest.Edge.SetCircuits(circuits)
	configRequest.Edge.SetInterfaces(interfaces)
	
	// Verify device is ready
	if err := verifyDevicePortalStatus(client, bearerToken, deviceID); err != nil {
		return fmt.Errorf("device not ready: %v", err)
	}
	
	// Push configuration
	_, _, err := client.DefaultAPI.
		V1DevicesDeviceIdConfigPut(context.Background(), deviceID).
		Authorization(bearerToken).
		V1DevicesDeviceIdConfigPutRequest(*configRequest).
		Execute()
	
	if err != nil {
		return fmt.Errorf("configuration failed: %v", err)
	}
	
	fmt.Printf("Configuration job submitted for device %d\n", deviceID)
	return nil
}
```

### BGP Monitoring

```go
func getBgpMonitoring(client *graphiant_sdk.APIClient, bearerToken string) error {
	// Create BGP monitoring request
	bgpReq := graphiant_sdk.NewV2MonitoringBgpPostRequest()
	
	// Configure selectors
	selectors := []graphiant_sdk.V2MonitoringBgpPostRequestSelectorsInner{
		{
			Type:  graphiant_sdk.PtrString("enterprise"),
			Value: graphiant_sdk.PtrString("your-enterprise-id"),
		},
	}
	bgpReq.SetSelectors(selectors)
	
	// Get BGP monitoring data
	bgpResp, _, err := client.DefaultAPI.
		V2MonitoringBgpPost(context.Background()).
		Authorization(bearerToken).
		V2MonitoringBgpPostRequest(*bgpReq).
		Execute()
	
	if err != nil {
		return fmt.Errorf("BGP monitoring failed: %v", err)
	}
	
	// Process BGP data
	bgpData := bgpResp.GetData()
	fmt.Printf("Retrieved %d BGP monitoring records\n", len(bgpData))
	
	for _, record := range bgpData {
		fmt.Printf("BGP Session: %s, State: %s\n", 
			record.GetSessionId(), record.GetState())
	}
	
	return nil
}
```

### Circuit Monitoring

```go
func getCircuitMonitoring(client *graphiant_sdk.APIClient, bearerToken string) error {
	// Create circuit monitoring request
	circuitReq := graphiant_sdk.NewV2MonitoringCircuitsSummaryPostRequest()
	
	// Configure selectors
	selectors := []graphiant_sdk.V2MonitoringCircuitsSummaryPostRequestSelectorsInner{
		{
			Type:  graphiant_sdk.PtrString("enterprise"),
			Value: graphiant_sdk.PtrString("your-enterprise-id"),
		},
	}
	circuitReq.SetSelectors(selectors)
	
	// Get circuit monitoring data
	circuitResp, _, err := client.DefaultAPI.
		V2MonitoringCircuitsSummaryPost(context.Background()).
		Authorization(bearerToken).
		V2MonitoringCircuitsSummaryPostRequest(*circuitReq).
		Execute()
	
	if err != nil {
		return fmt.Errorf("circuit monitoring failed: %v", err)
	}
	
	// Process circuit data
	circuitData := circuitResp.GetData()
	fmt.Printf("Retrieved %d circuit monitoring records\n", len(circuitData))
	
	for _, circuit := range circuitData {
		fmt.Printf("Circuit: %s, Status: %s, Bandwidth: %d Mbps\n",
			circuit.GetCircuitId(), circuit.GetStatus(), circuit.GetBandwidthMbps())
	}
	
	return nil
}
```

### Error Handling

```go
func handleApiErrors(client *graphiant_sdk.APIClient, bearerToken string) {
	resp, httpRes, err := client.DefaultAPI.
		V1EdgesSummaryGet(context.Background()).
		Authorization(bearerToken).
		Execute()
	
	if err != nil {
		// Handle different types of errors
		if apiErr, ok := err.(*graphiant_sdk.GenericOpenAPIError); ok {
			fmt.Printf("API Error: %s\n", apiErr.Error())
			fmt.Printf("Response Body: %s\n", string(apiErr.Body()))
			
			// Check HTTP status code
			if httpRes != nil {
				switch httpRes.StatusCode {
				case 401:
					fmt.Println("Unauthorized: Check your credentials")
				case 403:
					fmt.Println("Forbidden: Check your permissions")
				case 404:
					fmt.Println("Not Found: Resource doesn't exist")
				case 500:
					fmt.Println("Server Error: Try again later")
				}
			}
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
		return
	}
	defer httpRes.Body.Close()
	
	// Process successful response
	edges := resp.GetEdgesSummary()
	fmt.Printf("Successfully retrieved %d edges\n", len(edges))
}
```

## 🎯 Convenient Wrapper Functions

The SDK includes high-level wrapper functions in [api_custom.go](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/api_custom.go) that simplify common operations:

### Allow Device Configuration only when the device is Ready

```go
import "github.com/Graphiant-Inc/graphiant-sdk-go"

// Poll device status and execute configuration when ready
func configureDeviceWhenReady(deviceID int64, config graphiant_sdk.V1DevicesDeviceIdConfigPutRequest) *http.Response {
    config := graphiant_sdk.NewConfiguration()
    client := graphiant_sdk.NewAPIClient(config)
    
    // Get authentication token
    apiClient, token := getAuthToken() // Your auth function
    
    // Poll device status every 30 seconds for up to 10 attempts
    // Execute configuration when device becomes ready
    return graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceID, config)
}
```

### Available Wrapper Functions
| `PollAndPutDeviceConfig(apiClient, token, deviceID, config)` | Poll device status and execute config when ready |

## 🛠️ Development

### Prerequisites

- Go 1.21+ (1.23 recommended)
- Git
- OpenAPI Generator (for code generation)

### CI/CD Workflows

This repository uses GitHub Actions for continuous integration and deployment:

- **Linting** ([lint.yml](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/.github/workflows/lint.yml)): Runs golangci-lint, gofmt, and go vet on pull requests and pushes
- **Testing** ([test.yml](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/.github/workflows/test.yml)): Runs `go test` with race detection and coverage across Go 1.21, 1.22, and 1.23
- **Building** ([build.yml](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/.github/workflows/build.yml)): Builds and verifies the Go module
- **Releasing** ([release.yml](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/.github/workflows/release.yml)): Creates git tags and GitHub releases (manual trigger, admin-only)

See [.github/workflows/README.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/.github/workflows/README.md) for detailed workflow documentation.

### Building from Source

```bash
# Clone repository
git clone https://github.com/Graphiant-Inc/graphiant-sdk-go
cd graphiant-sdk-go

# Install dependencies
go mod tidy

# Build the project
go build ./...

# Run tests
go get github.com/stretchr/testify/assert
go test ./...
```

### Code Generation

To regenerate the SDK from the latest API specification:

```bash
# Install OpenAPI Generator
brew install openapi-generator  # macOS
# or download from: https://github.com/OpenAPITools/openapi-generator

# Generate SDK
openapi-generator generate \
  -i graphiant_api_docs_v26.4.1.json \
  -g go \
  --git-user-id Graphiant-Inc \
  --git-repo-id graphiant-sdk-go \
  --package-name graphiant_sdk \
  --additional-properties=packageVersion=26.4.1
```

> **Note:** Download the latest API bundle from the Graphiant portal under **Support Hub** → **Developer Tools**. Set **`packageVersion`** to the SDK release you are publishing (this branch: **26.4.1**). The **`-i`** filename reflects the API doc bundle version (here `graphiant_api_docs_v26.4.1.json`) and may stay the same across patch releases when the spec is unchanged.

### Testing

```bash
# Install test dependencies
go mod download

# Run all tests (tests requiring credentials will skip if not configured)
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detection
go test -race ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test
go test -v ./... -run Test_edge_summary
```

**Environment Variables for Tests:**

Tests that require API access will automatically use the following environment variables if set:

```bash
export GRAPHIANT_HOST="https://api.graphiant.com"  # Optional: API host (defaults to https://api.graphiant.io)
export GRAPHIANT_USERNAME="your_username"           # Required for integration tests
export GRAPHIANT_PASSWORD="your_password"          # Required for integration tests
```

- **`GRAPHIANT_HOST`** (optional): API host URL. If not set, defaults to `https://api.graphiant.io`. Supports formats like `https://api.test.graphiant.io` or `gcs:https://api.test.graphiant.io` (the `gcs:` prefix is automatically removed).
- **`GRAPHIANT_USERNAME`** (required for integration tests): Your Graphiant API username
- **`GRAPHIANT_PASSWORD`** (required for integration tests): Your Graphiant API password

Tests that require credentials will automatically skip if `GRAPHIANT_USERNAME` or `GRAPHIANT_PASSWORD` are not set, allowing the test suite to run successfully without credentials.

**Note**: The CI/CD pipeline automatically runs tests across multiple Go versions (1.21-1.23) on every pull request and push to main/develop branches. The pipeline reads credentials from GitHub secrets/variables when available.

### Project Structure

```
graphiant-sdk-go/
├── api_default.go              # Main API service
├── api_custom.go               # Convenient wrapper functions
├── client.go                   # HTTP client implementation
├── configuration.go            # Configuration management
├── model_*.go                  # Generated data models
├── response.go                 # Response handling
├── utils.go                    # Utility functions
├── version.go                  # Version information
├── test/                       # Test files
│   ├── sanity_test.go          # Basic functionality tests
│   ├── api_default_test.go     # API tests
│   └── version_test.go         # Version tests
├── docs/                       # Generated documentation
│   ├── api/                    # API documentation
│   ├── examples/               # Usage examples
│   └── *.md                    # Model documentation
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## 📖 API Reference

### Core Classes

- **`Configuration`**: Client configuration with authentication
- **`APIClient`**: HTTP client for API requests
- **`DefaultAPI`**: Main API interface with all endpoints

### Key Models

- **`V1AuthLoginPostRequest`**: Authentication request
- **`V1EdgesSummaryGetResponse`**: Device summary response
- **`V1DevicesDeviceIdConfigPutRequest`**: Device configuration request
- **`V2MonitoringBgpPostRequest`**: BGP monitoring request
- **`V2MonitoringCircuitsSummaryPostRequest`**: Circuit monitoring request

### Common Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/auth/login` | POST | Authenticate and get bearer token |
| `/v1/edges/summary` | GET | Get all device summaries |
| `/v1/devices/{device_id}` | GET | Get device details |
| `/v1/devices/{device_id}/config` | PUT | Update device configuration |
| `/v2/monitoring/bgp` | POST | Get BGP monitoring data |
| `/v2/monitoring/circuits/summary` | POST | Get circuit monitoring data |

## 🔐 Security

- **Authentication**: Bearer token-based authentication
- **HTTPS**: All API communications use HTTPS
- **Credentials**: Store credentials securely using environment variables
- **Token Management**: Bearer tokens expire and should be refreshed as needed

### Environment Variables

| Variable | Purpose |
|----------|---------|
| **`GRAPHIANT_ACCESS_TOKEN`** | Raw bearer JWT (same as Python SDK). Use with **`AuthorizationBearerFromEnv()`** for API calls. Typical source: `graphiant login` then `source ~/.graphiant/env.sh`. |
| **`GRAPHIANT_API_HOST`** | API base URL (e.g. `https://api.graphiant.com`). Preferred name; use with **`ConfigureHostFromEnv`**. |
| **`GRAPHIANT_HOST`** | Alternative API base URL for tools/tests; used if **`GRAPHIANT_API_HOST`** is empty. May include a `gcs:` prefix (stripped by **`ConfigureHostFromEnv`**). |
| **`GRAPHIANT_USERNAME`** / **`GRAPHIANT_PASSWORD`** | Used when **`GRAPHIANT_ACCESS_TOKEN`** is unset: **`AuthorizationBearerFromEnvOrLogin`** calls **`POST /v1/auth/login`**. Also used by integration tests. |

```bash
# Bearer-only automation (aligned with graphiant-sdk-python)
export GRAPHIANT_ACCESS_TOKEN="your_jwt"
export GRAPHIANT_API_HOST="https://api.graphiant.com"   # optional

# Password / test flows
export GRAPHIANT_USERNAME="your_username"
export GRAPHIANT_PASSWORD="your_password"
export GRAPHIANT_HOST="https://api.graphiant.com"      # optional; see table above
```

```go
token := graphiant_sdk.AccessTokenFromEnv()
authHeader := graphiant_sdk.AuthorizationBearerFromEnv()
// Or: authHeader, err := graphiant_sdk.AuthorizationBearerFromEnvOrLogin(ctx, client)
username := os.Getenv("GRAPHIANT_USERNAME")
password := os.Getenv("GRAPHIANT_PASSWORD")
```

**Note**: For detailed security policies, vulnerability reporting, and security best practices, see [SECURITY.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/SECURITY.md).

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes and ensure they pass local checks:
   ```bash
   # Format code
   gofmt -s -w .
   
   # Run linting
   golangci-lint run
   
   # Run static analysis
   go vet ./...
   
   # Run tests
   go test -v -race ./...
   ```
4. Commit your changes with a clear message (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

**Note**: All pull requests automatically run CI/CD checks (linting, testing across multiple Go versions). Ensure all checks pass before requesting review.

See [CONTRIBUTING.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/CONTRIBUTING.md) for detailed contribution guidelines.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/LICENSE) file for details.

## 📋 Version History

For a complete list of changes, new features, and version history, see [CHANGELOG.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/CHANGELOG.md).

The changelog follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) format and includes:
- All releases from v25.6.1 (initial release) to the latest version
- Breaking changes and deprecations
- Bug fixes and security updates

## 🆘 Support

- **Official Documentation**: [Graphiant SDK Go Guide](https://docs.graphiant.com/docs/graphiant-sdk-go) <-> [Graphiant Automation Docs](https://docs.graphiant.com/docs/automation)
- **API Reference**: [Graphiant SDK Go API Docs](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/docs/DefaultAPI.md) <-> [Graphiant Portal REST API Guide](https://docs.graphiant.com/docs/graphiant-portal-rest-api)
- **Changelog**: [CHANGELOG.md](https://github.com/Graphiant-Inc/graphiant-sdk-go/blob/main/CHANGELOG.md) - Version history and release notes
- **Issues**: [GitHub Issues](https://github.com/Graphiant-Inc/graphiant-sdk-go/issues)
- **Email**: support@graphiant.com

## 🔗 Related Projects

- [Graphiant SDK Python](https://github.com/Graphiant-Inc/graphiant-sdk-python)
- [Graphiant Playbooks](https://github.com/Graphiant-Inc/graphiant-playbooks)

---

**Made with ❤️ by the Graphiant Team**
