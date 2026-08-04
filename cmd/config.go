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
	Long:  `Create or extend the configuration file. If the file already exists, you can overwrite it, add a named profile without losing other settings, or cancel. You can store credentials under named profiles (AWS-style) and set current_profile.`,
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

var profileCmd = &cobra.Command{
	Use:     "profile",
	Aliases: []string{"instance"},
	Short:   "List or select named configuration profiles",
	Long: `Named profiles (similar to AWS ~/.aws/config) keep separate api, auth, and token data per profile.

Use the "profiles" map in YAML (legacy "instances" is still read). The default profile name is stored in current_profile.
Set the default with "config profile use <name>", override per command with --profile / -p, or set CYVER_PROFILE.`,
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured profile names",
	Run: func(cmd *cobra.Command, args []string) {
		aliases := ListProfileNames()
		if len(aliases) == 0 {
			fmt.Println("No profiles configured. Add a \"profiles\" map to your config file (each key is a profile name).")
			return
		}
		def := EffectiveCurrentProfileName()
		fmt.Println("Configured profiles:")
		for _, a := range aliases {
			mark := ""
			if def != "" && a == def {
				mark = " (default)"
			}
			fmt.Printf("  %s%s\n", a, mark)
		}
	},
}

var profileUseCmd = &cobra.Command{
	Use:   "use PROFILE",
	Short: "Set the default profile (saved as current_profile)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		if !profileExists(alias) {
			fmt.Printf("Unknown profile %q. Run \"cyverApiCli config profile list\" to see names.\n", alias)
			os.Exit(1)
		}
		viper.Set("current_profile", alias)
		viper.Set("current_instance", "")
		configPath := viper.ConfigFileUsed()
		if configPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("Error resolving config path: %v\n", err)
				os.Exit(1)
			}
			configPath = filepath.Join(home, ".cyverApiCli.yaml")
			viper.SetConfigFile(configPath)
		}
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Default profile set to %q (current_profile).\n", alias)
	},
}

func init() {
	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(refreshTokenCmd)
	configCmd.AddCommand(reAuthCmd)
	configCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileUseCmd)
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

// validateProfileName checks a safe profile name (like AWS [profile name]: letters, digits, hyphens, underscores).
func validateProfileName(s string) bool {
	if len(s) < 1 || len(s) > 40 {
		return false
	}
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
		case r >= 'A' && r <= 'Z':
		case r >= '0' && r <= '9':
		case r == '-' || r == '_':
		default:
			return false
		}
	}
	return true
}

// promptNamedProfile asks whether to use named profiles (AWS-style) and returns the chosen name.
func promptNamedProfile() (named bool, name string) {
	var useStr string
	fmt.Println("\nNamed profile (AWS-style)")
	fmt.Println("-------------------------")
	fmt.Println("Named profiles keep separate API URLs and tokens (like [default] and [profile dev-user] in AWS config).")
	fmt.Println("You can add more profiles later by editing the config file or running config init again.")
	for {
		fmt.Print("\nStore this API connection under a named profile? (y/N): ")
		fmt.Scanln(&useStr)
		if useStr == "" || useStr == "n" || useStr == "N" {
			return false, ""
		}
		if useStr == "y" || useStr == "Y" {
			break
		}
		fmt.Println("Please answer 'y' or 'n'")
	}
	for {
		fmt.Print("Profile name [default]: ")
		fmt.Scanln(&name)
		if name == "" {
			name = "default"
		}
		if validateProfileName(name) {
			return true, name
		}
		fmt.Println("Invalid name. Use 1–40 characters: letters, digits, hyphens, underscores.")
	}
}

// applyProfileLayout sets top-level api and, when named, current_profile and profiles.<name>.api.
func applyProfileLayout(config map[string]interface{}, named bool, name string, apiBlock map[string]interface{}) {
	config["api"] = apiBlock
	if !named || name == "" {
		return
	}
	config["current_profile"] = name
	config["profiles"] = map[string]interface{}{
		name: map[string]interface{}{
			"api": apiBlock,
		},
	}
}

