package graphiant_sdk_test

import (
	"context"
	"testing"

	openapiclient "github.com/Graphiant-Inc/graphiant-sdk-go"
)

func TestAccessTokenFromEnv(t *testing.T) {
	t.Setenv(openapiclient.EnvAccessToken, "")
	if openapiclient.AccessTokenFromEnv() != "" {
		t.Fatalf("expected empty")
	}
	t.Setenv(openapiclient.EnvAccessToken, "  tok  ")
	if openapiclient.AccessTokenFromEnv() != "tok" {
		t.Fatalf("expected trimmed token")
	}
}

func TestAuthorizationBearerFromEnv(t *testing.T) {
	t.Setenv(openapiclient.EnvAccessToken, "")
	if openapiclient.AuthorizationBearerFromEnv() != "" {
		t.Fatalf("expected empty when env unset")
	}
	t.Setenv(openapiclient.EnvAccessToken, "abc")
	if g := openapiclient.AuthorizationBearerFromEnv(); g != "Bearer abc" {
		t.Fatalf("got %q", g)
	}
	t.Setenv(openapiclient.EnvAccessToken, "Bearer xyz")
	if g := openapiclient.AuthorizationBearerFromEnv(); g != "Bearer xyz" {
		t.Fatalf("got %q", g)
	}
	t.Setenv(openapiclient.EnvAccessToken, "bearer lowercase")
	if g := openapiclient.AuthorizationBearerFromEnv(); g != "bearer lowercase" {
		t.Fatalf("got %q", g)
	}
}

func TestAuthorizationBearerFromEnvOrLogin_TokenPreferred(t *testing.T) {
	t.Setenv(openapiclient.EnvAccessToken, "abc")
	t.Setenv(openapiclient.EnvUsername, "u")
	t.Setenv(openapiclient.EnvPassword, "p")
	client := openapiclient.NewAPIClient(openapiclient.NewConfiguration())
	s, err := openapiclient.AuthorizationBearerFromEnvOrLogin(context.Background(), client)
	if err != nil || s != "Bearer abc" {
		t.Fatalf("got %q err %v", s, err)
	}
}

func TestAuthorizationBearerFromEnvOrLogin_ErrWhenNoTokenNoCreds(t *testing.T) {
	t.Setenv(openapiclient.EnvAccessToken, "")
	t.Setenv(openapiclient.EnvUsername, "")
	t.Setenv(openapiclient.EnvPassword, "")
	client := openapiclient.NewAPIClient(openapiclient.NewConfiguration())
	_, err := openapiclient.AuthorizationBearerFromEnvOrLogin(context.Background(), client)
	if err == nil {
		t.Fatal("expected error when token and credentials are missing")
	}
}

func TestConfigureHostFromEnv(t *testing.T) {
	t.Setenv(openapiclient.EnvAPIHost, "")
	t.Setenv(openapiclient.EnvHost, "")
	cfg := openapiclient.NewConfiguration()
	openapiclient.ConfigureHostFromEnv(cfg)
	if cfg.Host != "" {
		t.Fatalf("expected no host override")
	}

	t.Setenv(openapiclient.EnvAPIHost, "https://api.example.com")
	cfg = openapiclient.NewConfiguration()
	openapiclient.ConfigureHostFromEnv(cfg)
	if cfg.Scheme != "https" || cfg.Host != "api.example.com" {
		t.Fatalf("scheme=%q host=%q", cfg.Scheme, cfg.Host)
	}

	t.Setenv(openapiclient.EnvAPIHost, "")
	t.Setenv(openapiclient.EnvHost, "api.other.com")
	cfg = openapiclient.NewConfiguration()
	openapiclient.ConfigureHostFromEnv(cfg)
	if cfg.Scheme != "https" || cfg.Host != "api.other.com" {
		t.Fatalf("scheme=%q host=%q", cfg.Scheme, cfg.Host)
	}

	t.Setenv(openapiclient.EnvAPIHost, "gcs:https://gcs.example.com")
	t.Setenv(openapiclient.EnvHost, "")
	cfg = openapiclient.NewConfiguration()
	openapiclient.ConfigureHostFromEnv(cfg)
	if cfg.Host != "gcs.example.com" {
		t.Fatalf("host=%q", cfg.Host)
	}
}
