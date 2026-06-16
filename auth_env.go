package graphiant_sdk

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Environment variable names aligned with graphiant-sdk-python and common automation.
const (
	// EnvAccessToken is the bearer token from the portal / graphiant CLI (e.g. after sourcing ~/.graphiant/env.sh).
	EnvAccessToken = "GRAPHIANT_ACCESS_TOKEN"
	// EnvAPIHost is the API base URL (Python SDK name). Example: https://api.graphiant.com
	EnvAPIHost = "GRAPHIANT_API_HOST"
	// EnvHost is the API base URL as used in this repo's tests and older examples. If set and EnvAPIHost is empty, ConfigureHostFromEnv uses this.
	EnvHost = "GRAPHIANT_HOST"
	// EnvUsername and EnvPassword are used for POST /v1/auth/login when [EnvAccessToken] is unset.
	EnvUsername = "GRAPHIANT_USERNAME"
	EnvPassword = "GRAPHIANT_PASSWORD" // #nosec G101 — env var name constant, not a credential value
)

// AccessTokenFromEnv returns the trimmed raw access token from [EnvAccessToken], or empty if unset.
func AccessTokenFromEnv() string {
	return strings.TrimSpace(os.Getenv(EnvAccessToken))
}

// AuthorizationBearerFromEnv returns a value suitable for generated API methods' Authorization(...):
// "Bearer <token>" when [EnvAccessToken] is set, or "" when unset.
// If the env value already starts with "Bearer " (case-insensitive), it is returned unchanged.
func AuthorizationBearerFromEnv() string {
	tok := AccessTokenFromEnv()
	if tok == "" {
		return ""
	}
	if len(tok) > 7 && strings.EqualFold(tok[:7], "bearer ") {
		return tok
	}
	return "Bearer " + tok
}

// apiHostFromEnv returns the first non-empty of GRAPHIANT_API_HOST, GRAPHIANT_HOST.
func apiHostFromEnv() string {
	if h := strings.TrimSpace(os.Getenv(EnvAPIHost)); h != "" {
		return h
	}
	return strings.TrimSpace(os.Getenv(EnvHost))
}

// UsernameFromEnv returns the trimmed value of [EnvUsername].
func UsernameFromEnv() string {
	return strings.TrimSpace(os.Getenv(EnvUsername))
}

// PasswordFromEnv returns the trimmed value of [EnvPassword].
func PasswordFromEnv() string {
	return strings.TrimSpace(os.Getenv(EnvPassword))
}

// LoginBearerFromEnvCredentials calls POST /v1/auth/login using [UsernameFromEnv] and [PasswordFromEnv].
// It returns a value suitable for Authorization(...): "Bearer <jwt>". If either credential is empty, it returns an error.
func LoginBearerFromEnvCredentials(ctx context.Context, client *APIClient) (string, error) {
	user := UsernameFromEnv()
	pass := PasswordFromEnv()
	if user == "" || pass == "" {
		return "", fmt.Errorf("%s and %s must be set when %s is unset", EnvUsername, EnvPassword, EnvAccessToken)
	}
	authReq := NewV1AuthLoginPostRequestWithDefaults()
	authReq.SetUsername(user)
	authReq.SetPassword(pass)
	resp, httpRes, err := client.DefaultAPI.V1AuthLoginPost(ctx).V1AuthLoginPostRequest(*authReq).Execute()
	if httpRes != nil {
		defer func() { _ = httpRes.Body.Close() }()
	}
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("login response body missing")
	}
	if !resp.GetAuth() {
		return "", fmt.Errorf("login failed (auth=false)")
	}
	tok := resp.GetToken()
	if tok == "" {
		return "", fmt.Errorf("login response missing token")
	}
	return "Bearer " + tok, nil
}

// AuthorizationBearerFromEnvOrLogin returns [AuthorizationBearerFromEnv] when [EnvAccessToken] is set;
// otherwise it performs [LoginBearerFromEnvCredentials]. This mirrors graphiant-sdk-python: token from env first, else username/password login.
func AuthorizationBearerFromEnvOrLogin(ctx context.Context, client *APIClient) (string, error) {
	if s := AuthorizationBearerFromEnv(); s != "" {
		return s, nil
	}
	return LoginBearerFromEnvCredentials(ctx, client)
}

// ConfigureHostFromEnv sets [Configuration.Scheme] and [Configuration.Host] when an API base URL
// is present in the environment ([EnvAPIHost] or [EnvHost]). Intended to mirror Python's GRAPHIANT_API_HOST.
// A "gcs:" prefix (used in some test setups) is stripped. Hostnames without a scheme get https.
// If no host env is set, cfg is unchanged.
func ConfigureHostFromEnv(cfg *Configuration) {
	raw := apiHostFromEnv()
	if raw == "" {
		return
	}
	raw = strings.TrimPrefix(strings.TrimSpace(raw), "gcs:")
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return
	}
	if !strings.Contains(raw, "://") {
		raw = "https://" + raw
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return
	}
	if u.Scheme != "" {
		cfg.Scheme = u.Scheme
	}
	cfg.Host = u.Host
}