// clearProfileKeysIfFlatProfile removes profiles/instances and current_* keys when saving a flat single-profile config.
func clearProfileKeysIfFlatProfile(named bool) {
	if named {
		return
	}
	viper.Set("profiles", map[string]interface{}{})
	viper.Set("instances", map[string]interface{}{})
	viper.Set("current_profile", "")
	viper.Set("current_instance", "")
}

// collectCoreAPISettings prompts for API version, base URL, and optional API key (shared by full init and add-profile).
func collectCoreAPISettings() (apiVersion, baseURL, apiKey string) {
	fmt.Println("\nCyver API Configuration")
	fmt.Println("======================")

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

	for {
		fmt.Print("\nEnter your API key (optional, press Enter to skip): ")
		fmt.Scanln(&apiKey)
		if apiKey == "" || validateAPIKey(apiKey) {
			break
		}
		fmt.Println("Invalid API key. API key must be at least 32 characters long.")
	}
	return apiVersion, baseURL, apiKey
}

// promptNewProfileNameForAdd asks for a profile name; if it already exists, offers to replace the API block only.
func promptNewProfileNameForAdd() string {
	for {
		fmt.Print("\nProfile name (letters, digits, hyphens, underscores) [dev-user]: ")
		var name string
		fmt.Scanln(&name)
		if name == "" {
			name = "dev-user"
		}
		if !validateProfileName(name) {
			fmt.Println("Invalid name. Use 1–40 characters: letters, digits, hyphens, underscores.")
			continue
		}
		if profileExists(name) {
			fmt.Printf("Profile %q already exists. Replace its stored API settings (auth/token are kept unless you edit the file)? (y/N): ", name)
			var r string
			fmt.Scanln(&r)
			if r != "y" && r != "Y" {
				continue
			}
		}
		return name
	}
}

// initializeConfigAddProfile merges a new or updated profile into the existing YAML without wiping other keys.
func initializeConfigAddProfile(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read existing config: %w", err)
	}

	fmt.Println("\nAdd or update a profile")
	fmt.Println("-------------------------")
	fmt.Println("Global settings (proxy, logging, output, client timeout) are unchanged. You can edit them with \"cyverApiCli config update\".")

	profileName := promptNewProfileNameForAdd()
	apiVersion, baseURL, apiKey := collectCoreAPISettings()

	apiBlock := map[string]interface{}{
		"version":  apiVersion,
		"base_url": baseURL,
		"api_key":  apiKey,
	}

	var setDefault string
	for {
		fmt.Print("\nSet this profile as the default (current_profile)? (y/N): ")
		fmt.Scanln(&setDefault)
		if setDefault == "" || setDefault == "n" || setDefault == "N" {
			setDefault = "n"
			break
		}
		if setDefault == "y" || setDefault == "Y" {
			setDefault = "y"
			break
		}
		fmt.Println("Please answer 'y' or 'n'")
	}

	rootKey := MergeProfileStorageRootForWrite()
	existing := toStringMap(viper.Get(rootKey))
	if existing == nil {
		existing = map[string]interface{}{}
	}

	prev := toStringMap(existing[profileName])
	entry := map[string]interface{}{"api": apiBlock}
	if prev != nil {
		if a, ok := prev["auth"]; ok {
			entry["auth"] = a
		}
		if t, ok := prev["token"]; ok {
			entry["token"] = t
		}
	}
	existing[profileName] = entry
	viper.Set(rootKey, existing)

	activeProfileResolved = profileName
	if err := ApplyProfile(profileName); err != nil {
		return fmt.Errorf("apply profile: %w", err)
	}
	if setDefault == "y" {
		viper.Set("current_profile", profileName)
		viper.Set("current_instance", "")
	}

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	if err := os.Chmod(configPath, 0600); err != nil {
		return fmt.Errorf("failed to set config file permissions: %w", err)
	}

	fmt.Printf("\nProfile %q saved under %q in %s\n", profileName, rootKey, configPath)
	if setDefault == "y" {
		fmt.Println("This profile is now the default (current_profile).")
	} else {
		fmt.Printf("Default profile is unchanged. Use: cyverApiCli --profile %s ... or: cyverApiCli config profile use %s\n", profileName, profileName)
	}
	fmt.Println("\nYou can now use the CLI with your configured settings.")
	fmt.Println("To view your current configuration, use 'cyverApiCli config view'")

	runInteractiveTokenAuthIfNoAPIKey(apiKey, "\nNo API key provided for this profile. Starting token authentication...")

	return nil
}

