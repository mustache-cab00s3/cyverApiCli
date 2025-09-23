package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/internal/api/versions"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	log "github.com/yourusername/cyverApiCli/logger"
)

// Helper function to safely get string value from pointer
func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
	Long:  `Manage CLI configuration settings and initialization.`,
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize CLI configuration",
	Long:  `Create a new configuration file with guided setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeConfig(); err != nil {
			fmt.Printf("Error initializing config: %v\n", err)
			os.Exit(1)
		}
	},
}

var viewConfigCmd = &cobra.Command{
	Use:   "view",
	Short: "View current configuration",
	Long:  `Display the current configuration settings with sensitive data partially obscured.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viewConfig(); err != nil {
			fmt.Printf("Error viewing config: %v\n", err)
			os.Exit(1)
		}
	},
}

var refreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "Refresh the access token using the refresh token",
	Long:  `Manually refresh the access token using the stored refresh token. This is useful when the access token has expired.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RefreshAccessToken(); err != nil {
			fmt.Printf("Error refreshing token: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Token refreshed successfully!")
	},
}

var reAuthCmd = &cobra.Command{
	Use:   "re-auth",
	Short: "Re-authenticate using stored email and prompt for password",
	Long:  `Re-authenticate using the stored email address from configuration. This will prompt for your password and update the stored tokens.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ReAuthenticate(); err != nil {
			fmt.Printf("Error re-authenticating: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(refreshTokenCmd)
	configCmd.AddCommand(reAuthCmd)
	rootCmd.AddCommand(configCmd)
}

// validateEmail checks if the provided string is a valid email address
func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// validateUsername checks if the provided string is a valid username or email address
func validateUsername(username string) bool {
	// Check if it's empty
	if username == "" {
		return false
	}

	// Check if it's a valid email address
	if validateEmail(username) {
		return true
	}

	// Check if it's a valid username (alphanumeric, underscores, hyphens, dots)
	// Username should be at least 3 characters and at most 50 characters
	if len(username) < 3 || len(username) > 50 {
		return false
	}

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	return usernameRegex.MatchString(username)
}

// validateAPIKey checks if the API key has a valid format
func validateAPIKey(key string) bool {
	// Basic validation: API key should be at least 32 characters
	// and should only contain alphanumeric characters and hyphens
	if len(key) < 32 {
		return false
	}
	keyRegex := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	return keyRegex.MatchString(key)
}

// validateURL checks if the provided string is a valid URL
func validateURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// validateAPIVersion checks if the provided version is valid
func validateAPIVersion(version string) bool {
	validVersions := map[string]bool{
		"v2.2":   true,
		"latest": true,
	}
	return validVersions[version]
}

func initializeConfig() error {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Set default config path
	configPath := filepath.Join(home, ".cyverApiCli.yaml")

	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("Configuration file already exists at %s\n", configPath)
		fmt.Print("Do you want to overwrite it? (y/N): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Configuration initialization cancelled.")
			return nil
		}
	}

	// Get API configuration
	fmt.Println("\nCyver API Configuration")
	fmt.Println("======================")

	// Get API version
	var apiVersion string
	for {
		fmt.Print("\nSelect API version (v2.2/latest) [latest]: ")
		fmt.Scanln(&apiVersion)
		if apiVersion == "" {
			apiVersion = "latest"
		}
		if validateAPIVersion(apiVersion) {
			break
		}
		fmt.Println("Invalid API version. Please select from: v2.2, latest")
	}

	// Get base URL
	var baseURL string
	for {
		fmt.Print("\nEnter API base URL [https://api.cyver.io]: ")
		fmt.Scanln(&baseURL)
		if baseURL == "" {
			baseURL = "https://api.cyver.io"
		}
		if validateURL(baseURL) {
			break
		}
		fmt.Println("Invalid URL. Please enter a valid URL starting with http:// or https://")
	}

	// Get API key (optional)
	var apiKey string
	for {
		fmt.Print("\nEnter your API key (optional, press Enter to skip): ")
		fmt.Scanln(&apiKey)
		if apiKey == "" || validateAPIKey(apiKey) {
			break
		}
		fmt.Println("Invalid API key. API key must be at least 32 characters long.")
	}

	// Get proxy settings
	fmt.Println("\nProxy Settings")
	fmt.Println("-------------")
	var useProxy string
	for {
		fmt.Print("Do you want to configure a proxy? (y/N): ")
		fmt.Scanln(&useProxy)
		if useProxy == "" || useProxy == "n" || useProxy == "N" {
			break
		}
		if useProxy == "y" || useProxy == "Y" {
			var proxyURL string
			for {
				fmt.Print("Enter proxy URL (e.g., http://proxy:port): ")
				fmt.Scanln(&proxyURL)
				if proxyURL == "" || validateURL(proxyURL) {
					break
				}
				fmt.Println("Invalid proxy URL. Please enter a valid URL starting with http:// or https://")
			}

			var proxyUser string
			fmt.Print("Enter proxy username (optional): ")
			fmt.Scanln(&proxyUser)

			var proxyPass string
			if proxyUser != "" {
				fmt.Print("Enter proxy password: ")
				fmt.Scanln(&proxyPass)
			}

			// Get logging settings
			fmt.Println("\nLogging Settings")
			fmt.Println("---------------")
			var logLevel string
			for {
				fmt.Print("Select log level (debug/info/warn/error) [info]: ")
				fmt.Scanln(&logLevel)
				if logLevel == "" {
					logLevel = "info"
				}
				if validateLogLevel(logLevel) {
					break
				}
				fmt.Println("Invalid log level. Please select from: debug, info, warn, error")
			}

			var logFile string
			fmt.Print("Enter log file path (leave empty for stdout): ")
			fmt.Scanln(&logFile)

			// Get output formatting settings
			fmt.Println("\nOutput Formatting")
			fmt.Println("----------------")
			var outputFormat string
			for {
				fmt.Print("Select output format (json/yaml/table) [table]: ")
				fmt.Scanln(&outputFormat)
				if outputFormat == "" {
					outputFormat = "table"
				}
				if validateOutputFormat(outputFormat) {
					break
				}
				fmt.Println("Invalid output format. Please select from: json, yaml, table")
			}

			var colorOutput string
			for {
				fmt.Print("Enable colored output? (y/N): ")
				fmt.Scanln(&colorOutput)
				if colorOutput == "" || colorOutput == "n" || colorOutput == "N" {
					colorOutput = "false"
					break
				}
				if colorOutput == "y" || colorOutput == "Y" {
					colorOutput = "true"
					break
				}
				fmt.Println("Please answer with 'y' or 'n'")
			}

			config := map[string]interface{}{
				"api": map[string]interface{}{
					"version":  apiVersion,
					"base_url": baseURL,
					"api_key":  apiKey,
				},
				"proxy": map[string]interface{}{
					"url":      proxyURL,
					"username": proxyUser,
					"password": proxyPass,
				},
				"client": map[string]interface{}{
					"timeout": 30,
				},
				"logging": map[string]interface{}{
					"level": logLevel,
					"file":  logFile,
				},
				"output": map[string]interface{}{
					"format": outputFormat,
					"color":  colorOutput == "y" || colorOutput == "Y",
				},
			}

			// Get client settings
			fmt.Println("\nClient Settings")
			fmt.Println("---------------")
			var timeout string
			for {
				fmt.Print("Enter request timeout in seconds [30]: ")
				fmt.Scanln(&timeout)
				if timeout == "" {
					timeout = "30"
				}
				if validateTimeout(timeout) {
					break
				}
				fmt.Println("Invalid timeout. Please enter a number between 1 and 300")
			}

			timeoutInt, _ := strconv.Atoi(timeout)
			config["client"] = map[string]interface{}{
				"timeout": timeoutInt,
			}

			// Write config to file
			viper.SetConfigFile(configPath)
			for key, value := range config {
				viper.Set(key, value)
			}
			if err := viper.WriteConfig(); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}

			// Set file permissions to user-only (600)
			if err := os.Chmod(configPath, 0600); err != nil {
				return fmt.Errorf("failed to set config file permissions: %w", err)
			}

			fmt.Printf("\nConfiguration saved to %s\n", configPath)
			fmt.Println("\nYou can now use the CLI with your configured settings.")
			fmt.Println("To modify these settings later, edit the config file or use 'cyverApiCli config init' again.")
			fmt.Println("To view your current configuration, use 'cyverApiCli config view'")

			return nil
		}
		fmt.Println("Please answer with 'y' or 'n'")
	}

	// If no proxy configuration, write basic config
	config := map[string]interface{}{
		"api": map[string]interface{}{
			"version":  apiVersion,
			"base_url": baseURL,
			"api_key":  apiKey,
		},
		"client": map[string]interface{}{
			"timeout": 30,
		},
		"logging": map[string]interface{}{
			"level": "info",
			"file":  "",
		},
		"output": map[string]interface{}{
			"format": "table",
			"color":  true,
		},
	}

	// Write config to file
	viper.SetConfigFile(configPath)
	for key, value := range config {
		viper.Set(key, value)
	}
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Set file permissions to user-only (600)
	if err := os.Chmod(configPath, 0600); err != nil {
		return fmt.Errorf("failed to set config file permissions: %w", err)
	}

	fmt.Printf("\nConfiguration saved to %s\n", configPath)
	fmt.Println("\nYou can now use the CLI with your configured settings.")
	fmt.Println("To modify these settings later, edit the config file or use 'cyverApiCli config init' again.")
	fmt.Println("To view your current configuration, use 'cyverApiCli config view'")

	// If API key is blank, trigger token authentication
	if apiKey == "" {
		fmt.Println("\nNo API key provided. Starting token authentication process...")

		// Prompt for username
		var username string
		for {
			fmt.Print("Enter your username or email address: ")
			fmt.Scanln(&username)
			if validateUsername(username) {
				break
			}
			fmt.Println("Please enter a valid username (3-50 characters, alphanumeric with dots, underscores, hyphens) or email address.")
		}

		// Store email for future re-authentication
		viper.Set("auth.email", username)

		// Create a temporary command to execute tokenAuthCmd
		tempCmd := &cobra.Command{}
		tempCmd.Flags().String("username", username, "")

		// Execute token authentication
		if err := handleClientSwitch(versionedApiClient(), tempCmd); err != nil {
			fmt.Printf("Token authentication failed: %v\n", err)
			fmt.Println("You can manually run 'cyverApiCli apiAuth getToken -u <username>' to authenticate later.")
		} else {
			fmt.Println("Token authentication completed successfully!")
		}
	}

	return nil
}

