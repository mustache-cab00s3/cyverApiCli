package versions

import (
	"fmt"
	"time"

	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
)

// APIVersion represents the available API versions
type APIVersion string

const (
	V2_2   APIVersion = "v2.2"
	Latest APIVersion = "latest"
)

// NewClient creates a new API client for the specified version
func NewClient(version APIVersion, baseURL, apiKey string) (interface{}, error) {
	timeout := 30 * time.Second // Default timeout
	switch version {
	case V2_2:
		client := v2_2.NewClient(baseURL, timeout, apiKey)
		client.SetAPIVersion(string(version))
		return client, nil
	case Latest:
		// Latest is synonymous with v2.2
		client := v2_2.NewClient(baseURL, timeout, apiKey)
		client.SetAPIVersion("v2.2") // Set to v2.2 instead of "latest"
		return client, nil
	default:
		return nil, fmt.Errorf("unsupported API version: %s", version)
	}
}
