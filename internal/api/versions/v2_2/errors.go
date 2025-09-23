package v2_2

import "fmt"

// AuthError represents a base error type for authentication-related errors
type AuthError struct {
	Message string
	Err     error
}

func (e *AuthError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

// InvalidCredentialsError represents an error when authentication credentials are invalid
type InvalidCredentialsError struct {
	*AuthError
}

// NewInvalidCredentialsError creates a new InvalidCredentialsError
func NewInvalidCredentialsError(err error) *InvalidCredentialsError {
	getLogger().Error("invalid credentials", "error", err)
	return &InvalidCredentialsError{
		AuthError: &AuthError{
			Message: "invalid credentials",
			Err:     err,
		},
	}
}

// TokenExpiredError represents an error when a token has expired
type TokenExpiredError struct {
	*AuthError
}

// NewTokenExpiredError creates a new TokenExpiredError
func NewTokenExpiredError(err error) *TokenExpiredError {
	getLogger().Error("token has expired", "error", err)
	return &TokenExpiredError{
		AuthError: &AuthError{
			Message: "token has expired",
			Err:     err,
		},
	}
}

// InvalidTokenError represents an error when a token is invalid
type InvalidTokenError struct {
	*AuthError
}

// NewInvalidTokenError creates a new InvalidTokenError
func NewInvalidTokenError(err error) *InvalidTokenError {
	getLogger().Error("invalid token", "error", err)
	return &InvalidTokenError{
		AuthError: &AuthError{
			Message: "invalid token",
			Err:     err,
		},
	}
}

// TwoFactorAuthError represents an error related to two-factor authentication
type TwoFactorAuthError struct {
	*AuthError
}

// NewTwoFactorAuthError creates a new TwoFactorAuthError
func NewTwoFactorAuthError(err error) *TwoFactorAuthError {
	getLogger().Error("invalid 2fa request", "error", err)
	return &TwoFactorAuthError{
		AuthError: &AuthError{
			Message: "invalid 2fa request",
			Err:     err,
		},
	}
}

// APIError represents an error returned by the API
type APIError struct {
	*AuthError
	StatusCode int
}

// NewAPIError creates a new APIError
func NewAPIError(statusCode int, message string, err error) *APIError {
	getLogger().Error("API error", "statusCode", statusCode, "message", message, "error", err)
	return &APIError{
		AuthError: &AuthError{
			Message: message,
			Err:     err,
		},
		StatusCode: statusCode,
	}
}