// validateLogLevel checks if the provided log level is valid
func validateLogLevel(level string) bool {
	validLevels := map[string]bool{
		"DEBUG": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	return validLevels[level]
}

// validateOutputFormat checks if the provided output format is valid
func validateOutputFormat(format string) bool {
	validFormats := map[string]bool{
		"json":  true,
		"yaml":  true,
		"table": true,
	}
	return validFormats[format]
}

// validateTimeout checks if the provided timeout is valid
func validateTimeout(timeout string) bool {
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return false
	}
	return timeoutInt >= 1 && timeoutInt <= 300
}

// obscureSensitiveData partially obscures sensitive data in the configuration
func obscureSensitiveData(value string) string {
	if len(value) <= 8 {
		return "********"
	}
	return value[:4] + "****" + value[len(value)-4:]
}

func viewConfig() error {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Set config path
	configPath := filepath.Join(home, ".cyverApiCli.yaml")
	viper.SetConfigFile(configPath)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Get and display configuration
	fmt.Println("\nCurrent Configuration:")
	fmt.Println("=====================")

	// API Configuration
	fmt.Println("\nAPI Settings:")
	fmt.Printf("  Version:  %s\n", viper.GetString("api.version"))
	fmt.Printf("  Base URL: %s\n", viper.GetString("api.base_url"))
	fmt.Printf("  API Key:  %s\n", obscureSensitiveData(viper.GetString("api.api_key")))

	// Proxy Configuration
	if viper.IsSet("proxy.url") {
		fmt.Println("\nProxy Settings:")
		fmt.Printf("  URL:      %s\n", viper.GetString("proxy.url"))
		if viper.GetString("proxy.username") != "" {
			fmt.Printf("  Username: %s\n", viper.GetString("proxy.username"))
			fmt.Printf("  Password: %s\n", obscureSensitiveData(viper.GetString("proxy.password")))
		}
	}

	// Logging Configuration
	fmt.Println("\nLogging Settings:")
	fmt.Printf("  Level:    %s\n", viper.GetString("logging.level"))
	fmt.Printf("  File:     %s\n", viper.GetString("logging.file"))

	// Output Configuration
	fmt.Println("\nOutput Settings:")
	fmt.Printf("  Format:   %s\n", viper.GetString("output.format"))
	fmt.Printf("  Color:    %v\n", viper.GetBool("output.color"))

	// Client Configuration
	fmt.Println("\nClient Settings:")
	fmt.Printf("  Timeout:  %d seconds\n", viper.GetInt("client.timeout"))

	// Authentication Configuration (if available)
	if viper.IsSet("auth.email") {
		fmt.Println("\nAuthentication Settings:")
		fmt.Printf("  Email:         %s\n", viper.GetString("auth.email"))
	}

	// Token Configuration (if available)
	if viper.IsSet("token.access_token") {
		fmt.Println("\nToken Settings:")
		fmt.Printf("  Access Token:  %s\n", obscureSensitiveData(viper.GetString("token.access_token")))
		if viper.IsSet("token.refresh_token") {
			fmt.Printf("  Refresh Token: %s\n", obscureSensitiveData(viper.GetString("token.refresh_token")))
		}
		if viper.IsSet("token.expireInSeconds") {
			fmt.Printf("  Expires In:    %d seconds\n", viper.GetInt32("token.expireInSeconds"))
		}
		if viper.IsSet("token.token_created_at") {
			fmt.Printf("  Created At:    %s\n", viper.GetString("token.token_created_at"))
		}

		// Check if token is expired
		if isExpired, err := IsTokenExpired(); err == nil {
			if isExpired {
				fmt.Println("  Status:        EXPIRED")
			} else {
				fmt.Println("  Status:        VALID")
			}
		}
	}

	fmt.Printf("\nConfiguration file location: %s\n", configPath)

	return nil
}

