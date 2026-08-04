package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/cmd/shared"
	v2_2 "github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	log "github.com/yourusername/cyverApiCli/logger"
	"golang.org/x/term"
)

// Prompt reads input from the user with optional masking
func Prompt(message string, mask bool) (string, error) {
	fmt.Print(message)
	var input strings.Builder

	if mask {
		// Set terminal to raw mode (no echo)
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			log.GetLogger(verboseLevel).Error("failed to set raw mode", "error", err)
			return "", err
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		// Read one byte at a time
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.GetLogger(verboseLevel).Error("failed to read input", "error", err)
				return "", err
			}

			char := buf[0]

			// Enter key
			if char == '\r' || char == '\n' {
				fmt.Println()
				break
			}

			// Backspace (ASCII 127 or 8)
			if char == 127 || char == 8 {
				if input.Len() > 0 {
					// Remove last character by rebuilding the string
					current := input.String()
					input.Reset()
					input.WriteString(current[:len(current)-1])

					// Clear line and reprint masked input
					fmt.Print("\r\033[K")
					fmt.Print(message + strings.Repeat("*", input.Len()))
				}
				continue
			}

			// Add character to input
			input.WriteByte(char)
			fmt.Print("*")
		}
	} else {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			log.GetLogger(verboseLevel).Error("failed to read input", "error", err)
			return "", err
		}
		input.WriteString(line)
	}

	return input.String(), nil
}

// getUsername retrieves the username or email from command flags
func getUsername(cmd *cobra.Command) (string, error) {
	username, err := cmd.Flags().GetString("username")
	if err != nil || len(username) == 0 {
		log.GetLogger(verboseLevel).Error("username is required", "error", err)
		return "", fmt.Errorf("username is required")
	}
	log.GetLogger(verboseLevel).Info("Set User as", "user", username)
	return username, nil
}

// getPassword prompts for and retrieves the password
func getPassword() (string, error) {
	password, err := Prompt("Enter password: ", true)
	if err != nil {
		log.GetLogger(verboseLevel).Error("failed to read password", "error", err)
		return "", err
	}
	log.GetLogger(verboseLevel).Info("Set password as", "password", password)
	return password, nil
}

// getOptionalAuthParams retrieves optional authentication parameters from command flags
func getOptionalAuthParams(cmd *cobra.Command) (v2_2.AuthenticateModel, error) {
	params := v2_2.AuthenticateModel{}

	// Retrieve optional flags
	if twoFactorCode, err := cmd.Flags().GetString("two-factor-code"); err == nil && twoFactorCode != "" {
		params.TwoFactorVerificationCode = &twoFactorCode
	}
	if rememberClient, err := cmd.Flags().GetBool("remember-client"); err == nil {
		params.RememberClient = rememberClient
	}
	if twoFactorToken, err := cmd.Flags().GetString("two-factor-token"); err == nil && twoFactorToken != "" {
		params.TwoFactorRememberClientToken = &twoFactorToken
	}
	if singleSignIn, err := cmd.Flags().GetBool("single-sign-in"); err == nil {
		params.SingleSignIn = &singleSignIn
	}
	if returnUrl, err := cmd.Flags().GetString("return-url"); err == nil && returnUrl != "" {
		params.ReturnUrl = &returnUrl
	}
	if captchaResponse, err := cmd.Flags().GetString("captcha-response"); err == nil && captchaResponse != "" {
		params.CaptchaResponse = &captchaResponse
	}

	return params, nil
}

// setupV2_2Parameters sets up authentication parameters for v2.2 client
func setupV2_2Parameters(cmd *cobra.Command) (v2_2.AuthenticateModel, error) {
	params := v2_2.AuthenticateModel{}

	// Get required parameters
	username, err := getUsername(cmd)
	if err != nil {
		return params, err
	}
	params.UserNameOrEmailAddress = username

	password, err := getPassword()
	if err != nil {
		return params, err
	}
	params.Password = password

	// Get optional parameters
	optionalParams, err := getOptionalAuthParams(cmd)
	if err != nil {
		return params, err
	}

	// Merge optional parameters
	params.TwoFactorVerificationCode = optionalParams.TwoFactorVerificationCode
	params.RememberClient = optionalParams.RememberClient
	params.TwoFactorRememberClientToken = optionalParams.TwoFactorRememberClientToken
	params.SingleSignIn = optionalParams.SingleSignIn
	params.ReturnUrl = optionalParams.ReturnUrl
	params.CaptchaResponse = optionalParams.CaptchaResponse

	return params, nil
}

