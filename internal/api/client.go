package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/internal/errors"
	log "github.com/yourusername/cyverApiCli/logger"
	"golang.org/x/net/http2"
)

var (
	verboseLevel int
	logger       = log.GetLogger(verboseLevel) // initialize with default verbosity 0
)

// APIClient represents the base API client
type TokenClient struct {
	BaseURL     string
	HTTPClient  *http.Client
	AccessToken string
	APIVersion  string
}

// APIClient represents the base API client
type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
	APIVersion string
}

// NewAPIClient creates a new API client with the given base URL, timeout, and API key
func NewAPIClient(baseURL string, timeout time.Duration, apiKey string) (*APIClient, error) {
	// Validate inputs
	if baseURL == "" {
		return nil, errors.NewCyverError(errors.ErrCodeConfigInvalid, "base URL cannot be empty", nil)
	}

	if timeout <= 0 {
		timeout = 30 * time.Second // Default timeout
	}

	// Configure HTTP/2.0 transport with connection pooling
	transport := &http.Transport{
		// Enable HTTP/2.0
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		// Connection pooling
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		// TLS configuration
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	// Explicitly configure transport for HTTP/2.0
	if err := http2.ConfigureTransport(transport); err != nil {
		logger.Error("Failed to configure HTTP/2.0 transport", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to configure HTTP/2.0 transport")
	}

	// Create a new cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Error("Failed to create cookie jar", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to create cookie jar")
	}
	logger.Debug("Cookie jar created successfully")

	// Initialize HTTP client with cookie jar
	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
		Jar:       jar,
	}

	return &APIClient{
		BaseURL:    baseURL,
		HTTPClient: client,
		APIKey:     apiKey,
	}, nil
}

// SetAPIVersion sets the API version for the client
func (c *APIClient) SetAPIVersion(version string) {
	c.APIVersion = version
}

// prepareRequestBody marshals the request body into JSON and returns a reader
func (c *APIClient) prepareRequestBody(body interface{}) (io.Reader, error) {
	if body == nil {
		logger.Debug("No body provided to prepareRequestBody")
		return nil, nil
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("Failed to marshal request body", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to marshal request body")
	}

	logger.Debug("Request body prepared", "size", len(jsonBody))
	return bytes.NewReader(jsonBody), nil
}

// createRequest constructs an HTTP request with the given method, URL, and body
func (c *APIClient) createRequest(method, url string, bodyReader io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		logger.Error("Failed to create request", "method", method, "url", url, "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to create HTTP request")
	}

	logger.Debug("HTTP request created", "method", method, "url", url)
	return req, nil
}

// setRequestHeaders adds necessary headers to the HTTP request
func (c *APIClient) setRequestHeaders(req *http.Request) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "CyverAPI-CLI/1.0")

	// Check if API key is provided
	if c.APIKey != "" {
		logger.Debug("Setting X-API-Key header")
		req.Header.Set("X-API-Key", c.APIKey)
	} else {
		// If API key is blank, try to load access token from viper config
		accessToken := viper.GetString("token.access_token")
		if accessToken != "" {
			logger.Debug("Setting Authorization Bearer header")
			req.Header.Set("Authorization", "Bearer "+accessToken)
		} else {
			logger.Warn("No API key or access token found for authentication")
			return errors.NewCyverError(errors.ErrCodeAuthFailed, "no authentication credentials found", nil)
		}
	}

	// Add API version header if set
	if c.APIVersion != "" {
		req.Header.Set("X-API-Version", c.APIVersion)
	}

	logger.Debug("Request headers set", "content_type", req.Header.Get("Content-Type"), "auth_present", req.Header.Get("Authorization") != "" || req.Header.Get("X-API-Key") != "")
	return nil
}

// logRequestDetails logs the request details based on verbosity level
func (c *APIClient) logRequestDetails(req *http.Request) error {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		logger.Error("Failed to dump request", "error", err)
		return errors.WrapError(err, errors.ErrCodeInternalError, "failed to dump request details")
	}

	logger.Debug("HTTP Request", "method", req.Method, "url", req.URL.String(), "headers", req.Header)
	logger.Debug("Request dump", "dump", string(requestDump))
	return nil
}

// executeRequest sends the HTTP request and returns the response
func (c *APIClient) executeRequest(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	duration := time.Since(start)

	if err != nil {
		logger.Error("HTTP request failed", "method", req.Method, "url", req.URL.String(), "duration", duration, "error", err)

		// Determine error type based on the error
		if isTimeoutError(err) {
			return nil, errors.NewCyverError(errors.ErrCodeAPITimeout, "request timeout", err)
		}
		if isNetworkError(err) {
			return nil, errors.NewCyverError(errors.ErrCodeAPINetworkError, "network error", err)
		}

		return nil, errors.WrapError(err, errors.ErrCodeAPINetworkError, "HTTP request failed")
	}

	logger.Debug("HTTP request completed", "method", req.Method, "url", req.URL.String(), "status", resp.StatusCode, "duration", duration)
	return resp, nil
}

// isTimeoutError checks if an error is a timeout error
func isTimeoutError(err error) bool {
	// Check for common timeout error patterns
	return fmt.Sprintf("%v", err) == "context deadline exceeded" ||
		fmt.Sprintf("%v", err) == "i/o timeout"
}

// isNetworkError checks if an error is a network error
func isNetworkError(err error) bool {
	// Check for common network error patterns
	errStr := fmt.Sprintf("%v", err)
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "no such host") ||
		strings.Contains(errStr, "network is unreachable")
}

