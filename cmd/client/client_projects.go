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

// Project commands
var getProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get projects",
	Long:  `Retrieve a list of projects with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
		maxResultCount, _ := cmd.Flags().GetInt("max-results")
		skipCount, _ := cmd.Flags().GetInt("skip-count")
		filter, _ := cmd.Flags().GetString("filter")

		// Validate input parameters
		if maxResultCount < 0 {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
			return
		}
		if skipCount < 0 {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "skip-count must be non-negative", nil))
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

			projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
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
				if err := shared.PrintJSONResponse(projects); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := shared.PrintSimpleProjectsList(interface{}(projects)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := shared.PrintSimpleProjectsTable(interface{}(projects)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(interface{}(projects), maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var getProjectByIDCmd = &cobra.Command{
	Use:   "get [project-id]",
	Short: "Get project by ID",
	Long:  `Retrieve detailed information about a specific project by its ID.`,
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

			project, err := client.ClientOps.GetProjectByID(projectID)
			if err != nil {
				shared.HandleError(cmd, err)
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
				if err := shared.PrintJSONResponse(project); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := shared.PrintProjectTable(interface{}(project)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(interface{}(project), maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var getProjectRequestFormsCmd = &cobra.Command{
	Use:   "request-forms",
	Short: "Get project request forms",
	Long:  `Retrieve project request forms with optional filtering.`,
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

			requestForms, err := client.ClientOps.GetProjectRequestForms(maxResultCount, skipCount, filter)
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
				if err := shared.PrintJSONResponse(requestForms); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := printSimpleRequestFormsList(requestForms); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := printSimpleRequestFormsTable(requestForms); err != nil {
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

var requestProjectCmd = &cobra.Command{
	Use:   "request",
	Short: "Request a new project",
	Long:  `Request a new project with the specified parameters.`,
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

			result, err := client.ClientOps.RequestProject(triggerEvents, body)
			if err != nil {
				shared.HandleError(cmd, err)
				return
			}

			shared.LogInfo("Project requested successfully", "result", result)

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

// Helper functions for project output formatting are now in utils.go

func init() {
	// Add flags to get projects command
	getProjectsCmd.Flags().String("status", "", "Filter by project status")
	getProjectsCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getProjectsCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getProjectsCmd.Flags().String("filter", "", "Search filter string")
	getProjectsCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getProjectsCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to get project by ID command
	getProjectByIDCmd.Flags().String("output", "table", "Output format: json (complete JSON), table (formatted table), or custom (interactive field selection)")
	getProjectByIDCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to get project request forms command
	getProjectRequestFormsCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getProjectRequestFormsCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getProjectRequestFormsCmd.Flags().String("filter", "", "Search filter string")
	getProjectRequestFormsCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getProjectRequestFormsCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to request project command
	requestProjectCmd.Flags().Bool("trigger-events", false, "Trigger events after request")
	requestProjectCmd.Flags().String("body", "", "JSON body for project request (required)")

	// Commands will be added to projects command group via InitClientCommands
}
