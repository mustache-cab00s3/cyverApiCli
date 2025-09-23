package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/internal/api/versions"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	log "github.com/yourusername/cyverApiCli/logger"
)

func versionedApiClient() interface{} {
	apiKey, baseURL, apiVersionString, err := LoadConfig()
	if err != nil {
		log.GetLogger(verboseLevel).Error("Failed to load config", err)
		return nil
	}

	// Use the client factory from internal/api/versions
	genericClient, err := versions.NewClient(versions.APIVersion(apiVersionString), baseURL, apiKey)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error creating API client for version", "apiVersionString", apiVersionString, "error", err)
		return nil
	}

	switch apiVersionString {
	case "v2.2", "latest":
		v2_2Client, ok := genericClient.(*v2_2.Client)
		if !ok {
			log.GetLogger(verboseLevel).Error("API client for version 'v2.2' is not of expected type *v2_2.Client. Got", "genericClient", genericClient)
			return nil
		}
		// Set the verbose level for the v2.2 client
		v2_2.SetVerboseLevel(verboseLevel)
		return v2_2Client
	default:
		log.GetLogger(verboseLevel).Error("Unsupported API version", "Supported Version", "v2.2, latest", "apiVersionString", apiVersionString)
		return nil
	}
}

// CustomURLResponse represents the response structure for custom URL requests
type CustomURLResponse struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       interface{}         `json:"body"`
	URL        string              `json:"url"`
	Method     string              `json:"method"`
	Duration   string              `json:"duration"`
}

// DirectHTTPRequest performs a direct HTTP request with the specified method, URL, and data
func DirectHTTPRequest(method, urlPath, data string) (*CustomURLResponse, error) {
	// Load configuration to get base URL and API key
	// LoadConfig() already handles token validation and refresh
	apiKey, baseURL, _, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Construct full URL
	var fullURL string
	if strings.HasPrefix(urlPath, "http://") || strings.HasPrefix(urlPath, "https://") {
		fullURL = urlPath
	} else {
		// Remove leading slash if present and construct full URL
		cleanPath := strings.TrimPrefix(urlPath, "/")
		fullURL = strings.TrimSuffix(baseURL, "/") + "/" + cleanPath
	}

	log.GetLogger(verboseLevel).Info("Making direct HTTP request", "method", method, "url", fullURL)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Prepare request body
	var bodyReader io.Reader
	if data != "" {
		bodyReader = strings.NewReader(data)
	}

	// Create request
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "cyverApiCli/1.0")

	// Check for access token in config first (prioritize token over API key)
	accessToken := viper.GetString("token.access_token")
	if accessToken != "" {
		log.GetLogger(verboseLevel).Info("Using access token for authentication")
		req.Header.Set("Authorization", "Bearer "+accessToken)
	} else if apiKey != "" {
		log.GetLogger(verboseLevel).Info("Using API key for authentication")
		req.Header.Set("Authorization", "Bearer "+apiKey)
	} else {
		log.GetLogger(verboseLevel).Warn("No authentication credentials found (no access token or API key)")
		log.GetLogger(verboseLevel).Info("Request will be made without authentication")
	}

	// Record start time
	startTime := time.Now()

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Record duration
	duration := time.Since(startTime)

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response if possible
	var bodyInterface interface{}
	if len(bodyBytes) > 0 {
		if err := json.Unmarshal(bodyBytes, &bodyInterface); err != nil {
			// If JSON parsing fails, return as string
			bodyInterface = string(bodyBytes)
		}
	}

	// Create response
	response := &CustomURLResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       bodyInterface,
		URL:        fullURL,
		Method:     method,
		Duration:   duration.String(),
	}

	return response, nil
}

// customUrlCmd represents the customUrl command
var customUrlCmd = &cobra.Command{
	Use:   "customUrl [URL_PATH]",
	Short: "Make a direct HTTP request to a custom URL path",
	Long: `Make a direct HTTP request to a custom URL path using the configured base URL and authentication.
The URL path can be relative to the configured base URL or a full URL.

Authentication:
  - Uses access token from config if available (prioritized)
  - Falls back to API key if no access token
  - Automatically validates and refreshes tokens before requests

Examples:
  cyverApiCli customUrl /api/v2.2/users
  cyverApiCli customUrl /api/v2.2/users --method POST --data '{"name":"John"}'
  cyverApiCli customUrl https://api.example.com/endpoint --method PUT --data @data.json`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlPath := args[0]
		method, _ := cmd.Flags().GetString("method")
		data, _ := cmd.Flags().GetString("data")

		// Handle data from file if @ prefix is used
		if strings.HasPrefix(data, "@") {
			filePath := strings.TrimPrefix(data, "@")
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				log.GetLogger(verboseLevel).Error("Failed to read data file", "file", filePath, "error", err)
				fmt.Printf("Error reading data file %s: %v\n", filePath, err)
				os.Exit(1)
			}
			data = string(fileData)
		}

		// Make the request
		response, err := DirectHTTPRequest(method, urlPath, data)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Direct HTTP request failed", "error", err)
			fmt.Printf("Error making request: %v\n", err)
			os.Exit(1)
		}

		// Output response as JSON
		jsonOutput, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			log.GetLogger(verboseLevel).Error("Failed to marshal response to JSON", "error", err)
			fmt.Printf("Error formatting response: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(jsonOutput))
	},
}

func init() {
	// Add customUrl command to root
	rootCmd.AddCommand(customUrlCmd)

	// Add flags to customUrl command
	customUrlCmd.Flags().StringP("method", "m", "GET", "HTTP method (GET, POST, PUT, DELETE)")
	customUrlCmd.Flags().StringP("data", "d", "", "Request data (JSON string or @filename for file input)")
}
