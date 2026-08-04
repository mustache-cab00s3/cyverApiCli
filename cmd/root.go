package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/cmd/shared"
	"github.com/yourusername/cyverApiCli/internal/api"
	"github.com/yourusername/cyverApiCli/internal/api/services"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	log "github.com/yourusername/cyverApiCli/logger"
)

var (
	verboseLevel int
	logger       = log.GetLogger(0) // initialize with default verbosity 0
)

// IsVerbose returns true if the specified verbosity level is enabled
func IsVerbose(level int) bool {
	return verboseLevel >= level
}

var rootCmd = &cobra.Command{
	Use:   "cyverApiCli",
	Short: "A CLI tool for interacting with the Cyver API",
	Long: `A command line interface tool that provides easy access to the Cyver API.
It allows you to interact with various endpoints and manage your resources.

Verbosity levels:
  -v    Show basic request information (method, URL)
  -vv   Show request headers and basic response info
  -vvv  Show full request/response details including body`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set verbosity and reinitialize the logger
		verboseLevel, _ = cmd.Flags().GetCount("verbose")
		logger = log.GetLogger(verboseLevel)
		logger.Info("Starting to execute command", "verbosity", verboseLevel)

		if err := ResolveAndApplyProfile(cmd); err != nil {
			logger.Error("Invalid API profile", "error", err)
			os.Exit(1)
		}

		// Update verbosity in shared package and API client
		shared.SetVerboseLevel(verboseLevel)
		api.SetVerboseLevel(verboseLevel)
		services.SetVerboseLevel(verboseLevel)
		v2_2.SetVerboseLevel(verboseLevel)

		// Handle manual continue flag (only works with -vvv)
		manualContinue, _ := cmd.Flags().GetBool("manual-continue")
		if manualContinue && verboseLevel >= 3 {
			api.SetManualContinue(true)
			logger.Info("Manual continue mode enabled (pauses after each API request for validation)")
		} else if manualContinue && verboseLevel < 3 {
			logger.Warn("--manual-continue requires -vvv verbosity level; ignoring flag")
		}
	},
}

func Execute() {
	logger.Info("Starting Execution...")
	if err := rootCmd.Execute(); err != nil {
		logger.Error("Failed to execute command", "error", err)
		os.Exit(1)
	}
}

func init() {
	// Add config file flag
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.cyverApiCli.yaml)")

	// Add verbose flag that can be specified multiple times
	rootCmd.PersistentFlags().CountP("verbose", "v", "increase verbosity level")

	// Add manual continue flag (requires -vvv)
	rootCmd.PersistentFlags().Bool("manual-continue", false, "pause after each API request for validation (requires -vvv)")

	// AWS-style profile selection (like AWS_PROFILE / ~/.aws/config)
	rootCmd.PersistentFlags().StringP("profile", "p", "", "use this named profile (overrides CYVER_PROFILE / current_profile)")

	// Initialize viper for configuration
	configPath := ""
	if configFile, _ := rootCmd.PersistentFlags().GetString("config"); configFile != "" {
		configPath = configFile
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			logger.Error("Failed to get home directory", "error", err)
			os.Exit(1)
		}
		configPath = filepath.Join(home, ".cyverApiCli.yaml")
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	_ = viper.BindEnv("current_profile", "CYVER_PROFILE")
	_ = viper.BindEnv("current_instance", "CYVER_INSTANCE")

	// Log config initialization
	logger.Info("Initializing config", "path", configPath)

	// Read the config file (if it exists)
	if err := viper.ReadInConfig(); err != nil {
		if os.IsNotExist(err) {
			logger.Info("Config file does not exist", "path", configPath)
		} else {
			logger.Error("Failed to read config file", "path", configPath, "error", err)
			os.Exit(1)
		}
	} else {
		logger.Info("Config file loaded successfully", "path", configPath)
	}
}
