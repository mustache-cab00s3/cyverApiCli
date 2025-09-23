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

// Findings commands
var getFindingsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get findings",
	Long:  `Retrieve a list of findings with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project-id")
		maxResultCount, _ := cmd.Flags().GetInt("max-results")
		skipCount, _ := cmd.Flags().GetInt("skip-count")

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

			findings, err := client.ClientOps.GetFindings(projectID, maxResultCount, skipCount)
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
				if err := shared.PrintJSONResponse(findings); err != nil {
					shared.HandleError(cmd, err)
				}
			case "short":
				if err := shared.PrintSimpleFindingsList(interface{}(findings)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := shared.PrintSimpleFindingsTable(interface{}(findings)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(interface{}(findings), maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var getFindingByIDCmd = &cobra.Command{
	Use:   "get [finding-id]",
	Short: "Get finding by ID",
	Long:  `Retrieve detailed information about a specific finding by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		findingID := args[0]
		includeEvidence, _ := cmd.Flags().GetBool("include-evidence")

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

			finding, err := client.ClientOps.GetFindingByID(findingID, includeEvidence)
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
				if err := shared.PrintJSONResponse(finding); err != nil {
					shared.HandleError(cmd, err)
				}
			case "table":
				if err := shared.PrintFindingTable(interface{}(finding)); err != nil {
					shared.HandleError(cmd, err)
				}
			case "custom":
				maxColumns, _ := cmd.Flags().GetInt("max-columns")
				if maxColumns <= 0 {
					maxColumns = 4 // Default to 4 columns
				}
				if err := shared.PrintCustomTable(interface{}(finding), maxColumns); err != nil {
					shared.HandleError(cmd, err)
				}
			}

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

var setFindingStatusCmd = &cobra.Command{
	Use:   "set-status [finding-id]",
	Short: "Set finding status",
	Long:  `Set the status of a specific finding.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		findingID := args[0]
		triggerEvents, _ := cmd.Flags().GetInt("trigger-events")
		statusBodyJSON, _ := cmd.Flags().GetString("status-body")

		if statusBodyJSON == "" {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "status-body is required", nil))
			return
		}

		// Parse the JSON body
		var statusBody interface{}
		if err := json.Unmarshal([]byte(statusBodyJSON), &statusBody); err != nil {
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "invalid JSON status body", err))
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

			err := client.ClientOps.SetFindingStatus(findingID, triggerEvents, statusBody)
			if err != nil {
				shared.HandleError(cmd, err)
				return
			}

			shared.LogInfo("Finding status updated successfully", "findingID", findingID)

		default:
			shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
			return
		}
	},
}

func init() {
	// Add flags to get findings command
	getFindingsCmd.Flags().String("project-id", "", "Filter by project ID")
	getFindingsCmd.Flags().Int("max-results", 10, "Maximum number of results")
	getFindingsCmd.Flags().Int("skip-count", 0, "Number of results to skip")
	getFindingsCmd.Flags().String("output", "table", "Output format: json (complete JSON), short (ID and name JSON), table (ID and name table), or custom (interactive field selection)")
	getFindingsCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to get finding by ID command
	getFindingByIDCmd.Flags().Bool("include-evidence", false, "Include evidence in response")
	getFindingByIDCmd.Flags().String("output", "table", "Output format: json (complete JSON), table (formatted table), or custom (interactive field selection)")
	getFindingByIDCmd.Flags().Int("max-columns", 4, "Maximum number of columns for custom table output")

	// Add flags to set finding status command
	setFindingStatusCmd.Flags().Int("trigger-events", 0, "Trigger events flag")
	setFindingStatusCmd.Flags().String("status-body", "", "JSON body for status update (required)")

	// Commands will be added to findings command group via InitClientCommands
}
