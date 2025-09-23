package errors

import (
	"fmt"
	"net/http"
	"time"
)

// ErrorCode represents a standardized error code
type ErrorCode string

const (
	// Configuration errors
	ErrCodeConfigMissing      ErrorCode = "CONFIG_MISSING"
	ErrCodeConfigInvalid      ErrorCode = "CONFIG_INVALID"
	ErrCodeConfigFileNotFound ErrorCode = "CONFIG_FILE_NOT_FOUND"

	// Authentication errors
	ErrCodeAuthFailed         ErrorCode = "AUTH_FAILED"
	ErrCodeTokenExpired       ErrorCode = "TOKEN_EXPIRED"
	ErrCodeTokenInvalid       ErrorCode = "TOKEN_INVALID"
	ErrCodeCredentialsInvalid ErrorCode = "CREDENTIALS_INVALID"
	ErrCodeTwoFactorRequired  ErrorCode = "TWO_FACTOR_REQUIRED"

	// API errors
	ErrCodeAPINotFound     ErrorCode = "API_NOT_FOUND"
	ErrCodeAPIBadRequest   ErrorCode = "API_BAD_REQUEST"
	ErrCodeAPIUnauthorized ErrorCode = "API_UNAUTHORIZED"
	ErrCodeAPIForbidden    ErrorCode = "API_FORBIDDEN"
	ErrCodeAPIRateLimited  ErrorCode = "API_RATE_LIMITED"
	ErrCodeAPIServerError  ErrorCode = "API_SERVER_ERROR"
	ErrCodeAPINetworkError ErrorCode = "API_NETWORK_ERROR"
	ErrCodeAPITimeout      ErrorCode = "API_TIMEOUT"

	// Validation errors
	ErrCodeValidationFailed ErrorCode = "VALIDATION_FAILED"
	ErrCodeInvalidInput     ErrorCode = "INVALID_INPUT"
	ErrCodeMissingRequired  ErrorCode = "MISSING_REQUIRED"

	// Output errors
	ErrCodeOutputFormatInvalid ErrorCode = "OUTPUT_FORMAT_INVALID"
	ErrCodeOutputRenderFailed  ErrorCode = "OUTPUT_RENDER_FAILED"

	// Internal errors
	ErrCodeInternalError  ErrorCode = "INTERNAL_ERROR"
	ErrCodeNotImplemented ErrorCode = "NOT_IMPLEMENTED"
	ErrCodeUnexpectedType ErrorCode = "UNEXPECTED_TYPE"
)

// ErrorSeverity represents the severity level of an error
type ErrorSeverity int

const (
	SeverityLow ErrorSeverity = iota
	SeverityMedium
	SeverityHigh
	SeverityCritical
)

// CyverError represents a standardized error for the Cyver API CLI
type CyverError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	Details    string                 `json:"details,omitempty"`
	Severity   ErrorSeverity          `json:"severity"`
	Timestamp  time.Time              `json:"timestamp"`
	Context    map[string]interface{} `json:"context,omitempty"`
	Retryable  bool                   `json:"retryable"`
	StatusCode int                    `json:"status_code,omitempty"`
	Err        error                  `json:"-"`
}

// Error implements the error interface
func (e *CyverError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *CyverError) Unwrap() error {
	return e.Err
}

// IsRetryable returns true if the error is retryable
func (e *CyverError) IsRetryable() bool {
	return e.Retryable
}

// GetSeverity returns the error severity
func (e *CyverError) GetSeverity() ErrorSeverity {
	return e.Severity
}

// GetStatusCode returns the HTTP status code if applicable
func (e *CyverError) GetStatusCode() int {
	return e.StatusCode
}

// AddContext adds additional context to the error
func (e *CyverError) AddContext(key string, value interface{}) {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
}

// NewCyverError creates a new CyverError
func NewCyverError(code ErrorCode, message string, err error) *CyverError {
	return &CyverError{
		Code:      code,
		Message:   message,
		Severity:  getSeverityForCode(code),
		Timestamp: time.Now(),
		Retryable: isRetryableCode(code),
		Err:       err,
	}
}

// NewCyverErrorWithDetails creates a new CyverError with additional details
func NewCyverErrorWithDetails(code ErrorCode, message, details string, err error) *CyverError {
	return &CyverError{
		Code:      code,
		Message:   message,
		Details:   details,
		Severity:  getSeverityForCode(code),
		Timestamp: time.Now(),
		Retryable: isRetryableCode(code),
		Err:       err,
	}
}

// NewAPIError creates a new API-related error
func NewAPIError(statusCode int, message string, err error) *CyverError {
	code := getErrorCodeForStatusCode(statusCode)
	return &CyverError{
		Code:       code,
		Message:    message,
		Severity:   getSeverityForCode(code),
		Timestamp:  time.Now(),
		Retryable:  isRetryableCode(code),
		StatusCode: statusCode,
		Err:        err,
	}
}

