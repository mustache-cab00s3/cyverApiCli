package v2_2

import (
	"time"

	"github.com/yourusername/cyverApiCli/internal/api"
	log "github.com/yourusername/cyverApiCli/logger"
)

var (
	verboseLevel int
	logger       = log.GetLogger(verboseLevel) // initialize with default verbosity 0
)

// getLogger returns a logger with the current verbose level
func getLogger() *log.Logger {
	return log.GetLogger(verboseLevel)
}

// SetVerboseLevel updates the verbose level for logging
func SetVerboseLevel(level int) {
	verboseLevel = level
}

// Client represents the V2.2 API client
type Client struct {
	*api.APIClient
	PentesterOps *PentesterOps
	TokenAuthOps *TokenAuthOps
	ClientOps    *ClientOps
}

// NewClient creates a new V2.2 API client
func NewClient(baseURL string, timeout time.Duration, apiKey string) *Client {
	apiClient, err := api.NewAPIClient(baseURL, timeout, apiKey)
	client := &Client{
		APIClient: apiClient,
	}
	client.PentesterOps = &PentesterOps{Client: client}
	client.TokenAuthOps = &TokenAuthOps{Client: client}
	client.ClientOps = &ClientOps{Client: client}
	if err != nil {
		logger.Error("failed to start NewClient %v", "error", err)
		return client
	}

	return client
}

// SetAPIVersion sets the API version in the base URL
func (c *Client) SetAPIVersion(version string) {
	// Set the APIVersion on the embedded base client. This is typically used
	// for setting the X-API-Version header.
	c.APIClient.SetAPIVersion(version)

	// DO NOT modify c.APIClient.BaseURL here.
	// The paths constructed in operations files (e.g., pentester_ops.go)
	// already include the full versioned path like "/api/v2.2/...".
	// The original line that caused the duplication was:
	// c.APIClient.BaseURL = fmt.Sprintf("%s/v2.2", c.APIClient.BaseURL)
}

// DoRequest performs an API request and unmarshals the response
func (c *Client) DoRequest(method, path string, body interface{}, result interface{}) (interface{}, error) {
	// c.APIClient.DoRequest is now assumed to return (interface{}, error)
	// Capture both return values and pass them through.
	data, err := c.APIClient.DoRequest(method, path, body, result)
	return data, err
}
