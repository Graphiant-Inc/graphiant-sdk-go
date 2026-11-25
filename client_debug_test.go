package graphiant_sdk

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRedactURL(t *testing.T) {
	cases := map[string]string{
		"https://api.graphiant.com/v1/users/jane%40example.com/password?token=s3cr3t": "https://api.graphiant.com/v1/users/[REDACTED]/password?token=[REDACTED]",
		"https://api.graphiant.com/v1/diagnostic/ping-stop/abc123":                    "https://api.graphiant.com/v1/diagnostic/ping-stop/[REDACTED]",
		"https://user:pw@api.graphiant.com/v1/devices?b=2&a=1":                        "https://api.graphiant.com/v1/devices?a=[REDACTED]&b=[REDACTED]",
		"https://api.graphiant.com/v1/devices/12345":                                  "https://api.graphiant.com/v1/devices/12345",
	}
	for in, want := range cases {
		u, err := url.Parse(in)
		if err != nil {
			t.Fatal(err)
		}
		if got := redactURL(u); got != want {
			t.Errorf("redactURL(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestRedactHeaders(t *testing.T) {
	h := http.Header{}
	h.Add("authorization", "Bearer secret-jwt")
	h.Add("X-Api-Key", "k-123")
	h.Add("Cookie", "session=abc")
	h.Add("Content-Type", "application/json")
	got := redactHeaders(h)
	for _, secret := range []string{"secret-jwt", "k-123", "session=abc"} {
		if strings.Contains(got, secret) {
			t.Errorf("redactHeaders leaked %q:\n%s", secret, got)
		}
	}
	if !strings.Contains(got, "Content-Type: application/json") {
		t.Errorf("redactHeaders dropped a non-sensitive header:\n%s", got)
	}
}

// End to end: with Debug on, nothing sensitive from the request or response
// reaches the log.
func TestCallAPIDebugLogRedacts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "resp-cookie-secret"})
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"token":"resp-body-secret"}`)
	}))
	defer srv.Close()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(io.Discard)

	cfg := NewConfiguration()
	cfg.Debug = true
	cfg.HTTPClient = srv.Client()
	c := NewAPIClient(cfg)

	req, _ := http.NewRequest(http.MethodPatch, srv.URL+"/v1/users/jane%40example.com/password?token=query-secret",
		strings.NewReader(`{"password":"body-secret"}`))
	req.Header.Set("Authorization", "Bearer jwt-secret")
	resp, err := c.callAPI(req)
	if err != nil {
		t.Fatal(err)
	}
	if err := resp.Body.Close(); err != nil {
		t.Fatal(err)
	}

	out := buf.String()
	for _, secret := range []string{"query-secret", "body-secret", "jwt-secret", "jane@example.com", "jane%40example.com", "resp-cookie-secret", "resp-body-secret"} {
		if strings.Contains(out, secret) {
			t.Errorf("debug log leaked %q:\n%s", secret, out)
		}
	}
	if !strings.Contains(out, "HTTP request: PATCH") || !strings.Contains(out, "200 OK") {
		t.Errorf("debug log is missing request/response metadata:\n%s", out)
	}
}