// WrapError wraps an existing error with CyverError context
func WrapError(err error, code ErrorCode, message string) *CyverError {
	if cyverErr, ok := err.(*CyverError); ok {
		// If it's already a CyverError, add context but don't wrap
		cyverErr.AddContext("wrapped_message", message)
		return cyverErr
	}

	return &CyverError{
		Code:      code,
		Message:   message,
		Severity:  getSeverityForCode(code),
		Timestamp: time.Now(),
		Retryable: isRetryableCode(code),
		Err:       err,
	}
}

// getSeverityForCode returns the appropriate severity for an error code
func getSeverityForCode(code ErrorCode) ErrorSeverity {
	switch code {
	case ErrCodeConfigMissing, ErrCodeConfigInvalid, ErrCodeConfigFileNotFound:
		return SeverityHigh
	case ErrCodeAuthFailed, ErrCodeTokenExpired, ErrCodeTokenInvalid, ErrCodeCredentialsInvalid:
		return SeverityHigh
	case ErrCodeAPINotFound, ErrCodeAPIBadRequest, ErrCodeAPIUnauthorized, ErrCodeAPIForbidden:
		return SeverityMedium
	case ErrCodeAPIRateLimited, ErrCodeAPIServerError, ErrCodeAPINetworkError, ErrCodeAPITimeout:
		return SeverityMedium
	case ErrCodeValidationFailed, ErrCodeInvalidInput, ErrCodeMissingRequired:
		return SeverityMedium
	case ErrCodeOutputFormatInvalid, ErrCodeOutputRenderFailed:
		return SeverityLow
	case ErrCodeInternalError, ErrCodeNotImplemented, ErrCodeUnexpectedType:
		return SeverityCritical
	default:
		return SeverityMedium
	}
}

// isRetryableCode returns true if the error code represents a retryable error
func isRetryableCode(code ErrorCode) bool {
	switch code {
	case ErrCodeAPIRateLimited, ErrCodeAPIServerError, ErrCodeAPINetworkError, ErrCodeAPITimeout:
		return true
	case ErrCodeTokenExpired:
		return true // Token can be refreshed
	default:
		return false
	}
}

// getErrorCodeForStatusCode maps HTTP status codes to error codes
func getErrorCodeForStatusCode(statusCode int) ErrorCode {
	switch statusCode {
	case http.StatusBadRequest:
		return ErrCodeAPIBadRequest
	case http.StatusUnauthorized:
		return ErrCodeAPIUnauthorized
	case http.StatusForbidden:
		return ErrCodeAPIForbidden
	case http.StatusNotFound:
		return ErrCodeAPINotFound
	case http.StatusTooManyRequests:
		return ErrCodeAPIRateLimited
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		return ErrCodeAPIServerError
	default:
		if statusCode >= 400 && statusCode < 500 {
			return ErrCodeAPIBadRequest
		}
		if statusCode >= 500 {
			return ErrCodeAPIServerError
		}
		return ErrCodeInternalError
	}
}

// IsCyverError checks if an error is a CyverError
func IsCyverError(err error) bool {
	_, ok := err.(*CyverError)
	return ok
}

// GetErrorCode extracts the error code from an error
func GetErrorCode(err error) ErrorCode {
	if cyverErr, ok := err.(*CyverError); ok {
		return cyverErr.Code
	}
	return ErrCodeInternalError
}

// GetErrorSeverity extracts the severity from an error
func GetErrorSeverity(err error) ErrorSeverity {
	if cyverErr, ok := err.(*CyverError); ok {
		return cyverErr.Severity
	}
	return SeverityMedium
}

// IsRetryable checks if an error is retryable
func IsRetryable(err error) bool {
	if cyverErr, ok := err.(*CyverError); ok {
		return cyverErr.IsRetryable()
	}
	return false
}

// ErrorCollection represents a collection of errors
type ErrorCollection struct {
	Errors []*CyverError `json:"errors"`
}

// Add adds an error to the collection
func (ec *ErrorCollection) Add(err *CyverError) {
	ec.Errors = append(ec.Errors, err)
}

// HasErrors returns true if the collection has any errors
func (ec *ErrorCollection) HasErrors() bool {
	return len(ec.Errors) > 0
}

// GetHighestSeverity returns the highest severity level in the collection
func (ec *ErrorCollection) GetHighestSeverity() ErrorSeverity {
	highest := SeverityLow
	for _, err := range ec.Errors {
		if err.Severity > highest {
			highest = err.Severity
		}
	}
	return highest
}

// Error returns a string representation of all errors
func (ec *ErrorCollection) Error() string {
	if len(ec.Errors) == 0 {
		return "no errors"
	}

	if len(ec.Errors) == 1 {
		return ec.Errors[0].Error()
	}

	msg := fmt.Sprintf("%d errors: ", len(ec.Errors))
	for i, err := range ec.Errors {
		if i > 0 {
			msg += "; "
		}
		msg += err.Error()
	}
	return msg
}
