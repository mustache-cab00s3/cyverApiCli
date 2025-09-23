package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourusername/cyverApiCli/cmd/shared"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	"github.com/yourusername/cyverApiCli/internal/errors"
)

// Users commands
var getUsersCmd = &cobra.Command{
	Use:   "list",
	Short: "Get users",
	Long:  `Retrieve a list of users with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		maxResultCount, _ := cmd.Flags().GetInt("max-results")
		skipCount, _ := cmd.Flags().GetInt("skip-count")
		filter, _ := cmd.Flags().GetString("filter")

		clientVersion := shared.GetVersionedApiClient()
		if clientVersion == nil {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeConfigInvalid, "failed to initialize API client", nil))
			return
		}

		// Type switch to handle different client versions
		switch client := clientVersion.(type) {
		case *v2_2.Client:
			if client.ClientOps == nil {
				shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, "ClientOps is nil for v2.2 client", nil))
				return
			}

			users, err := client.ClientOps.GetUsers(maxResultCount, skipCount, filter)
			if err != nil {
				shared.HandleError(cmd, err)
				return
			}

			// Get the output format option
			outputFormat, _ := cmd.Flags().GetString("output")

			// Validate output format
			validFormats := []string{"json", "short", "table"}
			isValidFormat := false
			for _, format := range validFormats {
				if outputFormat == format {
					isValidFormat = true
					break
				}
			}

			if !isValidFormat {
				shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed,
					fmt.Sprintf("Invalid output format '%s'. Valid options are: %s", outputFormat, strings.Join(validFormats, ", ")), nil))
				return
			}

			// Use the output format-specific function
			switch outputFormat {
			case "json":
				if err := shared.PrintJSONResponse(users); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := printSimpleUsersList(users); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := printSimpleUsersTable(users); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(users, maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var createUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user with the specified parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		bodyJSON, _ := cmd.Flags().GetString("body")

		if bodyJSON == "" {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "body is required", nil))
			return
		}

		// Parse the JSON body
		var body interface{}
		if err := json.Unmarshal([]byte(bodyJSON), &body); err != nil {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "invalid JSON body", err))
			return
		}

		clientVersion := shared.GetVersionedApiClient()
		if clientVersion == nil {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeConfigInvalid, "failed to initialize API client", nil))
			return
		}

		// Type switch to handle different client versions
		switch client := clientVersion.(type) {
		case *v2_2.Client:
			if client.ClientOps == nil {
				shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, "ClientOps is nil for v2.2 client", nil))
				return
			}

			result, err := client.ClientOps.CreateUser(body)
			if err != nil {
				shared.HandleError(cmd, err)
				return
			}

			shared.LogInfo("User created successfully", "result", result)

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

// Helper functions for users output formatting are now in utils.go

func init() {
	// Add flags to get users command
	getUsersCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getUsersCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getUsersCmd.Flags().String("filter", "", "Search filter string")
	getUsersCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getUsersCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to create user command
	createUserCmd.Flags().String("body", "", "JSON body for user creation (required)")

	// Commands will be added to users command group via InitClientCommands
}
