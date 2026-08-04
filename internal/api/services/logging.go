package services

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/yourusername/cyverApiCli/internal/api"
	log "github.com/yourusername/cyverApiCli/logger"
)

var (
	verboseLevel int
	logger       = log.GetLogger(verboseLevel)
)

// SetVerboseLevel updates verbosity for non-supported service request logging.
func SetVerboseLevel(level int) {
	verboseLevel = level
	logger = log.GetLogger(verboseLevel)
}

func getLogger() *log.Logger {
	return log.GetLogger(verboseLevel)
}

func logServiceRequest(req *http.Request) {
	getLogger().Debug("Service HTTP request", "method", req.Method, "url", req.URL.String(), "headers", api.SanitizeHeaders(req.Header))
	if verboseLevel < 3 {
		return
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		getLogger().Warn("Failed to dump service request", "error", err)
		return
	}

	fmt.Fprintf(os.Stderr, "\n=== RAW SERVICE HTTP REQUEST ===\n")
	fmt.Fprintf(os.Stderr, "%s\n", api.SanitizeHTTPDump(dump))
	fmt.Fprintf(os.Stderr, "===============================\n\n")
}

func logServiceResponse(resp *http.Response, body []byte, duration time.Duration) {
	safeHeaders := api.SanitizeHeaders(resp.Header)
	getLogger().Debug(
		"Service HTTP response",
		"status", resp.StatusCode,
		"url", resp.Request.URL.String(),
		"duration", duration.String(),
		"headers", safeHeaders,
		"body_size", len(body),
	)
	if verboseLevel < 3 {
		return
	}

	fmt.Fprintf(os.Stderr, "\n=== RAW SERVICE HTTP RESPONSE ===\n")
	fmt.Fprintf(os.Stderr, "Status: %s\n", resp.Status)
	fmt.Fprintf(os.Stderr, "Headers:\n")
	for k, vals := range safeHeaders {
		fmt.Fprintf(os.Stderr, "  %s: %s\n", k, strings.Join(vals, ", "))
	}
	fmt.Fprintf(os.Stderr, "\nBody:\n")
	if isBinaryContentType(resp.Header.Get("Content-Type")) {
		fmt.Fprintf(os.Stderr, "<binary body omitted; %d bytes>\n", len(body))
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", string(body))
	}
	fmt.Fprintf(os.Stderr, "=================================\n\n")
}

func isBinaryContentType(contentType string) bool {
	ct := strings.ToLower(contentType)
	if ct == "" {
		return false
	}
	return !(strings.Contains(ct, "application/json") ||
		strings.Contains(ct, "text/") ||
		strings.Contains(ct, "application/xml") ||
		strings.Contains(ct, "application/javascript"))
}