// LoadConfig loads and validates the configuration file using Viper.
// It returns an error if the config file is missing, inaccessible, or lacks required fields.
// It also automatically validates and refreshes tokens if necessary.
// ConfigLoader implements the shared.ConfigLoader interface
type ConfigLoader struct{}

func (c *ConfigLoader) LoadConfig() (apiKey, baseURL, apiVersion string, err error) {
	return LoadConfig()
}

func LoadConfig() (apiKey, baseURL, apiVersion string, err error) {
	// Debug: Print config file path and check if it exists
	configFile := viper.ConfigFileUsed()
	log.GetLogger(verboseLevel).Info("Config file path:", configFile)
	if configFile == "" {
		return "", "", "", fmt.Errorf("no config file specified. Run 'cyverApiCli config init' to configure")
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return "", "", "", fmt.Errorf("config file does not exist: %s. Run 'cyverApiCli config init' to create it", configFile)
	} else if err != nil {
		return "", "", "", fmt.Errorf("failed to access config file: %v", err)
	}

	// Debug: Print all loaded config keys
	log.GetLogger(verboseLevel).Info("Loaded config keys:", "Keys", viper.AllKeys())

	// Load configuration using viper
	apiKey = viper.GetString("api.api_key")
	baseURL = viper.GetString("api.base_url")
	apiVersion = viper.GetString("api.version")

	// Debug: Print loaded values
	log.GetLogger(verboseLevel).Info("Loaded values:", "api_key", obscureSensitiveData(apiKey), "base_url", baseURL, "api_version", apiVersion)

	// Validate configuration
	// Note: API key is now optional, so we don't validate it here
	if baseURL == "" {
		return "", "", "", fmt.Errorf("base URL is missing in config file: %s. Run 'cyverApiCli config init' to configure", configFile)
	}
	if apiVersion == "" {
		return "", "", "", fmt.Errorf("API version is missing in config file: %s. Run 'cyverApiCli config init' to configure", configFile)
	}

	// Validate and refresh token if necessary
	if err := ValidateAndRefreshToken(); err != nil {
		log.GetLogger(verboseLevel).Warn("Token validation/refresh failed", "error", err)
		// Don't fail the entire config load if token refresh fails
		// The user can still use API key authentication or manually re-authenticate
	}

	return apiKey, baseURL, apiVersion, nil
}