// processResponse handles the HTTP response, logging and parsing the body
func (c *APIClient) processResponse(resp *http.Response, result interface{}) (interface{}, error) {
	defer resp.Body.Close()

	// Log response details
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		logger.Error("Failed to dump response", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to dump response details")
	}
	logger.Debug("HTTP Response", "status", resp.StatusCode, "headers", resp.Header)
	logger.Debug("Response dump", "dump", string(responseDump))

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to read response body")
	}

	// Check for error response
	if resp.StatusCode >= 400 {
		return c.handleErrorResponse(resp.StatusCode, respBody)
	}

	// Handle successful response
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return c.handleSuccessResponse(respBody, result)
	}

	// Handle other status codes
	logger.Warn("Unexpected status code", "status", resp.StatusCode)
	return respBody, nil
}

// handleErrorResponse processes error responses
func (c *APIClient) handleErrorResponse(statusCode int, respBody []byte) (interface{}, error) {
	// Try to parse error response
	var errResp struct {
		Error   string `json:"error"`
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	if err := json.Unmarshal(respBody, &errResp); err == nil {
		// Use structured error response
		message := errResp.Message
		if message == "" {
			message = errResp.Error
		}
		if message == "" {
			message = "API request failed"
		}

		logger.Error("API error response", "status", statusCode, "message", message, "code", errResp.Code)
		return nil, errors.NewAPIError(statusCode, message, nil)
	}

	// Fallback to generic error
	message := fmt.Sprintf("API request failed with status %d", statusCode)
	logger.Error("API request failed", "status", statusCode, "body", string(respBody))
	return nil, errors.NewAPIError(statusCode, message, nil)
}

// handleSuccessResponse processes successful responses
func (c *APIClient) handleSuccessResponse(respBody []byte, result interface{}) (interface{}, error) {
	// Attempt to unmarshal JSON response
	var jsonData interface{}
	if err := json.Unmarshal(respBody, &jsonData); err != nil {
		logger.Debug("Response is not JSON, returning raw body")
		return respBody, nil
	}

	// Pretty-print JSON for logging
	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		logger.Warn("Failed to pretty-print JSON", "error", err)
	} else {
		logger.Debug("JSON response formatted", "size", len(prettyJSON))
	}

	// Unmarshal into result if provided
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			logger.Warn("Failed to unmarshal response into result", "error", err)
		}
	}

	return prettyJSON, nil
}

// DoRequest performs an HTTP request with the specified method, path, body, and result
func (c *APIClient) DoRequest(method, path string, body interface{}, result interface{}) (interface{}, error) {
	url := c.BaseURL + path

	logger.Debug("Starting API request", "method", method, "path", path, "url", url)

	// Prepare request body
	bodyReader, err := c.prepareRequestBody(body)
	if err != nil {
		logger.Error("Failed to prepare request body", "error", err)
		return nil, err
	}

	// Create request
	req, err := c.createRequest(method, url, bodyReader)
	if err != nil {
		logger.Error("Failed to create request", "error", err)
		return nil, err
	}

	// Set headers
	if err := c.setRequestHeaders(req); err != nil {
		logger.Error("Failed to set request headers", "error", err)
		return nil, err
	}

	// Log request details
	if err := c.logRequestDetails(req); err != nil {
		logger.Error("Failed to log request details", "error", err)
		return nil, err
	}

	// Execute request
	resp, err := c.executeRequest(req)
	if err != nil {
		logger.Error("Failed to execute request", "error", err)
		return nil, err
	}

	// Process response
	response, err := c.processResponse(resp, result)
	if err != nil {
		logger.Error("Failed to process response", "error", err)
		return nil, err
	}

	logger.Debug("API request completed successfully")
	return response, nil
}

// DoRequestRaw performs an HTTP request and returns the raw response body
func (c *APIClient) DoRequestRaw(method, path string, body interface{}) ([]byte, error) {
	url := c.BaseURL + path

	logger.Debug("Starting raw API request", "method", method, "path", path, "url", url)

	// Prepare request body
	bodyReader, err := c.prepareRequestBody(body)
	if err != nil {
		logger.Error("Failed to prepare request body", "error", err)
		return nil, err
	}

	// Create request
	req, err := c.createRequest(method, url, bodyReader)
	if err != nil {
		logger.Error("Failed to create request", "error", err)
		return nil, err
	}

	// Set headers
	if err := c.setRequestHeaders(req); err != nil {
		logger.Error("Failed to set request headers", "error", err)
		return nil, err
	}

	// Log request details
	if err := c.logRequestDetails(req); err != nil {
		logger.Error("Failed to log request details", "error", err)
		return nil, err
	}

	// Execute request
	resp, err := c.executeRequest(req)
	if err != nil {
		logger.Error("Failed to execute request", "error", err)
		return nil, err
	}

	// Read raw response body
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeInternalError, "failed to read response body")
	}

	// Check for error response
	if resp.StatusCode >= 400 {
		_, err := c.handleErrorResponse(resp.StatusCode, respBody)
		return nil, err
	}

	logger.Debug("Raw API request completed successfully", "body_size", len(respBody))
	return respBody, nil
}

// DoRequestWithRetry performs an HTTP request with retry logic
func (c *APIClient) DoRequestWithRetry(ctx context.Context, method, path string, body interface{}, result interface{}, config *errors.RetryConfig) (interface{}, error) {
	return errors.RetryWithResult(ctx, func() (interface{}, error) {
		return c.DoRequest(method, path, body, result)
	}, config)
}

// DoRequestRawWithRetry performs a raw HTTP request with retry logic
func (c *APIClient) DoRequestRawWithRetry(ctx context.Context, method, path string, body interface{}, config *errors.RetryConfig) ([]byte, error) {
	return errors.RetryWithResult(ctx, func() ([]byte, error) {
		return c.DoRequestRaw(method, path, body)
	}, config)
}
