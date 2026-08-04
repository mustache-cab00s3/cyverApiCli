package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/yourusername/cyverApiCli/internal/api"
)

// NonSupportedServiceClient is for non-documented web-app service endpoints under /api/services/...
// and related app URLs. These endpoints are not guaranteed stable and should only be used for specific instances.
type NonSupportedServiceClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func (c *NonSupportedServiceClient) requireBearerToken() (string, error) {
	token := strings.TrimSpace(c.Token)
	if token == "" {
		return "", fmt.Errorf("authorization bearer token is required for services requests")
	}
	return token, nil
}

// NewNonSupportedServiceClient returns a client with BaseURL trimmed of trailing slashes.
func NewNonSupportedServiceClient(baseURL string, timeout time.Duration) *NonSupportedServiceClient {
	return &NonSupportedServiceClient{
		BaseURL: strings.TrimSuffix(baseURL, "/"),
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// DoGET performs an authenticated GET to BaseURL+endpoint with optional query string.
func (c *NonSupportedServiceClient) DoGET(ctx context.Context, endpoint string, query url.Values) ([]byte, int, error) {
	bearer, err := c.requireBearerToken()
	if err != nil {
		return nil, 0, err
	}

	reqURL := c.BaseURL + endpoint
	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}
	getLogger().Info("Making non-supported service request", "method", http.MethodGet, "url", reqURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", api.ChromeUserAgent)
	req.Header.Set("Authorization", "Bearer "+bearer)
	logServiceRequest(req)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response body: %w", err)
	}
	logServiceResponse(resp, body, time.Since(start))
	if resp.StatusCode >= 400 {
		return body, resp.StatusCode, fmt.Errorf("service request failed with status %d", resp.StatusCode)
	}

	return body, resp.StatusCode, nil
}

// DoJSON performs an authenticated JSON request (e.g. POST, PUT) to BaseURL+endpoint with optional query and body.
func (c *NonSupportedServiceClient) DoJSON(ctx context.Context, method, endpoint string, query url.Values, payload interface{}) ([]byte, int, error) {
	bearer, err := c.requireBearerToken()
	if err != nil {
		return nil, 0, err
	}

	reqURL := c.BaseURL + endpoint
	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}
	getLogger().Info("Making non-supported service request", "method", method, "url", reqURL)

	var bodyReader io.Reader
	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, 0, fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return nil, 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", api.ChromeUserAgent)
	req.Header.Set("Authorization", "Bearer "+bearer)
	logServiceRequest(req)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response body: %w", err)
	}
	logServiceResponse(resp, respBody, time.Since(start))
	if resp.StatusCode >= 400 {
		return respBody, resp.StatusCode, fmt.Errorf("service request failed with status %d", resp.StatusCode)
	}

	return respBody, resp.StatusCode, nil
}

// DoForm performs an authenticated form-encoded request (typically POST) to BaseURL+endpoint.
func (c *NonSupportedServiceClient) DoForm(ctx context.Context, method, endpoint string, query, form url.Values) ([]byte, int, error) {
	bearer, err := c.requireBearerToken()
	if err != nil {
		return nil, 0, err
	}

	reqURL := c.BaseURL + endpoint
	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}
	getLogger().Info("Making non-supported service request", "method", method, "url", reqURL)

	var bodyReader io.Reader
	if form != nil {
		bodyReader = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return nil, 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "text/html, */*; q=0.01")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", api.ChromeUserAgent)
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	logServiceRequest(req)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response body: %w", err)
	}
	logServiceResponse(resp, respBody, time.Since(start))
	if resp.StatusCode >= 400 {
		return respBody, resp.StatusCode, fmt.Errorf("service request failed with status %d", resp.StatusCode)
	}

	return respBody, resp.StatusCode, nil
}