// StoreToken saves the token response to a file with 0600 permissions
func StoreToken(tokenResp *v2_2.AuthenticateResultModel, filePath string) error {
	// Convert token response to JSON
	tokenData, err := json.MarshalIndent(tokenResp, "", "  ")
	if err != nil {
		return err
	}

	// Write to file with 0600 permissions
	// Using 0600 ensures only the owner can read/write
	err = os.WriteFile(filePath, tokenData, fs.FileMode(0600))
	if err != nil {
		return err
	}

	return nil
}

// authenticateV2_2 handles the v2.2 authentication process
func authenticateV2_2(client *v2_2.Client, params v2_2.AuthenticateModel) (*v2_2.AuthenticateResultModel, error) {
	// Perform initial authentication
	response, err := client.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error during authentication", "error", err)
		return nil, err
	}

	// Check if authentication was successful
	if !response.Success {
		errorMsg := "authentication failed"
		if response.Error != nil && response.Error.Message != nil {
			errorMsg = *response.Error.Message
		}
		log.GetLogger(verboseLevel).Error("Authentication failed", "error", errorMsg)
		return nil, fmt.Errorf("authentication failed: %s", errorMsg)
	}

	if response.Result == nil {
		return nil, fmt.Errorf("authentication failed: no result in response")
	}

	// Print the response
	if err := shared.PrintJSONResponse(response); err != nil {
		log.GetLogger(verboseLevel).Error("Error printing authentication response", "error", err)
		return nil, err
	}

	// MFA only when the API explicitly requires it (do not treat nil/empty result as MFA — fixes non-MFA sandboxes)
	if response.Result.RequiresTwoFactorVerification {
		log.GetLogger(verboseLevel).Info("Two-factor authentication required")

		// Get user ID for 2FA request
		userId := response.Result.UserId
		if userId == "" {
			return nil, fmt.Errorf("user ID is required for 2FA but was not provided in response")
		}

		provider := twoFactorProviderForSend(response.Result.TwoFactorAuthProviders)
		log.GetLogger(verboseLevel).Info("Using two-factor provider from API", "provider", provider)

		// Send 2FA code request
		twoFactorRequest := v2_2.SendTwoFactorAuthCodeModel{
			UserId:   userId,
			Provider: stringPtr(provider),
		}

		_, err = client.TokenAuthOps.ApiTokenauthSendtwofactorauthcodePost(twoFactorRequest)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Failed to send 2FA code", "error", err)
			return nil, fmt.Errorf("failed to send 2FA code: %w", err)
		}

		fmt.Println("2FA code sent to your registered device")
		twoFactorCode, err := Prompt("Enter your 2FA code: ", false)
		if err != nil {
			log.GetLogger(verboseLevel).Error("TwoFactorVerificationCode Error", "error", err)
			return nil, err
		}
		if twoFactorCode == "" {
			return nil, fmt.Errorf("2FA code is required")
		}
		params.TwoFactorVerificationCode = &twoFactorCode
		log.GetLogger(verboseLevel).Info("You entered", "TwoFactorVerificationCode", twoFactorCode)

		// Complete authentication with 2FA code
		response, err = client.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
		if err != nil {
			log.GetLogger(verboseLevel).Error("2FA authentication failed", "error", err)
			return nil, fmt.Errorf("2FA authentication failed: %w", err)
		}

		if !response.Success {
			errorMsg := "2FA authentication failed"
			if response.Error != nil && response.Error.Message != nil {
				errorMsg = *response.Error.Message
			}
			log.GetLogger(verboseLevel).Error("2FA authentication failed", "error", errorMsg)
			return nil, fmt.Errorf("2FA authentication failed: %s", errorMsg)
		}

		if response.Result == nil {
			return nil, fmt.Errorf("authentication failed: no result in response after 2FA")
		}
	}

	// Authentication successful (with or without 2FA)
	log.GetLogger(verboseLevel).Info("Authentication successful!")
	if response.Result.AccessToken != nil {
		log.GetLogger(verboseLevel).Info("Token: ", "token", *response.Result.AccessToken)
	}
	log.GetLogger(verboseLevel).Debug("Successfully retrieved token from v2.2 of the API")
	return response.Result, nil
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}

