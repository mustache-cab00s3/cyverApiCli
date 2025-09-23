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
	// Get User ID using the new API function
	userIdResponse, err := client.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error during getting UserId", "error", err)
		return nil, err
	}
	if err := shared.PrintJSONResponse(userIdResponse); err != nil {
		log.GetLogger(verboseLevel).Error("Error printing UserId response", "error", err)
		return nil, err
	}

	// Extract user ID from response
	var userId string
	if userIdResponse.Result != nil && userIdResponse.Result.UserId != "" {
		userId = userIdResponse.Result.UserId
	} else {
		return nil, fmt.Errorf("failed to get user ID from authentication response")
	}

	// Send 2FA code using the new API function
	twoFactorRequest := v2_2.SendTwoFactorAuthCodeModel{
		UserId:   userId,
		Provider: stringPtr("GoogleAuthenticator"),
	}
	_, err = client.TokenAuthOps.ApiTokenauthSendtwofactorauthcodePost(twoFactorRequest)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error during sending 2FA code", "error", err)
		return nil, err
	}

	// Prompt for 2FA code
	twoFactorCode, err := Prompt("Enter your 2fa Code: ", false)
	if err != nil {
		log.GetLogger(verboseLevel).Error("TwoFactorVerificationCode Error", "error", err)
		return nil, err
	}
	params.TwoFactorVerificationCode = &twoFactorCode
	log.GetLogger(verboseLevel).Info("You entered", "TwoFactorVerificationCode", twoFactorCode)

	// Authenticate and get token using the new API function
	tokenResponse, err := client.TokenAuthOps.ApiTokenauthAuthenticatePost(params)
	if err != nil {
		log.GetLogger(verboseLevel).Error("Error during authentication", "error", err)
		return nil, err
	}

	if tokenResponse.Result == nil {
		return nil, fmt.Errorf("authentication failed: no result in response")
	}

	log.GetLogger(verboseLevel).Info("Authentication successful!")
	if tokenResponse.Result.AccessToken != nil {
		log.GetLogger(verboseLevel).Info("Token: ", "token", *tokenResponse.Result.AccessToken)
	}
	log.GetLogger(verboseLevel).Debug("Successfully retrieved token from v2.2 of the API")
	return tokenResponse.Result, nil
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
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
			viper.Set("token.access_token", *tokenAuth.AccessToken)
		}
		if tokenAuth.RefreshToken != nil {
			viper.Set("token.refresh_token", *tokenAuth.RefreshToken)
		}
		viper.Set("token.expireInSeconds", tokenAuth.ExpireInSeconds)
		viper.Set("token.refresh_expires_in", tokenAuth.RefreshTokenExpireInSeconds)

		// Set token creation timestamps
		now := time.Now()
		viper.Set("token.token_created_at", now.Format(time.RFC3339))
		viper.Set("token.refresh_token_created_at", now.Format(time.RFC3339))
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