// TokenInfo represents token information stored in config
type TokenInfo struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	ExpiresIn             int32     `json:"expireInSeconds"`
	RefreshExpiresIn      int32     `json:"refresh_expires_in"`
	TokenCreatedAt        time.Time `json:"token_created_at"`
	RefreshTokenCreatedAt time.Time `json:"refresh_token_created_at"`
}

// IsTokenExpired checks if the access token has expired
func IsTokenExpired() (bool, error) {
	// Get token creation time and expiration duration
	tokenCreatedAtStr := viper.GetString("token.token_created_at")
	expiresIn := viper.GetInt32("token.expireInSeconds")

	// Debug logging to understand what's happening
	log.GetLogger(verboseLevel).Debug("Token expiration check - raw values",
		"token_created_at", tokenCreatedAtStr,
		"expireInSeconds", expiresIn)

	if tokenCreatedAtStr == "" {
		log.GetLogger(verboseLevel).Debug("Token creation time is empty, considering token expired")
		return true, nil
	}

	if expiresIn == 0 {
		log.GetLogger(verboseLevel).Debug("Token expiration duration is 0, considering token expired")
		return true, nil
	}

	// Parse the creation time
	tokenCreatedAt, err := time.Parse(time.RFC3339, tokenCreatedAtStr)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Failed to parse token creation time", "error", err)
		return true, fmt.Errorf("invalid token creation time: %w", err)
	}

	// Calculate expiration time
	expirationTime := tokenCreatedAt.Add(time.Duration(expiresIn) * time.Second)

	// Check if token is expired (with 5 minute buffer)
	bufferTime := 5 * time.Minute
	isExpired := time.Now().Add(bufferTime).After(expirationTime)

	log.GetLogger(verboseLevel).Debug("Token expiration check",
		"created_at", tokenCreatedAt,
		"expireInSeconds", expiresIn,
		"expiration_time", expirationTime,
		"is_expired", isExpired)

	return isExpired, nil
}

