package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/cyverApiCli/internal/errors"
	log "github.com/yourusername/cyverApiCli/logger"
)

// HandleError handles errors in command execution
func HandleError(cmd *cobra.Command, err error) {
	if err == nil {
		return
	}

	// Get logger instance
	logger := log.GetLogger(verboseLevel)

	// Check if it's a CyverError
	if cyverErr, ok := err.(*errors.CyverError); ok {
		handleCyverError(cmd, cyverErr, logger)
	} else {
		handleGenericError(cmd, err, logger)
	}
}

// handleCyverError handles CyverError instances
func handleCyverError(_ *cobra.Command, cyverErr *errors.CyverError, logger *log.Logger) {
	// Log the error with appropriate level
	switch cyverErr.Severity {
	case errors.SeverityLow:
		logger.Warn("Command warning", "error", cyverErr.Error())
	case errors.SeverityMedium:
		logger.Error("Command error", "error", cyverErr.Error())
	case errors.SeverityHigh:
		logger.Error("Command failed", "error", cyverErr.Error())
	case errors.SeverityCritical:
		logger.Error("Critical error", "error", cyverErr.Error())
	}

	// Print user-friendly error message
	fmt.Fprintf(os.Stderr, "Error: %s\n", cyverErr.Message)
	if cyverErr.Details != "" {
		fmt.Fprintf(os.Stderr, "Details: %s\n", cyverErr.Details)
	}

	// Exit with appropriate code
	exitCode := getExitCodeForSeverity(cyverErr.Severity)
	os.Exit(exitCode)
}

// handleGenericError handles generic errors
func handleGenericError(_ *cobra.Command, err error, logger *log.Logger) {
	logger.Error("Unexpected error", "error", err)
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

// getExitCodeForSeverity returns the appropriate exit code for error severity
func getExitCodeForSeverity(severity errors.ErrorSeverity) int {
	switch severity {
	case errors.SeverityLow:
		return 0 // Warning, don't exit
	case errors.SeverityMedium:
		return 1 // General error
	case errors.SeverityHigh:
		return 2 // High severity error
	case errors.SeverityCritical:
		return 3 // Critical error
	default:
		return 1
	}
}

// ValidateAndHandleErrors validates input and handles any validation errors
func ValidateAndHandleErrors(_ *cobra.Command, data interface{}) {
	// Validate the data
	validationErrors := errors.ValidateStruct(data)
	if validationErrors.HasErrors() {
		// Log validation errors
		logger := log.GetLogger(verboseLevel)
		for _, err := range validationErrors.Errors {
			logger.Error("Validation error", "error", err.Error())
		}

		// Print user-friendly error message
		fmt.Fprintf(os.Stderr, "Validation failed: %s\n", validationErrors.Error())
		os.Exit(1)
	}
}

// CheckAPIError checks if an error is an API error and handles it appropriately
func CheckAPIError(_ *cobra.Command, err error) {
	if err == nil {
		return
	}

	// Check if it's an API error
	if cyverErr, ok := err.(*errors.CyverError); ok {
		switch cyverErr.Code {
		case errors.ErrCodeAPIUnauthorized:
			fmt.Fprintf(os.Stderr, "Authentication failed. Please check your API key or token.\n")
		case errors.ErrCodeAPIForbidden:
			fmt.Fprintf(os.Stderr, "Access denied. You don't have permission to perform this action.\n")
		case errors.ErrCodeAPINotFound:
			fmt.Fprintf(os.Stderr, "Resource not found. Please check the resource ID or path.\n")
		case errors.ErrCodeAPIRateLimited:
			fmt.Fprintf(os.Stderr, "Rate limit exceeded. Please wait before making another request.\n")
		case errors.ErrCodeAPIServerError:
			fmt.Fprintf(os.Stderr, "Server error. Please try again later.\n")
		case errors.ErrCodeAPINetworkError:
			fmt.Fprintf(os.Stderr, "Network error. Please check your connection.\n")
		case errors.ErrCodeAPITimeout:
			fmt.Fprintf(os.Stderr, "Request timeout. Please try again.\n")
		}
	}
}

// RetryableErrorHandler handles retryable errors with user feedback
func RetryableErrorHandler(_ *cobra.Command, err error, attempt int, maxAttempts int) bool {
	if err == nil {
		return false
	}

	// Check if error is retryable
	if !errors.IsRetryable(err) {
		return false
	}

	// Check if we should retry
	if attempt >= maxAttempts {
		return false
	}

	// Log retry attempt
	logger := log.GetLogger(verboseLevel)
	logger.Warn("Retrying request", "attempt", attempt+1, "max_attempts", maxAttempts, "error", err)

	// Print user feedback
	fmt.Fprintf(os.Stderr, "Retrying request (attempt %d/%d)...\n", attempt+1, maxAttempts)

	return true
}

// PrintErrorSummary prints a summary of errors for batch operations
func PrintErrorSummary(errorList []error) {
	if len(errorList) == 0 {
		return
	}

	fmt.Fprintf(os.Stderr, "\nError Summary:\n")
	fmt.Fprintf(os.Stderr, "==============\n")

	errorCounts := make(map[string]int)
	for _, err := range errorList {
		if cyverErr, ok := err.(*errors.CyverError); ok {
			errorCounts[string(cyverErr.Code)]++
		} else {
			errorCounts["UNKNOWN"]++
		}
	}

	for errorType, count := range errorCounts {
		fmt.Fprintf(os.Stderr, "%s: %d occurrences\n", errorType, count)
	}

	fmt.Fprintf(os.Stderr, "Total errors: %d\n", len(errorList))
}
