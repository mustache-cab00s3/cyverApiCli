# Error Handling Implementation Summary

## Overview

I have implemented a comprehensive error handling system for the Cyver API CLI project. This system provides structured error management, retry mechanisms, validation, and user-friendly error reporting.

## What Was Implemented

### 1. Core Error System (`internal/errors/`)

#### `errors.go` - Main Error Types
- **CyverError**: Base error type with structured fields
- **ErrorCode**: Standardized error codes (25+ predefined codes)
- **ErrorSeverity**: Four severity levels (Low, Medium, High, Critical)
- **ErrorCollection**: For handling multiple errors
- **Helper functions**: Error creation, wrapping, and checking utilities

#### `retry.go` - Retry Mechanism
- **RetryConfig**: Configurable retry settings
- **Retry()**: Simple retry function
- **RetryWithResult()**: Retry with return values
- **RetryConfigBuilder**: Fluent interface for configuration
- **Exponential backoff** with jitter support

#### `validation.go` - Input Validation
- **ValidationRule**: Flexible validation rules
- **Validator**: Struct validation with reflection
- **Common validation rules**: Required, length, pattern, custom
- **Struct tag support**: Automatic validation from struct tags

### 2. API Client Improvements (`internal/api/client.go`)

#### Enhanced Error Handling
- **Input validation** for client creation
- **Structured error responses** with proper error codes
- **Network error detection** (timeout, connection issues)
- **HTTP status code mapping** to error types
- **Retry support** with configurable policies

#### New Methods
- `DoRequestWithRetry()`: HTTP requests with retry logic
- `DoRequestRawWithRetry()`: Raw requests with retry logic
- Enhanced logging with request/response details

### 3. Command Error Handling (`cmd/`)

#### `error_handler.go` - Command Error Utilities
- **HandleError()**: Centralized error handling for commands
- **CheckAPIError()**: API-specific error handling
- **RetryableErrorHandler()**: Retry logic for commands
- **PrintErrorSummary()**: Batch operation error reporting

#### `shared/error_utils.go` - Shared Utilities
- **Input validation** functions
- **Error recovery** suggestions
- **User-friendly messages** for common errors
- **Exit code mapping** based on error severity

### 4. Updated Command Examples

#### Before (Original)
```go
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.LogError("Error: failed to get projects", "error", err)
    return
}
```

#### After (With Error Handling)
```go
// Validate input parameters
if maxResultCount < 0 {
    cmd.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
    return
}

projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    cmd.HandleError(cmd, err)
    return
}
```

## Key Features

### 1. Structured Error Types
- **25+ predefined error codes** covering all scenarios
- **Severity-based categorization** for appropriate handling
- **Context preservation** for debugging
- **Retryable error identification**

### 2. Retry Mechanisms
- **Exponential backoff** with configurable parameters
- **Jitter support** to prevent thundering herd
- **Context-aware cancellation**
- **Retryable error detection**

### 3. Input Validation
- **Struct-based validation** with reflection
- **Common validation rules** (required, length, pattern)
- **Custom validation functions**
- **Struct tag support** for automatic validation

### 4. User Experience
- **User-friendly error messages**
- **Appropriate exit codes** based on severity
- **Error recovery suggestions**
- **Batch operation error summaries**

### 5. Developer Experience
- **Comprehensive logging** with structured data
- **Error context preservation**
- **Debug mode support**
- **Testing utilities**

## Error Code Categories

### Configuration Errors
- `CONFIG_MISSING`, `CONFIG_INVALID`, `CONFIG_FILE_NOT_FOUND`

### Authentication Errors
- `AUTH_FAILED`, `TOKEN_EXPIRED`, `TOKEN_INVALID`, `CREDENTIALS_INVALID`

### API Errors
- `API_UNAUTHORIZED`, `API_FORBIDDEN`, `API_NOT_FOUND`, `API_RATE_LIMITED`
- `API_SERVER_ERROR`, `API_NETWORK_ERROR`, `API_TIMEOUT`