// IsRefreshTokenExpired checks if the refresh token has expired
func IsRefreshTokenExpired() (bool, error) {
	// Get refresh token creation time and expiration duration
	refreshTokenCreatedAtStr := viper.GetString("token.refresh_token_created_at")
	refreshExpiresIn := viper.GetInt32("token.refresh_expires_in")

	if refreshTokenCreatedAtStr == "" || refreshExpiresIn == 0 {
		// No refresh token information available, consider it expired
		return true, nil
	}

	// Parse the creation time
	refreshTokenCreatedAt, err := time.Parse(time.RFC3339, refreshTokenCreatedAtStr)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Failed to parse refresh token creation time", "error", err)
		return true, fmt.Errorf("invalid refresh token creation time: %w", err)
	}

	// Calculate expiration time
	expirationTime := refreshTokenCreatedAt.Add(time.Duration(refreshExpiresIn) * time.Second)

	// Check if refresh token is expired
	isExpired := time.Now().After(expirationTime)

	log.GetLogger(verboseLevel).Debug("Refresh token expiration check",
		"created_at", refreshTokenCreatedAt,
		"expireInSeconds", refreshExpiresIn,
		"expiration_time", expirationTime,
		"is_expired", isExpired)

	return isExpired, nil
}

// RefreshAccessToken refreshes the access token using the refresh token
func RefreshAccessToken() error {
	log.GetLogger(verboseLevel).Info("Starting token refresh process")

	// Check if refresh token exists
	refreshToken := viper.GetString("token.refresh_token")
	if refreshToken == "" {
		return fmt.Errorf("no refresh token available")
	}

	// Check if refresh token is expired
	isRefreshExpired, err := IsRefreshTokenExpired()
	if err != nil {
		return fmt.Errorf("failed to check refresh token expiration: %w", err)
	}

	if isRefreshExpired {
		return fmt.Errorf("refresh token has expired, please re-authenticate")
	}

	// Get API configuration
	baseURL := viper.GetString("api.base_url")
	apiVersion := viper.GetString("api.version")

	if baseURL == "" || apiVersion == "" {
		return fmt.Errorf("missing API configuration (base_url or version)")
	}

	// Create API client for token refresh directly without going through LoadConfig
	// to avoid infinite recursion
	apiKey := viper.GetString("api.api_key")
	genericClient, err := versions.NewClient(versions.APIVersion(apiVersion), baseURL, apiKey)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error creating API client for token refresh", "apiVersion", apiVersion, "error", err)
		return fmt.Errorf("failed to create API client for token refresh: %w", err)
	}

	// Perform token refresh based on API version
	var newTokenInfo *TokenInfo
	switch clientVersion := genericClient.(type) {
	case *v2_2.Client:
		if clientVersion.TokenAuthOps == nil {
			return fmt.Errorf("TokenAuthOps is nil for v2.2 client")
		}

		// Call refresh token API
		response, err := clientVersion.TokenAuthOps.ApiTokenauthRefreshtokenPost(refreshToken)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Token refresh API call failed", "error", err)
			return fmt.Errorf("token refresh API call failed: %w", err)
		}

		if !response.Success {
			errorMsg := "token refresh failed"
			if response.Error != nil && response.Error.Message != nil {
				errorMsg = *response.Error.Message
			}
			log.GetLogger(verboseLevel).Error("Token refresh failed", "error", errorMsg)
			return fmt.Errorf("token refresh failed: %s", errorMsg)
		}

		// Extract new token information
		if response.Result == nil {
			return fmt.Errorf("empty response from token refresh")
		}

		// Debug: Log the response details
		log.GetLogger(verboseLevel).Debug("Token refresh response",
			"AccessToken", response.Result.AccessToken,
			"RefreshToken", response.Result.RefreshToken,
			"ExpireInSeconds", response.Result.ExpireInSeconds)

		// Get the new refresh token, but keep the existing one if none is provided
		newRefreshToken := getStringValue(response.Result.RefreshToken)
		if newRefreshToken == "" {
			// If no new refresh token is provided, keep the existing one
			newRefreshToken = viper.GetString("token.refresh_token")
			log.GetLogger(verboseLevel).Debug("No new refresh token provided, keeping existing one")
		} else {
			log.GetLogger(verboseLevel).Debug("New refresh token provided, updating")
		}

		// Determine refresh token creation time
		// During token refresh, we should update the refresh token creation time
		// because the refresh operation represents a new use of the refresh token
		var refreshTokenCreatedAt time.Time
		hasNewRefreshToken := response.Result.RefreshToken != nil && getStringValue(response.Result.RefreshToken) != ""

		if hasNewRefreshToken {
			// New refresh token provided, update creation time
			refreshTokenCreatedAt = time.Now()
			log.GetLogger(verboseLevel).Debug("New refresh token provided, updating creation time to current time")
		} else {
			// No new refresh token, but we're doing a refresh operation
			// so we should update the creation time to reflect the new refresh operation
			refreshTokenCreatedAt = time.Now()
			log.GetLogger(verboseLevel).Debug("Token refresh operation completed, updating refresh token creation time to current time")
		}

		newTokenInfo = &TokenInfo{
			AccessToken:           getStringValue(response.Result.AccessToken),
			RefreshToken:          newRefreshToken,
			ExpiresIn:             response.Result.ExpireInSeconds,
			RefreshExpiresIn:      viper.GetInt32("token.refresh_expires_in"), // Keep existing refresh token expiry
			TokenCreatedAt:        time.Now(),
			RefreshTokenCreatedAt: refreshTokenCreatedAt,
		}

		// Debug: Log what we're about to save
		log.GetLogger(verboseLevel).Debug("New token info to save",
			"AccessToken", newTokenInfo.AccessToken,
			"RefreshToken", newTokenInfo.RefreshToken,
			"ExpiresIn", newTokenInfo.ExpiresIn,
			"RefreshTokenCreatedAt", newTokenInfo.RefreshTokenCreatedAt)

	default:
		return fmt.Errorf("token refresh not supported for API version: %T", clientVersion)
	}

	// Update configuration with new token information
	configPath := viper.ConfigFileUsed()
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		configPath = filepath.Join(home, ".cyverApiCli.yaml")
	}

	viper.SetConfigFile(configPath)
	_ = viper.ReadInConfig() // ignore error, file may not exist yet

	// Update token information in config
	viper.Set("token.access_token", newTokenInfo.AccessToken)
	viper.Set("token.refresh_token", newTokenInfo.RefreshToken)
	viper.Set("token.expireInSeconds", newTokenInfo.ExpiresIn)
	viper.Set("token.refresh_expires_in", newTokenInfo.RefreshExpiresIn)
	viper.Set("token.token_created_at", newTokenInfo.TokenCreatedAt.Format(time.RFC3339))
	viper.Set("token.refresh_token_created_at", newTokenInfo.RefreshTokenCreatedAt.Format(time.RFC3339))

	// Debug logging to verify what's being saved
	log.GetLogger(verboseLevel).Debug("Saving token information to config",
		"access_token_length", len(newTokenInfo.AccessToken),
		"expireInSeconds", newTokenInfo.ExpiresIn,
		"token_created_at", newTokenInfo.TokenCreatedAt.Format(time.RFC3339))

	// Write updated config
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to save refreshed token to config: %w", err)
	}

	log.GetLogger(verboseLevel).Info("Token refresh completed successfully")
	return nil
}