// runInteractiveTokenAuthIfNoAPIKey prompts for username and runs token authentication when init completed without an API key.
func runInteractiveTokenAuthIfNoAPIKey(apiKey, introLine string) {
	if apiKey != "" {
		return
	}
	fmt.Println(introLine)
	var username string
	for {
		fmt.Print("Enter your username or email address: ")
		fmt.Scanln(&username)
		if validateUsername(username) {
			break
		}
		fmt.Println("Please enter a valid username (3-50 characters, alphanumeric with dots, underscores, hyphens) or email address.")
	}
	SetAuthViperKey("email", username)
	tempCmd := &cobra.Command{}
	tempCmd.Flags().String("username", username, "")
	if err := handleClientSwitch(versionedApiClient(), tempCmd); err != nil {
		fmt.Printf("Token authentication failed: %v\n", err)
		fmt.Println("You can manually run 'cyverApiCli apiAuth getToken -u <username>' to authenticate later.")
	} else {
		fmt.Println("Token authentication completed successfully!")
	}
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
		fmt.Println("Choose how to proceed:")
		fmt.Println("  O — Overwrite the entire file (full guided setup; replaces all settings)")
		fmt.Println("  A — Add or replace one named profile (keeps proxy, logging, and other profiles)")
		fmt.Println("  C — Cancel")
		fmt.Print("Your choice [C]: ")
		var response string
		fmt.Scanln(&response)
		response = strings.TrimSpace(strings.ToUpper(response))
		if response == "A" {
			return initializeConfigAddProfile(configPath)
		}
		if response != "O" && response != "Y" {
			fmt.Println("Configuration initialization cancelled.")
			return nil
		}
	}

	apiVersion, baseURL, apiKey := collectCoreAPISettings()

	useNamed, profileName := promptNamedProfile()

	apiBlock := map[string]interface{}{
		"version":  apiVersion,
		"base_url": baseURL,
		"api_key":  apiKey,
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
			applyProfileLayout(config, useNamed, profileName, apiBlock)

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
			clearProfileKeysIfFlatProfile(useNamed)
			for key, value := range config {
				viper.Set(key, value)
			}
			if useNamed {
				activeProfileResolved = profileName
			}
			if err := viper.WriteConfig(); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}

			// Set file permissions to user-only (600)
			if err := os.Chmod(configPath, 0600); err != nil {
				return fmt.Errorf("failed to set config file permissions: %w", err)
			}

			fmt.Printf("\nConfiguration saved to %s\n", configPath)
			if useNamed {
				fmt.Printf("Named profile %q is set as default (current_profile). Use \"cyverApiCli config profile list\" to see names.\n", profileName)
			}
			fmt.Println("\nYou can now use the CLI with your configured settings.")
			fmt.Println("To modify these settings later, edit the config file or use 'cyverApiCli config init' again.")
			fmt.Println("To view your current configuration, use 'cyverApiCli config view'")

			runInteractiveTokenAuthIfNoAPIKey(apiKey, "\nNo API key provided. Starting token authentication process...")
			return nil
		}
		fmt.Println("Please answer with 'y' or 'n'")
	}

	// If no proxy configuration, write basic config
	config := map[string]interface{}{
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
	applyProfileLayout(config, useNamed, profileName, apiBlock)

	// Write config to file
	viper.SetConfigFile(configPath)
	clearProfileKeysIfFlatProfile(useNamed)
	for key, value := range config {
		viper.Set(key, value)
	}
	if useNamed {
		activeProfileResolved = profileName
	}
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Set file permissions to user-only (600)
	if err := os.Chmod(configPath, 0600); err != nil {
		return fmt.Errorf("failed to set config file permissions: %w", err)
	}

	fmt.Printf("\nConfiguration saved to %s\n", configPath)
	if useNamed {
		fmt.Printf("Named profile %q is set as default (current_profile). Use \"cyverApiCli config profile list\" to see names.\n", profileName)
	}
	fmt.Println("\nYou can now use the CLI with your configured settings.")
	fmt.Println("To modify these settings later, edit the config file or use 'cyverApiCli config init' again.")
	fmt.Println("To view your current configuration, use 'cyverApiCli config view'")

	runInteractiveTokenAuthIfNoAPIKey(apiKey, "\nNo API key provided. Starting token authentication process...")

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

	if aliases := ListProfileNames(); len(aliases) > 0 {
		fmt.Println("\nProfiles (AWS-style):")
		def := EffectiveCurrentProfileName()
		if def == "" {
			def = "(not set)"
		}
		fmt.Printf("  Default (current_profile / legacy current_instance): %s\n", def)
		if a := ActiveProfileName(); a != "" {
			fmt.Printf("  Active for this process: %s\n", a)
		}
		fmt.Printf("  Configured profile names: %s\n", strings.Join(aliases, ", "))
	}

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
	SetTokenViperKey("access_token", newTokenInfo.AccessToken)
	SetTokenViperKey("refresh_token", newTokenInfo.RefreshToken)
	SetTokenViperKey("expireInSeconds", newTokenInfo.ExpiresIn)
	SetTokenViperKey("refresh_expires_in", newTokenInfo.RefreshExpiresIn)
	SetTokenViperKey("token_created_at", newTokenInfo.TokenCreatedAt.Format(time.RFC3339))
	SetTokenViperKey("refresh_token_created_at", newTokenInfo.RefreshTokenCreatedAt.Format(time.RFC3339))

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

		if response.Result == nil {
			return fmt.Errorf("empty response from authentication")
		}

		// MFA only when the API sets requiresTwoFactorVerification (password-only / sandbox skips this)
		if response.Result.RequiresTwoFactorVerification {
			log.GetLogger(verboseLevel).Info("Two-factor authentication required")

			userId := response.Result.UserId
			if userId == "" {
				fmt.Print("Enter your User ID for 2FA: ")
				fmt.Scanln(&userId)
				if userId == "" {
					return fmt.Errorf("user ID is required for 2FA")
				}
			}

			provider := twoFactorProviderForSend(response.Result.TwoFactorAuthProviders)
			log.GetLogger(verboseLevel).Info("Using two-factor provider from API", "provider", provider)

			twoFactorRequest := v2_2.SendTwoFactorAuthCodeModel{
				UserId:   userId,
				Provider: stringPtr(provider),
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
		SetAuthViperKey("email", storedEmail)

		// Handle pointer types safely and update token information
		if response.Result.AccessToken != nil {
			SetTokenViperKey("access_token", *response.Result.AccessToken)
		}
		if response.Result.RefreshToken != nil {
			SetTokenViperKey("refresh_token", *response.Result.RefreshToken)
		}
		SetTokenViperKey("expireInSeconds", response.Result.ExpireInSeconds)
		SetTokenViperKey("refresh_expires_in", response.Result.RefreshTokenExpireInSeconds)

		// Set token creation timestamps
		now := time.Now()
		SetTokenViperKey("token_created_at", now.Format(time.RFC3339))
		SetTokenViperKey("refresh_token_created_at", now.Format(time.RFC3339))

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
