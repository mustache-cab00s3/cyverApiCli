# Error Handling in Cyver API CLI

This document describes the comprehensive error handling system implemented in the Cyver API CLI.

## Overview

The error handling system provides:
- **Structured error types** with consistent interfaces
- **Error categorization** by severity and type
- **Retry mechanisms** for transient failures
- **User-friendly error messages** with appropriate exit codes
- **Comprehensive logging** for debugging

## Error Types

### CyverError

The base error type that all application errors extend:

```go
type CyverError struct {
    Code       ErrorCode     `json:"code"`
    Message    string        `json:"message"`
    Details    string        `json:"details,omitempty"`
    Severity   ErrorSeverity `json:"severity"`
    Timestamp  time.Time     `json:"timestamp"`
    Context    map[string]interface{} `json:"context,omitempty"`
    Retryable  bool          `json:"retryable"`
    StatusCode int           `json:"status_code,omitempty"`
    Err        error         `json:"-"`
}
```

### Error Codes

| Code | Description | Severity | Retryable |
|------|-------------|----------|-----------|
| `CONFIG_MISSING` | Configuration missing | High | No |
| `CONFIG_INVALID` | Invalid configuration | High | No |
| `AUTH_FAILED` | Authentication failed | High | No |
| `TOKEN_EXPIRED` | Token expired | High | Yes |
| `API_UNAUTHORIZED` | API unauthorized | Medium | No |
| `API_FORBIDDEN` | API forbidden | Medium | No |
| `API_RATE_LIMITED` | Rate limited | Medium | Yes |
| `API_SERVER_ERROR` | Server error | Medium | Yes |
| `API_NETWORK_ERROR` | Network error | Medium | Yes |
| `API_TIMEOUT` | Request timeout | Medium | Yes |
| `VALIDATION_FAILED` | Validation failed | Medium | No |
| `INTERNAL_ERROR` | Internal error | Critical | No |

### Error Severity Levels

- **Low**: Warnings that don't prevent operation
- **Medium**: Errors that prevent current operation
- **High**: Errors that prevent application functionality
- **Critical**: Errors that indicate system failure

## Usage Examples

### Creating Errors

```go
// Simple error
err := errors.NewCyverError(errors.ErrCodeValidationFailed, "invalid input", nil)

// Error with details
err := errors.NewCyverErrorWithDetails(
    errors.ErrCodeAPIUnauthorized,
    "authentication failed",
    "token expired at 2024-01-01T00:00:00Z",
    nil,
)

// API error with status code
err := errors.NewAPIError(401, "unauthorized", nil)

// Wrapping existing error
err := errors.WrapError(originalErr, errors.ErrCodeInternalError, "failed to process request")
```

### Error Handling in Commands

```go
func runCommand(cmd *cobra.Command, args []string) {
    // Validate input
    if err := validateInput(args); err != nil {
        cmd.HandleError(cmd, err)
        return
    }

    // Make API call
    result, err := apiClient.DoRequest("GET", "/projects", nil, &projects)
    if err != nil {
        cmd.HandleError(cmd, err)
        return
    }

    // Process result
    if err := processResult(result); err != nil {
        cmd.HandleError(cmd, err)
        return
    }
}
```

### Retry Logic

```go
// Simple retry
result, err := errors.Retry(ctx, func() error {
    return apiClient.DoRequest("GET", "/projects", nil, &projects)
}, errors.DefaultRetryConfig())

// Retry with custom configuration
config := errors.NewRetryConfigBuilder().
    WithMaxAttempts(5).
    WithBaseDelay(time.Second).
    WithMaxDelay(30 * time.Second).
    Build()

result, err := errors.RetryWithResult(ctx, func() (interface{}, error) {
    return apiClient.DoRequest("GET", "/projects", nil, &projects)
}, config)
```

### Validation

```go
// Validate struct with tags
type ProjectRequest struct {
    Name        string `validate:"required,min=1,max=100"`
    Description string `validate:"max=500"`
    Status      string `validate:"required,oneof=active inactive"`
}

// Validate manually
validator := errors.NewValidator().
    AddRule(errors.RequiredString("Name")).
    AddRule(errors.StringWithLength("Description", 0, 500))

errors := validator.Validate(projectRequest)
if errors.HasErrors() {
    // Handle validation errors
}
```

## Best Practices

### 1. Error Creation
- Always use appropriate error codes
- Provide clear, user-friendly messages
- Include context when helpful
- Set appropriate severity levels

### 2. Error Handling
- Handle errors at the appropriate level
- Don't ignore errors
- Provide meaningful error messages to users
- Log errors with sufficient context

### 3. Retry Logic
- Only retry retryable errors
- Use exponential backoff
- Set reasonable retry limits
- Provide user feedback during retries

### 4. Validation
- Validate input early
- Use structured validation rules
- Provide clear validation error messages
- Validate both required and optional fields

### 5. Logging
- Log errors with appropriate levels
- Include context in log messages
- Use structured logging
- Don't log sensitive information

## Error Recovery

### Automatic Recovery
- Retry transient errors
- Fallback to alternative endpoints
- Use cached data when available

### User Recovery
- Provide clear error messages
- Suggest corrective actions
- Offer alternative commands
- Display help information

## Testing Error Handling

### Unit Tests
```go
func TestErrorHandling(t *testing.T) {
    // Test error creation
    err := errors.NewCyverError(errors.ErrCodeValidationFailed, "test error", nil)
    assert.Equal(t, errors.ErrCodeValidationFailed, err.Code)
    assert.Equal(t, errors.SeverityMedium, err.Severity)
    assert.False(t, err.IsRetryable())

    // Test error wrapping
    originalErr := fmt.Errorf("original error")
    wrappedErr := errors.WrapError(originalErr, errors.ErrCodeInternalError, "wrapped")
    assert.Equal(t, originalErr, wrappedErr.Unwrap())
}
```

### Integration Tests
```go
func TestAPIRetry(t *testing.T) {
    // Test retry logic
    ctx := context.Background()
    config := errors.DefaultRetryConfig()
    
    result, err := errors.RetryWithResult(ctx, func() (string, error) {
        // Simulate transient error
        if attempt < 2 {
            return "", errors.NewCyverError(errors.ErrCodeAPIServerError, "server error", nil)
        }
        return "success", nil
    }, config)
    
    assert.NoError(t, err)
    assert.Equal(t, "success", result)
}
```

## Monitoring and Alerting

### Error Metrics
- Error rate by type
- Error rate by severity
- Retry success rate
- Error recovery time

### Alerting Rules
- Critical errors: Immediate alert
- High severity errors: Alert within 5 minutes
- Medium severity errors: Alert within 15 minutes
- Low severity errors: Log only

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Check API key configuration
   - Verify token expiration
   - Ensure proper permissions

2. **Network Errors**
   - Check internet connection
   - Verify API endpoint availability
   - Check firewall settings

3. **Validation Errors**
   - Review input parameters
   - Check data format requirements
   - Verify required fields

4. **Rate Limiting**
   - Implement exponential backoff
   - Reduce request frequency
   - Use batch operations when possible

### Debug Mode

Enable debug mode for detailed error information:

```bash
cyver-api-cli --verbose=3 command
```

This will show:
- Detailed error context
- Request/response dumps
- Retry attempts
- Validation details
