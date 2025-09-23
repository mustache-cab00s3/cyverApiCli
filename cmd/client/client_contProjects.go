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

// Continuous Project commands
var getContinuousProjectsCmd = &cobra.Command{
	Use:   "continuous-projects",
	Short: "Get continuous projects",
	Long:  `Retrieve a list of continuous projects with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
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

			continuousProjects, err := client.ClientOps.GetContinuousProjects(status, maxResultCount, skipCount, filter)
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
				if err := shared.PrintJSONResponse(continuousProjects); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := printSimpleContinuousProjectsList(continuousProjects); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := printSimpleContinuousProjectsTable(continuousProjects); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(continuousProjects, maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var getContinuousProjectByIDCmd = &cobra.Command{
	Use:   "continuous-project [project-id]",
	Short: "Get continuous project by ID",
	Long:  `Retrieve detailed information about a specific continuous project by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

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

			continuousProject, err := client.ClientOps.GetContinuousProjectByID(projectID)
			if err != nil {
				shared.LogError("Error: failed to get continuous project", "error", err)
				return
			}

			// Get the output format option
			outputFormat, _ := cmd.Flags().GetString("output")

			// Validate output format
			validFormats := []string{"json", "table", "custom"}
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
				if err := shared.PrintJSONResponse(continuousProject); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := printContinuousProjectTable(continuousProject); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(continuousProject, maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var getContinuousProjectRequestFormsCmd = &cobra.Command{
	Use:   "continuous-request-forms",
	Short: "Get continuous project request forms",
	Long:  `Retrieve continuous project request forms with optional filtering.`,
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

			requestForms, err := client.ClientOps.GetContinuousProjectRequestForms(maxResultCount, skipCount, filter)
			if err != nil {
				shared.LogError("Error: failed to get continuous project request forms", "error", err)
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
				if err := shared.PrintJSONResponse(requestForms); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := printSimpleContinuousRequestFormsList(requestForms); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := printSimpleContinuousRequestFormsTable(requestForms); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(requestForms, maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var requestContinuousProjectCmd = &cobra.Command{
	Use:   "continuous-request",
	Short: "Request a new continuous project",
	Long:  `Request a new continuous project with the specified parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		triggerEvents, _ := cmd.Flags().GetBool("trigger-events")
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

			result, err := client.ClientOps.RequestContinuousProject(triggerEvents, body)
			if err != nil {
				shared.LogError("Error: failed to request continuous project", "error", err)
				return
			}

			shared.LogInfo("Continuous project requested successfully", "result", result)

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

// Helper functions for continuous project output formatting are now in utils.go

// InitContinuousProjectsCommands initializes continuous projects commands and adds them to the client command group
func InitContinuousProjectsCommands(clientCmd *cobra.Command) {
	// Add flags to get continuous projects command
	getContinuousProjectsCmd.Flags().String("status", "", "Filter by project status")
	getContinuousProjectsCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getContinuousProjectsCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getContinuousProjectsCmd.Flags().String("filter", "", "Search filter string")
	getContinuousProjectsCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getContinuousProjectsCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to get continuous project by ID command
	getContinuousProjectByIDCmd.Flags().String("output", "table", "Output format: json (complete JSON), table (formatted table), or custom (interactive field selection)")
	getContinuousProjectByIDCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to get continuous project request forms command
	getContinuousProjectRequestFormsCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getContinuousProjectRequestFormsCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getContinuousProjectRequestFormsCmd.Flags().String("filter", "", "Search filter string")
	getContinuousProjectRequestFormsCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getContinuousProjectRequestFormsCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to request continuous project command
	requestContinuousProjectCmd.Flags().Bool("trigger-events", false, "Trigger events after request")
	requestContinuousProjectCmd.Flags().String("body", "", "JSON body for continuous project request (required)")

	// Add continuous project commands to client command group
	clientCmd.AddCommand(getContinuousProjectsCmd)
	clientCmd.AddCommand(getContinuousProjectByIDCmd)
	clientCmd.AddCommand(getContinuousProjectRequestFormsCmd)
	clientCmd.AddCommand(requestContinuousProjectCmd)
}