// ValidateAndRefreshToken validates the current token and refreshes it if necessary
func ValidateAndRefreshToken() error {
	log.GetLogger(verboseLevel).Debug("Starting token validation and refresh check")

	// Check if we have token information
	accessToken := viper.GetString("token.access_token")
	if accessToken == "" {
		log.GetLogger(verboseLevel).Debug("No access token found, skipping validation")
		return nil // No token to validate
	}

	log.GetLogger(verboseLevel).Debug("Access token found, checking expiration",
		"token_length", len(accessToken))

	// Check if token is expired
	isExpired, err := IsTokenExpired()
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error checking token expiration", "error", err)
		return fmt.Errorf("failed to check token expiration: %w", err)
	}

	log.GetLogger(verboseLevel).Debug("Token expiration check result", "is_expired", isExpired)

	if !isExpired {
		log.GetLogger(verboseLevel).Debug("Token is still valid, no refresh needed")
		return nil
	}

	log.GetLogger(verboseLevel).Info("Access token has expired, attempting to refresh")

	// Attempt to refresh the token
	if err := RefreshAccessToken(); err != nil {
		log.GetLogger(verboseLevel).Error("Token refresh failed", "error", err)
		return fmt.Errorf("token refresh failed: %w", err)
	}

	log.GetLogger(verboseLevel).Info("Token successfully refreshed")
	return nil
}