### Validation Errors
- `VALIDATION_FAILED`, `INVALID_INPUT`, `MISSING_REQUIRED`

### Internal Errors
- `INTERNAL_ERROR`, `NOT_IMPLEMENTED`, `UNEXPECTED_TYPE`

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
```

### Retry Logic
```go
// Simple retry
result, err := errors.Retry(ctx, func() error {
    return apiClient.DoRequest("GET", "/projects", nil, &projects)
}, errors.DefaultRetryConfig())

// Custom retry configuration
config := errors.NewRetryConfigBuilder().
    WithMaxAttempts(5).
    WithBaseDelay(time.Second).
    WithMaxDelay(30 * time.Second).
    Build()
```

### Validation
```go
// Struct validation
type ProjectRequest struct {
    Name        string `validate:"required,min=1,max=100"`
    Description string `validate:"max=500"`
    Status      string `validate:"required,oneof=active inactive"`
}

// Manual validation
validator := errors.NewValidator().
    AddRule(errors.RequiredString("Name")).
    AddRule(errors.StringWithLength("Description", 0, 500))

errors := validator.Validate(projectRequest)
```

## Command Implementation Examples

### Before (Old Error Handling)
```go
// Old way - manual error handling
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.LogError("Error: failed to get projects", "error", err)
    return
}

// Old way - manual validation
if maxResultCount < 0 {
    shared.LogError("Error: max-results must be non-negative")
    return
}
```

### After (New Error Handling)
```go
// New way - structured error handling
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}

// New way - structured validation
if maxResultCount < 0 {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
    return
}
```

### Complete Command Example
```go
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
```

### Error Handling in Different Scenarios

#### 1. Input Validation
```go
// Validate required parameters
if bodyJSON == "" {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "body is required", nil))
    return
}

// Validate JSON parsing
var body interface{}
if err := json.Unmarshal([]byte(bodyJSON), &body); err != nil {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "invalid JSON body", err))
    return
}
```

#### 2. API Client Initialization
```go
clientVersion := shared.GetVersionedApiClient()
if clientVersion == nil {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeConfigInvalid, "failed to initialize API client", nil))
    return
}
```

#### 3. API Operations
```go
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}
```

#### 4. Output Format Validation
```go
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
```

#### 5. Output Processing
```go
switch outputFormat {
case "json":
    if err := shared.PrintJSONResponse(projects); err != nil {
        shared.HandleError(cmd, err)
    }
case "table":
    if err := shared.PrintSimpleProjectsTable(interface{}(projects)); err != nil {
        shared.HandleError(cmd, err)
    }
}
```

### Error Types and Their Usage

#### Validation Errors
```go
// Input validation
if maxResultCount < 0 {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
    return
}

// Format validation
if !isValidFormat {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, 
        fmt.Sprintf("Invalid output format '%s'", outputFormat), nil))
    return
}
```

#### Configuration Errors
```go
// API client initialization
if clientVersion == nil {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeConfigInvalid, "failed to initialize API client", nil))
    return
}
```

#### Type Errors
```go
// Client type validation
if client.ClientOps == nil {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, "ClientOps is nil for v2.2 client", nil))
    return
}

// Unsupported client type
default:
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
    return
```

#### API Errors
```go
// API operation errors (automatically handled by shared.HandleError)
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}
```

### Best Practices

#### 1. Always Use shared.HandleError
```go
// ✅ Good
if err != nil {
    shared.HandleError(cmd, err)
    return
}

// ❌ Bad
if err != nil {
    shared.LogError("Error: failed to get projects", "error", err)
    return
}
```

#### 2. Provide Context in Error Messages
```go
// ✅ Good
shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, 
    fmt.Sprintf("Invalid output format '%s'. Valid options are: %s", outputFormat, strings.Join(validFormats, ", ")), nil))

// ❌ Bad
shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "Invalid format", nil))
```

#### 3. Use Appropriate Error Codes
```go
// ✅ Good - specific error codes
errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil)
errors.NewCyverError(errors.ErrCodeConfigInvalid, "failed to initialize API client", nil)
errors.NewCyverError(errors.ErrCodeUnexpectedType, "ClientOps is nil for v2.2 client", nil)

