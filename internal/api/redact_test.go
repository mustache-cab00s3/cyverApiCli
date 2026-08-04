package api

import (
	"net/http"
	"strings"
	"testing"
)

func TestSanitizeHeaders_AuthorizationAndAPIKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer super-secret-token")
	h.Set("X-API-Key", "my-api-key")
	h.Set("Content-Type", "application/json")

	got := SanitizeHeaders(h)
	if got.Get("Authorization") != "Bearer [REDACTED]" {
		t.Fatalf("Authorization = %q, want Bearer [REDACTED]", got.Get("Authorization"))
	}
	if got.Get("X-API-Key") != "[REDACTED]" {
		t.Fatalf("X-API-Key = %q, want [REDACTED]", got.Get("X-API-Key"))
	}
	if got.Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type changed: %q", got.Get("Content-Type"))
	}
	// Original headers must remain untouched for the live request.
	if h.Get("Authorization") != "Bearer super-secret-token" {
		t.Fatalf("original Authorization was mutated: %q", h.Get("Authorization"))
	}
	if h.Get("X-API-Key") != "my-api-key" {
		t.Fatalf("original X-API-Key was mutated: %q", h.Get("X-API-Key"))
	}
}

func TestSanitizeHTTPDump(t *testing.T) {
	dump := "POST /api/v2.2/pentester/findings HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.abc\r\n" +
		"X-API-Key: abc123xyz\r\n" +
		"Content-Type: application/json\r\n" +
		"\r\n" +
		`{"title":"test"}`

	got := SanitizeHTTPDumpString(dump)
	if strings.Contains(got, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9") {
		t.Fatalf("token still present in dump:\n%s", got)
	}
	if strings.Contains(got, "abc123xyz") {
		t.Fatalf("api key still present in dump:\n%s", got)
	}
	if !strings.Contains(got, "Authorization: Bearer [REDACTED]") {
		t.Fatalf("missing redacted Authorization:\n%s", got)
	}
	if !strings.Contains(got, "X-API-Key: [REDACTED]") {
		t.Fatalf("missing redacted X-API-Key:\n%s", got)
	}
	if !strings.Contains(got, `{"title":"test"}`) {
		t.Fatalf("request body was altered:\n%s", got)
	}
}
