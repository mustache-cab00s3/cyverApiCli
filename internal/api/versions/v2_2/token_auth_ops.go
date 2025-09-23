package v2_2

import (
	"errors"
	"net/http"

	"github.com/yourusername/cyverApiCli/internal/api"
)

const (
	// TwoFactorMethodEmail represents the email method for 2FA
	TwoFactorMethodEmail = "Email"
	// TwoFactorMethodSMS represents the SMS method for 2FA
	TwoFactorMethodAuthenticator = "GoogleAuthenticator"
)

// TokenAuthOps implements the TokenAuthInterface for V2.2
type TokenAuthOps struct {
	*Client
}

func (c *TokenAuthOps) ApiTokenauthAuthenticatePost(body interface{}) (*AuthenticateResultModelAjaxResponse, error) {
	getLogger().Debug("Starting ApiTokenauthAuthenticatePost request")

	var response AuthenticateResultModelAjaxResponse
	path := "/api/TokenAuth/Authenticate"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiTokenauthAuthenticatePost", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiTokenauthAuthenticatePost")
	return &response, nil
}

// Get userId
func (t *TokenAuthOps) GetUserId(authRequest AuthenticateModel) (*AuthenticateResultModel, error) {
	getLogger().Debug("Starting GetUserId request", "userNameOrEmailAddress", authRequest.UserNameOrEmailAddress)

	if authRequest.UserNameOrEmailAddress == "" || authRequest.Password == "" {
		getLogger().Error("Invalid credentials provided", "userNameOrEmailAddress", authRequest.UserNameOrEmailAddress)
		return nil, NewInvalidCredentialsError(errors.New("username and password are required and are showing blank"))
	}

	data := ReducedAuthenticateModel{
		UserNameOrEmailAddress: authRequest.UserNameOrEmailAddress,
		Password:               authRequest.Password,
	}

	getLogger().Info("Making API request", "method", http.MethodPost, "path", "/api/TokenAuth/Authenticate")
	var response AuthenticateResultModelAjaxResponse
	_, err := t.Client.DoRequest(http.MethodPost, "/api/TokenAuth/Authenticate", data, &response)
	if err != nil {
		getLogger().Error("Authentication request failed", "userNameOrEmailAddress", authRequest.UserNameOrEmailAddress, "error", err)
		return nil, NewAPIError(http.StatusInternalServerError, "authentication request failed", err)
	}

	if !response.Success {
		errorMsg := "authentication failed"
		if response.Error != nil && response.Error.Message != nil {
			errorMsg = *response.Error.Message
		}
		getLogger().Error("Authentication failed", "userNameOrEmailAddress", authRequest.UserNameOrEmailAddress, "error", errorMsg)
		return nil, NewInvalidCredentialsError(errors.New("authentication failed: " + errorMsg))
	}

	getLogger().Debug("Successfully authenticated user", "userNameOrEmailAddress", authRequest.UserNameOrEmailAddress)
	return response.Result, nil
}

func (c *TokenAuthOps) ApiTokenauthRefreshtokenPost(refreshToken string) (*RefreshTokenResultAjaxResponse, error) {
	getLogger().Debug("Starting ApiTokenauthRefreshtokenPost request", "refreshTokenLength", len(refreshToken))

	var response RefreshTokenResultAjaxResponse
	path := "/api/TokenAuth/RefreshToken"

	// Add refreshToken as query parameter
	pathWithQuery := path + "?refreshToken=" + refreshToken

	getLogger().Info("Making API request", "method", http.MethodPost, "path", pathWithQuery)
	_, err := c.DoRequest(http.MethodPost, pathWithQuery, nil, &response)
	if err != nil {
		getLogger().Error("Failed ApiTokenauthRefreshtokenPost", "refreshTokenLength", len(refreshToken), "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiTokenauthRefreshtokenPost", "refreshTokenLength", len(refreshToken))
	return &response, nil
}

func (c *TokenAuthOps) ApiTokenauthSendtwofactorauthcodePost(body interface{}) (*interface{}, error) {
	getLogger().Debug("Starting ApiTokenauthSendtwofactorauthcodePost request")

	var response interface{}
	path := "/api/TokenAuth/SendTwoFactorAuthCode"

	getLogger().Info("Making API request", "method", http.MethodPost, "path", path)
	_, err := c.DoRequest(http.MethodPost, path, body, &response)
	if err != nil {
		getLogger().Error("Failed ApiTokenauthSendtwofactorauthcodePost", "error", err)
		return nil, err
	}

	getLogger().Debug("Successfully completed ApiTokenauthSendtwofactorauthcodePost")
	return &response, nil
}

// GenerateToken is a legacy method that is maintained for backward compatibility.
// It is recommended to use Authenticate instead.
//
// Deprecated: Use Authenticate instead.
func (t *TokenAuthOps) GenerateToken() (*api.TokenInfo, error) {
	return nil, errors.New("use Authenticate instead")
}

// ValidateToken validates an existing token.
// Currently, it only performs basic validation. In a production environment,
// this should make an API call to validate the token.
//
// Parameters:
//   - token: The token to validate
//
// Returns:
//   - error: Returns an error if validation fails
func (t *TokenAuthOps) ValidateToken(token string) error {
	getLogger().Debug("Starting ValidateToken request", "tokenLength", len(token))

	if token == "" {
		getLogger().Error("Token validation failed", "reason", "token is empty")
		return NewInvalidTokenError(errors.New("token is required"))
	}

	getLogger().Debug("Successfully validated token", "tokenLength", len(token))
	return nil
}

// RevokeToken revokes an existing token.
// Currently, it only performs basic validation. In a production environment,
// this should make an API call to revoke the token.
//
// Parameters:
//   - token: The token to revoke
//
// Returns:
//   - error: Returns an error if revocation fails
func (t *TokenAuthOps) RevokeToken(token string) error {
	getLogger().Debug("Starting RevokeToken request", "tokenLength", len(token))

	if token == "" {
		getLogger().Error("Token revocation failed", "reason", "token is empty")
		return NewInvalidTokenError(errors.New("token is required"))
	}

	getLogger().Debug("Successfully revoked token", "tokenLength", len(token))
	return nil
}