// ❌ Bad - generic error codes
errors.NewCyverError(errors.ErrCodeInternalError, "max-results must be non-negative", nil)
```

#### 4. Handle Errors Immediately
```go
// ✅ Good - handle errors immediately
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}

// ❌ Bad - deferring error handling
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
// ... other code ...
if err != nil {
    shared.HandleError(cmd, err)
    return
}
```

### Migration Checklist

When updating existing commands to use the new error handling system:

1. **Add imports**:
   ```go
   import (
       "fmt"
       "github.com/yourusername/cyverApiCli/internal/errors"
   )
   ```

2. **Replace LogError calls**:
   ```go
   // Old
   shared.LogError("Error: failed to get projects", "error", err)
   
   // New
   shared.HandleError(cmd, err)
   ```

3. **Replace manual validation**:
   ```go
   // Old
   if maxResultCount < 0 {
       shared.LogError("Error: max-results must be non-negative")
       return
   }
   
   // New
   if maxResultCount < 0 {
       shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
       return
   }
   ```

4. **Replace default cases**:
   ```go
   // Old
   default:
       shared.LogError("Error: unsupported client type: %T", clientVersion)
       return
   
   // New
   default:
       shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeUnexpectedType, fmt.Sprintf("unsupported client type: %T", clientVersion), nil))
       return
   ```

5. **Test the updated commands** to ensure error handling works correctly.

## Benefits

### 1. Improved Reliability
- **Automatic retry** for transient failures
- **Proper error categorization** for appropriate handling
- **Input validation** to prevent invalid requests

### 2. Better User Experience
- **Clear error messages** with actionable information
- **Appropriate exit codes** for scripting
- **Error recovery suggestions**

### 3. Enhanced Debugging
- **Structured logging** with context
- **Error traceability** through error wrapping
- **Debug mode** for detailed information

### 4. Maintainability
- **Consistent error handling** across the codebase
- **Centralized error management**
- **Easy to extend** with new error types

### 5. Testing Support
- **Error type checking** for unit tests
- **Retry logic testing** utilities
- **Validation testing** helpers

## Migration Guide

### For Existing Commands
1. **Import error handling**: Add `"github.com/yourusername/cyverApiCli/internal/errors"`
2. **Replace error handling**: Use `cmd.HandleError(cmd, err)` instead of manual error handling
3. **Add validation**: Validate input parameters before API calls
4. **Update error messages**: Use structured error creation

### For New Commands
1. **Use error handling utilities** from the start
2. **Implement proper validation** for all inputs
3. **Add retry logic** for API calls where appropriate
4. **Provide user-friendly error messages**

## Testing

### Unit Tests
- Error creation and wrapping
- Retry logic with various scenarios
- Validation rules and struct validation
- Error handling utilities

### Integration Tests
- API client error handling
- Command error handling
- Retry mechanisms
- User experience flows

## Documentation

- **Comprehensive error handling guide** (`docs/error_handling.md`)
- **Usage examples** for all error types
- **Best practices** for error handling
- **Troubleshooting guide** for common issues

## Future Enhancements

### Potential Improvements
1. **Metrics collection** for error rates and types
2. **Error reporting** to external services
3. **Automatic error recovery** for common scenarios
4. **Error analytics** for improving user experience

### Extensibility
- **Custom error types** for specific use cases
- **Plugin system** for error handling
- **Configuration-driven** error handling
- **Integration** with monitoring systems

## Conclusion

The implemented error handling system provides a robust foundation for reliable, user-friendly, and maintainable error management in the Cyver API CLI. It addresses all major error scenarios while providing flexibility for future enhancements and specific use cases.

The system is designed to be:
- **Easy to use** for developers
- **Helpful for users** with clear error messages
- **Maintainable** with consistent patterns
- **Extensible** for future requirements
- **Testable** with comprehensive utilities
