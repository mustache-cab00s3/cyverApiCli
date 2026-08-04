package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateConfigCmd = &cobra.Command{
	Use:   "update",
	Short: "Update values in the CLI config file",
	Long: `Update settings in the YAML config without re-running the full init wizard.

API and auth fields can be scoped to a named profile: use --for-profile, or rely on global --profile / CYVER_PROFILE / current_profile (legacy instances map in YAML still works).
Proxy, logging, output, and client settings are always global.

Examples:
  cyverApiCli config update --base-url https://api.cyver.io
  cyverApiCli config update --for-profile dev-user --base-url https://example.com
  cyverApiCli config update --log-level info --client-timeout 60`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := UpdateConfig(cmd); err != nil {
			fmt.Printf("Error updating config: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Configuration updated successfully.")
	},
}

func init() {
	updateConfigCmd.Flags().String("for-profile", "", "scope api and auth flags to this profile (default: global --profile / CYVER_PROFILE / current_profile)")
	updateConfigCmd.Flags().String("api-version", "", "API version (v2.2 or latest)")
	updateConfigCmd.Flags().String("base-url", "", "API base URL (https://...)")
	updateConfigCmd.Flags().String("api-key", "", "API key (use empty string with --api-key \"\" to clear)")
	updateConfigCmd.Flags().String("auth-email", "", "stored email for token re-authentication")
	updateConfigCmd.Flags().String("proxy-url", "", "HTTP proxy URL (empty clears proxy settings)")
	updateConfigCmd.Flags().String("proxy-user", "", "proxy username")
	updateConfigCmd.Flags().String("proxy-password", "", "proxy password")
	updateConfigCmd.Flags().String("log-level", "", "log level: debug, info, warn, error")
	updateConfigCmd.Flags().String("log-file", "", "log file path (empty for stdout)")
	updateConfigCmd.Flags().String("output-format", "", "default output: json, yaml, or table")
	updateConfigCmd.Flags().Bool("output-color", false, "enable colored output (use with --output-color=true or --output-color=false)")
	updateConfigCmd.Flags().Lookup("output-color").NoOptDefVal = "true"
	updateConfigCmd.Flags().Int("client-timeout", 0, "HTTP client timeout in seconds (1–300)")
	configCmd.AddCommand(updateConfigCmd)
}

func ensureConfigReadable() error {
	path := viper.ConfigFileUsed()
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("resolve home directory: %w", err)
		}
		path = filepath.Join(home, ".cyverApiCli.yaml")
		viper.SetConfigFile(path)
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found: %s (run cyverApiCli config init)", path)
		}
		return fmt.Errorf("config file: %w", err)
	}
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}
	return nil
}

func validateProfileTarget(name string) error {
	if name == "" {
		return nil
	}
	if !profileExists(name) {
		return fmt.Errorf("unknown profile %q (use cyverApiCli config profile list)", name)
	}
	return nil
}

func validateLogLevelFlexible(level string) bool {
	switch strings.ToLower(level) {
	case "debug", "info", "warn", "error":
		return true
	default:
		return false
	}
}

// UpdateConfig applies changed flags from the update command to viper and writes the config file.
func UpdateConfig(cmd *cobra.Command) error {
	if err := ensureConfigReadable(); err != nil {
		return err
	}
	if err := ResolveAndApplyProfile(cmd); err != nil {
		return err
	}

	target := EffectiveProfileForConfigUpdate(cmd)
	if err := validateProfileTarget(target); err != nil {
		return err
	}

	var n int

	if cmd.Flags().Changed("api-version") {
		v, _ := cmd.Flags().GetString("api-version")
		if !validateAPIVersion(v) {
			return fmt.Errorf("invalid API version %q (use v2.2 or latest)", v)
		}
		SetAPIFieldForProfile("version", v, target)
		n++
	}

	if cmd.Flags().Changed("base-url") {
		v, _ := cmd.Flags().GetString("base-url")
		if !validateURL(v) {
			return fmt.Errorf("invalid base URL %q (must start with http:// or https://)", v)
		}
		SetAPIFieldForProfile("base_url", v, target)
		n++
	}

	if cmd.Flags().Changed("api-key") {
		v, _ := cmd.Flags().GetString("api-key")
		if v != "" && !validateAPIKey(v) {
			return fmt.Errorf("invalid API key format")
		}
		SetAPIFieldForProfile("api_key", v, target)
		n++
	}

	if cmd.Flags().Changed("auth-email") {
		raw, _ := cmd.Flags().GetString("auth-email")
		v := strings.TrimSpace(raw)
		if v != "" && !validateUsername(v) {
			return fmt.Errorf("invalid auth-email (username or email)")
		}
		SetAuthFieldForProfile("email", v, target)
		n++
	}

	if cmd.Flags().Changed("proxy-url") {
		v, _ := cmd.Flags().GetString("proxy-url")
		if v == "" {
			viper.Set("proxy", map[string]interface{}{})
			n++
		} else {
			if !validateURL(v) {
				return fmt.Errorf("invalid proxy URL %q", v)
			}
			viper.Set("proxy.url", v)
			n++
		}
	}

	if cmd.Flags().Changed("proxy-user") {
		u, _ := cmd.Flags().GetString("proxy-user")
		viper.Set("proxy.username", u)
		n++
	}

	if cmd.Flags().Changed("proxy-password") {
		p, _ := cmd.Flags().GetString("proxy-password")
		viper.Set("proxy.password", p)
		n++
	}

	if cmd.Flags().Changed("log-level") {
		v, _ := cmd.Flags().GetString("log-level")
		if !validateLogLevelFlexible(v) {
			return fmt.Errorf("invalid log level %q (use debug, info, warn, error)", v)
		}
		viper.Set("logging.level", v)
		n++
	}

	if cmd.Flags().Changed("log-file") {
		lf, _ := cmd.Flags().GetString("log-file")
		viper.Set("logging.file", lf)
		n++
	}

	if cmd.Flags().Changed("output-format") {
		v, _ := cmd.Flags().GetString("output-format")
		if !validateOutputFormat(v) {
			return fmt.Errorf("invalid output format %q (use json, yaml, table)", v)
		}
		viper.Set("output.format", v)
		n++
	}

	if cmd.Flags().Changed("output-color") {
		col, err := cmd.Flags().GetBool("output-color")
		if err != nil {
			return err
		}
		viper.Set("output.color", col)
		n++
	}

	if cmd.Flags().Changed("client-timeout") {
		t, _ := cmd.Flags().GetInt("client-timeout")
		if t < 1 || t > 300 {
			return fmt.Errorf("client-timeout must be between 1 and 300")
		}
		viper.Set("client.timeout", t)
		n++
	}

	if n == 0 {
		return fmt.Errorf("no updates: specify at least one flag (see cyverApiCli config update --help)")
	}

	configPath := viper.ConfigFileUsed()
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("resolve config path: %w", err)
		}
		configPath = filepath.Join(home, ".cyverApiCli.yaml")
		viper.SetConfigFile(configPath)
	}

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	if err := os.Chmod(configPath, 0600); err != nil {
		return fmt.Errorf("set config permissions: %w", err)
	}

	return nil
}