// twoFactorProviderForSend picks the provider for SendTwoFactorAuthCode from the API response.
// When the server does not require 2FA or lists no providers, this is only called from the MFA branch;
// if no providers are listed, Google Authenticator is used as the default (legacy behavior).
func twoFactorProviderForSend(providers []string) string {
	if len(providers) > 0 {
		return providers[0]
	}
	return v2_2.TwoFactorMethodAuthenticator
}

// handleClientSwitch processes the client version and performs authentication
func handleClientSwitch(clientVersion interface{}, cmd *cobra.Command) error {
	switch client := clientVersion.(type) {
	case *v2_2.Client:
		if client.TokenAuthOps == nil {
			log.GetLogger(verboseLevel).Error("Error: PentesterOps is nil for v2.2 client")
			return fmt.Errorf("PentesterOps is nil for v2.2 client")
		}

		// Set up parameters
		params, err := setupV2_2Parameters(cmd)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Error setting up parameters", "error", err)
			return err
		}

		// Perform authentication
		tokenAuth, err := authenticateV2_2(client, params)
		if err != nil {
			log.GetLogger(verboseLevel).Error("Error during authentication", "error", err)
			return err
		}

		// Store token in config file
		home, err := os.UserHomeDir()
		if err != nil {
			log.GetLogger(verboseLevel).Error("Error during getting home directory", "error", err)
			return err
		}

		// Write token data to config file using viper
		configPath := filepath.Join(home, ".cyverApiCli.yaml")
		viper.SetConfigFile(configPath)
		// Read config to avoid overwriting
		_ = viper.ReadInConfig() // ignore error, file may not exist yet

		// Handle pointer types safely
		if tokenAuth.AccessToken != nil {
			SetTokenViperKey("access_token", *tokenAuth.AccessToken)
		}
		if tokenAuth.RefreshToken != nil {
			SetTokenViperKey("refresh_token", *tokenAuth.RefreshToken)
		}
		SetTokenViperKey("expireInSeconds", tokenAuth.ExpireInSeconds)
		SetTokenViperKey("refresh_expires_in", tokenAuth.RefreshTokenExpireInSeconds)

		// Set token creation timestamps
		now := time.Now()
		SetTokenViperKey("token_created_at", now.Format(time.RFC3339))
		SetTokenViperKey("refresh_token_created_at", now.Format(time.RFC3339))
		if err := viper.WriteConfig(); err != nil {
			log.GetLogger(verboseLevel).Error("Error writing token to config file", "error", err)
			return err
		}

	default:
		log.GetLogger(verboseLevel).Error("Error: unsupported client type: %T", clientVersion)
		return fmt.Errorf("unsupported client type: %T", clientVersion)
	}
	return nil
}

var apiAuthCmd = &cobra.Command{
	Use:   "apiAuth",
	Short: "Authenticate to the API",
	Long:  `Perform the Authentication process with the api. Username/Email and Password with MFA may be required`,
}

var tokenAuthCmd = &cobra.Command{
	Use:   "getToken",
	Short: "Get API Authentication token",
	Long:  `Perform the Token Authentication Process.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientVersion := versionedApiClient()
		if clientVersion == nil {
			log.GetLogger(verboseLevel).Error("Error: failed to initialize API client")
			os.Exit(1)
		}

		// Handle client version
		if err := handleClientSwitch(clientVersion, cmd); err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// Define flags with correct types
	tokenAuthCmd.Flags().StringP("username", "u", "", "Username or email address (required)")
	tokenAuthCmd.Flags().String("two-factor-code", "", "Two-factor verification code")
	tokenAuthCmd.Flags().Bool("remember-client", true, "Remember the client")
	tokenAuthCmd.Flags().String("two-factor-token", "", "Two-factor remember client token")
	tokenAuthCmd.Flags().Bool("single-sign-in", false, "Enable single sign-in")
	tokenAuthCmd.Flags().String("return-url", "", "Return URL after authentication")
	tokenAuthCmd.Flags().String("captcha-response", "", "CAPTCHA response")

	// Mark required flags
	tokenAuthCmd.MarkFlagRequired("username")

	apiAuthCmd.AddCommand(tokenAuthCmd)
	rootCmd.AddCommand(apiAuthCmd)
}