// ReAuthenticate re-authenticates using stored email from config and prompts for password
func ReAuthenticate() error {
	log.GetLogger(verboseLevel).Info("Starting re-authentication process")

	// Get stored email from config
	storedEmail := viper.GetString("auth.email")
	if storedEmail == "" {
		return fmt.Errorf("no stored email found in configuration. Please run 'cyverApiCli config init' first")
	}

	// Display stored email and prompt for password
	fmt.Printf("\nRe-authenticating for: %s\n", storedEmail)
	fmt.Println("======================")

	// Get password
	password, err := getPassword()
	if err != nil {
		return fmt.Errorf("failed to get password: %w", err)
	}

	// Get API configuration
	baseURL := viper.GetString("api.base_url")
	apiVersion := viper.GetString("api.version")

	if baseURL == "" || apiVersion == "" {
		return fmt.Errorf("missing API configuration (base_url or version). Please run 'cyverApiCli config init' first")
	}

	// Create API client for authentication
	client := versionedApiClient()
	if client == nil {
		return fmt.Errorf("failed to create API client")
	}

	// Perform authentication based on API version
	switch clientVersion := client.(type) {
	case *v2_2.Client:
		if clientVersion.TokenAuthOps == nil {
			return fmt.Errorf("TokenAuthOps is nil for v2.2 client")
		}

		// Set up initial authentication parameters
		params := v2_2.AuthenticateModel{
			UserNameOrEmailAddress: storedEmail,
			Password:               password,
			RememberClient:         true, // Remember client for future use
		}

		// Perform initial authentication
		response, err := clientVersion.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Authentication API call failed", "error", err)
			return fmt.Errorf("authentication failed: %w", err)
		}

		// Check if authentication was successful
		if !response.Success {
			errorMsg := "authentication failed"
			if response.Error != nil && response.Error.Message != nil {
				errorMsg = *response.Error.Message
			}
			log.GetLogger(verboseLevel).Error("Authentication failed", "error", errorMsg)
			return fmt.Errorf("authentication failed: %s", errorMsg)
		}

		// Check if 2FA is required
		if response.Result == nil || response.Result.RequiresTwoFactorVerification {
			log.GetLogger(verboseLevel).Info("Two-factor authentication required")

			// Get user ID for 2FA request
			userId := ""
			if response.Result != nil && response.Result.UserId != "" {
				userId = response.Result.UserId
			} else {
				// If no user ID in response, we need to get it another way
				// For now, we'll prompt the user to enter it
				fmt.Print("Enter your User ID for 2FA: ")
				fmt.Scanln(&userId)
				if userId == "" {
					return fmt.Errorf("user ID is required for 2FA")
				}
			}

			// Send 2FA code request
			twoFactorRequest := v2_2.SendTwoFactorAuthCodeModel{
				UserId:   userId,
				Provider: stringPtr("GoogleAuthenticator"), // Default to Google Authenticator
			}

			_, err = clientVersion.TokenAuthOps.ApiTokenauthSendtwofactorauthcodePost(twoFactorRequest)
			if err != nil {
				log.GetLogger(verboseLevel).Error("Failed to send 2FA code", "error", err)
				return fmt.Errorf("failed to send 2FA code: %w", err)
			}

			fmt.Println("2FA code sent to your registered device")
			fmt.Print("Enter the 2FA code: ")

			// Get 2FA code from user
			var twoFactorCode string
			fmt.Scanln(&twoFactorCode)
			if twoFactorCode == "" {
				return fmt.Errorf("2FA code is required")
			}

			// Complete authentication with 2FA code
			params.TwoFactorVerificationCode = &twoFactorCode
			response, err = clientVersion.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
			if err != nil {
				log.GetLogger(verboseLevel).Error("2FA authentication failed", "error", err)
				return fmt.Errorf("2FA authentication failed: %w", err)
			}

			if !response.Success {
				errorMsg := "2FA authentication failed"
				if response.Error != nil && response.Error.Message != nil {
					errorMsg = *response.Error.Message
				}
				log.GetLogger(verboseLevel).Error("2FA authentication failed", "error", errorMsg)
				return fmt.Errorf("2FA authentication failed: %s", errorMsg)
			}
		}

		if response.Result == nil {
			return fmt.Errorf("empty response from authentication")
		}

		// Update configuration with new token information
		configPath := viper.ConfigFileUsed()
		if configPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			configPath = filepath.Join(home, ".cyverApiCli.yaml")
		}

		viper.SetConfigFile(configPath)
		_ = viper.ReadInConfig() // ignore error, file may not exist yet

		// Store email for future re-authentication
		viper.Set("auth.email", storedEmail)

		// Handle pointer types safely and update token information
		if response.Result.AccessToken != nil {
			viper.Set("token.access_token", *response.Result.AccessToken)
		}
		if response.Result.RefreshToken != nil {
			viper.Set("token.refresh_token", *response.Result.RefreshToken)
		}
		viper.Set("token.expireInSeconds", response.Result.ExpireInSeconds)
		viper.Set("token.refresh_expires_in", response.Result.RefreshTokenExpireInSeconds)

		// Set token creation timestamps
		now := time.Now()
		viper.Set("token.token_created_at", now.Format(time.RFC3339))
		viper.Set("token.refresh_token_created_at", now.Format(time.RFC3339))

		// Write updated config
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save authentication data to config: %w", err)
		}

		log.GetLogger(verboseLevel).Info("Re-authentication completed successfully")
		fmt.Println("Re-authentication completed successfully!")
		return nil

	default:
		return fmt.Errorf("re-authentication not supported for API version: %T", clientVersion)
	}
}
