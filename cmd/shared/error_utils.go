package shared

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/cyverApiCli/internal/errors"
	log "github.com/yourusername/cyverApiCli/logger"
)

// LogErrorWithContext logs an error with appropriate level and context
func LogErrorWithContext(message string, keysAndValues ...interface{}) {
	logger := log.GetLogger(VerboseLevel)
	logger.Error(message, keysAndValues...)
}

// LogWarningWithContext logs a warning with appropriate level and context
func LogWarningWithContext(message string, keysAndValues ...interface{}) {
	logger := log.GetLogger(VerboseLevel)
	logger.Warn(message, keysAndValues...)
}

// LogInfoWithContext logs an info message with appropriate level and context
func LogInfoWithContext(message string, keysAndValues ...interface{}) {
	logger := log.GetLogger(VerboseLevel)
	logger.Info(message, keysAndValues...)
}

// LogDebugWithContext logs a debug message with appropriate level and context
func LogDebugWithContext(message string, keysAndValues ...interface{}) {
	logger := log.GetLogger(VerboseLevel)
	logger.Debug(message, keysAndValues...)
}

// HandleAPIError handles API-specific errors with user-friendly messages
func HandleAPIError(err error) {
	if err == nil {
		return
	}

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
		default:
			fmt.Fprintf(os.Stderr, "Error: %s\n", cyverErr.Message)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

// ValidateInput validates common input parameters
func ValidateInput(params map[string]interface{}) error {
	for key, value := range params {
		switch key {
		case "max-results", "skip-count":
			if intVal, ok := value.(int); ok && intVal < 0 {
				return errors.NewCyverError(errors.ErrCodeValidationFailed, fmt.Sprintf("%s must be non-negative", key), nil)
			}
		case "status":
			if strVal, ok := value.(string); ok && strVal != "" {
				validStatuses := []string{"active", "inactive", "completed", "pending"}
				isValid := false
				for _, status := range validStatuses {
					if strVal == status {
						isValid = true
						break
					}
				}
				if !isValid {
					return errors.NewCyverError(errors.ErrCodeValidationFailed, fmt.Sprintf("invalid status: %s. Valid values are: %v", strVal, validStatuses), nil)
				}
			}
		}
	}
	return nil
}

// CheckRetryableError checks if an error is retryable and provides user feedback
func CheckRetryableError(err error, attempt int, maxAttempts int) bool {
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
	LogWarningWithContext("Retrying request", "attempt", attempt+1, "max_attempts", maxAttempts, "error", err)

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

// GetExitCodeForError returns the appropriate exit code for an error
func GetExitCodeForError(err error) int {
	if cyverErr, ok := err.(*errors.CyverError); ok {
		switch cyverErr.Severity {
		case errors.SeverityLow:
			return 0 // Warning, don't exit
		case errors.SeverityMedium:
			return 1 // General error
		case errors.SeverityHigh:
			return 2 // High severity error
		case errors.SeverityCritical:
			return 3 // Critical error
		}
	}
	return 1 // Default error code
}

// HandleError handles errors in command execution
func HandleError(cmd *cobra.Command, err error) {
	if err == nil {
		return
	}

	// Get logger instance
	logger := log.GetLogger(VerboseLevel)

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
	exitCode := GetExitCodeForError(cyverErr)
	os.Exit(exitCode)
}

// handleGenericError handles generic errors
func handleGenericError(_ *cobra.Command, err error, logger *log.Logger) {
	logger.Error("Unexpected error", "error", err)
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
