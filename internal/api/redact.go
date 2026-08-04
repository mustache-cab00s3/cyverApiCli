package api

import (
	"net/http"
	"regexp"
	"strings"
)

const redactedValue = "[REDACTED]"

var (
	// Matches Authorization header lines in httputil dumps (case-insensitive header name).
	dumpAuthorizationRe = regexp.MustCompile(`(?im)^(Authorization:\s*)(.+)$`)
	// Matches X-API-Key header lines in httputil dumps (hyphen variants).
	dumpAPIKeyRe = regexp.MustCompile(`(?im)^(X-API-Key:\s*)(.+)$`)
)

// SanitizeHeaders returns a copy of headers with Authorization and X-API-Key values redacted.
func SanitizeHeaders(headers http.Header) http.Header {
	if headers == nil {
		return nil
	}
	out := make(http.Header, len(headers))
	for key, values := range headers {
		if isSensitiveHeader(key) {
			redacted := make([]string, len(values))
			for i, v := range values {
				redacted[i] = sanitizeAuthHeaderValue(key, v)
			}
			out[key] = redacted
			continue
		}
		copied := make([]string, len(values))
		copy(copied, values)
		out[key] = copied
	}
	return out
}

// SanitizeHTTPDump redacts Authorization and X-API-Key values from an HTTP request/response dump.
func SanitizeHTTPDump(dump []byte) string {
	return SanitizeHTTPDumpString(string(dump))
}

// SanitizeHTTPDumpString redacts Authorization and X-API-Key values from dump text.
func SanitizeHTTPDumpString(dump string) string {
	if dump == "" {
		return dump
	}
	sanitized := dumpAuthorizationRe.ReplaceAllStringFunc(dump, func(line string) string {
		parts := dumpAuthorizationRe.FindStringSubmatch(line)
		if len(parts) != 3 {
			return line
		}
		return parts[1] + sanitizeBearerOrSecret(parts[2])
	})
	sanitized = dumpAPIKeyRe.ReplaceAllStringFunc(sanitized, func(line string) string {
		parts := dumpAPIKeyRe.FindStringSubmatch(line)
		if len(parts) != 3 {
			return line
		}
		return parts[1] + redactedValue
	})
	return sanitized
}

func isSensitiveHeader(name string) bool {
	switch http.CanonicalHeaderKey(name) {
	case "Authorization", "X-Api-Key":
		return true
	default:
		return false
	}
}

func sanitizeAuthHeaderValue(headerName, value string) string {
	if http.CanonicalHeaderKey(headerName) == "Authorization" {
		return sanitizeBearerOrSecret(value)
	}
	return redactedValue
}

func sanitizeBearerOrSecret(value string) string {
	trimmed := strings.TrimSpace(value)
	const bearerPrefix = "Bearer "
	if len(trimmed) >= len(bearerPrefix) && strings.EqualFold(trimmed[:len(bearerPrefix)], bearerPrefix) {
		return "Bearer " + redactedValue
	}
	if trimmed == "" {
		return value
	}
	return redactedValue
}
