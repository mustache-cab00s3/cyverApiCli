# Cyver API CLI

A comprehensive command-line interface tool for interacting with the Cyver API. This CLI provides easy access to manage clients, projects, findings, users, teams, and continuous projects through a well-structured command hierarchy.

## Features

- **Multi-role Support**: Separate command groups for client and pentester operations
- **Flexible Output Formats**: Support for table, JSON, and custom output formats
- **Configurable Verbosity**: Multiple verbosity levels for detailed logging
- **Token Management**: Automatic token refresh and secure storage
- **API Version Support**: Support for multiple API versions (v2.2)
- **Interactive Configuration**: Guided setup and configuration management

## Installation

### From Source
```bash
go install github.com/yourusername/cyverApiCli@latest
```

### From Repository
```bash
git clone https://github.com/yourusername/cyverApiCli.git
cd cyverApiCli
go mod download
go build
```

## Quick Start

1. **Initialize Configuration**:
```bash
cyverApiCli config init
```

2. **Authenticate**:
```bash
cyverApiCli apiAuth getToken --username your-email@example.com
```

3. **List Available Commands**:
```bash
cyverApiCli --help
```

## Command Tree

```
cyverApiCli
├── apiAuth                    # API Authentication
│   └── getToken              # Get authentication token
├── client                    # Client Operations
│   ├── get-client-info       # Get client information
│   ├── update-client-info    # Update client information
│   ├── list-clients          # List all clients
│   ├── get-projects          # Get client projects
│   ├── get-project-by-id     # Get specific project by ID
│   ├── get-project-request-forms # Get project request forms
│   ├── request-project       # Request a new project
│   ├── get-continuous-projects # Get continuous projects
│   ├── get-continuous-project-by-id # Get specific continuous project
│   ├── get-continuous-project-request-forms # Get continuous project forms
│   ├── request-continuous-project # Request continuous project
│   ├── get-findings          # Get findings
│   ├── get-finding-by-id     # Get specific finding by ID
│   ├── set-finding-status    # Update finding status
│   ├── get-assets            # Get assets
│   ├── create-asset          # Create new asset
│   ├── delete-asset          # Delete asset
│   ├── update-asset          # Update asset
│   ├── get-users             # Get users
│   └── create-user           # Create new user
├── config                    # Configuration Management
│   ├── init                  # Initialize CLI configuration
│   ├── view                  # View current configuration
│   ├── refresh-token         # Refresh access token
│   └── re-auth               # Re-authenticate
├── pentester                 # Pentester Operations
│   ├── clients               # Client Management (Pentester View)
│   │   ├── list              # List pentester clients
│   │   ├── get               # Get specific client
│   │   ├── create            # Create new client
│   │   ├── update            # Update client
│   │   ├── delete            # Delete client
│   │   ├── get-assets        # Get client assets
│   │   ├── create-asset      # Create client asset
│   │   └── update-asset      # Update client asset
│   ├── findings              # Findings Management
│   │   ├── list              # List findings
│   │   ├── get               # Get specific finding
│   │   ├── create            # Create new finding
│   │   ├── update            # Update finding
│   │   ├── delete            # Delete finding
│   │   └── import            # Import findings
│   ├── projects              # Projects Management
│   │   ├── list              # List projects
│   │   ├── create            # Create new project
│   │   ├── get               # Get specific project
│   │   ├── delete            # Delete project
│   │   ├── update-status     # Update project status
│   │   ├── set-assets        # Set project assets
│   │   ├── set-users         # Set project users
│   │   ├── set-teams         # Set project teams
│   │   ├── get-checklists    # Get project checklists
│   │   ├── get-compliance-norms # Get compliance norms
│   │   ├── get-report-versions # Get report versions
│   │   ├── get-report        # Get project report
│   │   └── upload-file       # Upload project file
│   ├── users                 # Users Management
│   │   ├── list              # List users
│   │   ├── create            # Create new user
│   │   ├── get               # Get specific user
│   │   ├── update            # Update user
│   │   └── delete            # Delete user
│   ├── teams                 # Teams Management
│   │   ├── list              # List teams
│   │   ├── create            # Create new team
│   │   ├── get               # Get specific team
│   │   ├── update            # Update team
│   │   └── delete            # Delete team
│   └── continuous-projects   # Continuous Projects Management
│       ├── list              # List continuous projects
│       ├── create            # Create continuous project
│       ├── get               # Get specific continuous project
│       ├── delete            # Delete continuous project
│       ├── update-status     # Update continuous project status
│       ├── set-assets        # Set continuous project assets
│       ├── set-users         # Set continuous project users
│       ├── set-teams         # Set continuous project teams
│       ├── get-runs          # Get continuous project runs
│       ├── complete-run      # Complete continuous project run
│       ├── get-report-versions # Get report versions
│       ├── get-report        # Get continuous project report
│       └── upload-file       # Upload file to continuous project
└── help                      # Help about any command
```

## Configuration

### Initial Setup
The CLI uses a YAML configuration file located at `~/.cyverApiCli.yaml`. Initialize it with:

```bash
cyverApiCli config init
```

### Configuration Options
- **API Base URL**: The base URL for the Cyver API
- **API Version**: Supported versions (v2.2)
- **Token Management**: Automatic token refresh and storage
- **Output Format**: Default output format (table, JSON, custom)

### Viewing Configuration
```bash
cyverApiCli config view
```

## Authentication

### Getting a Token
```bash
cyverApiCli apiAuth getToken --username your-email@example.com
```

### Token Refresh
```bash
cyverApiCli config refresh-token
```

### Re-authentication
```bash
cyverApiCli config re-auth
```

## Usage Examples

### Client Operations
```bash
# List all clients
cyverApiCli client list-clients

# Get specific project
cyverApiCli client get-project-by-id --project-id 12345

# Create a new asset
cyverApiCli client create-asset --body '{"name": "Web Server", "type": "server"}'

# Get findings with custom output
cyverApiCli client get-findings --output custom --max-columns 6
```

### Pentester Operations
```bash
# List all projects (pentester view)
cyverApiCli pentester projects list

# Create a new finding
cyverApiCli pentester findings create --body '{"title": "SQL Injection", "severity": "high"}'

# Update project status
cyverApiCli pentester projects update-status --project-id 12345 --status "in-progress"

# Get team information
cyverApiCli pentester teams get --team-id 67890
```

## Output Formats

The CLI supports multiple output formats:

- **table**: Human-readable table format (default)
- **json**: Complete JSON response
- **short**: Simplified JSON with ID and name only
- **custom**: Interactive field selection for table output

### Output Format Examples
```bash
# Table output (default)
cyverApiCli client get-projects

# Full JSON output
cyverApiCli client get-projects --output json

# Custom table with specific columns
cyverApiCli client get-projects --output custom --max-columns 4
```

## Verbosity Levels

Control the amount of logging output:

- **-v**: Show basic request information (method, URL)
- **-vv**: Show request headers and basic response info
- **-vvv**: Show full request/response details including body

### Verbosity Examples
```bash
# Basic verbosity
cyverApiCli -v client get-projects

# High verbosity for debugging
cyverApiCli -vvv pentester findings list
```

## Global Flags

- `-c, --config`: Specify config file path
- `-v, --verbose`: Increase verbosity level (can be used multiple times)
- `-h, --help`: Show help information

## Development

### Prerequisites
- Go 1.23.0 or later
- Git

### Building from Source
```bash
# Clone the repository
git clone https://github.com/yourusername/cyverApiCli.git
cd cyverApiCli

# Install dependencies
go mod download

# Build the project
go build

# Run tests
go test ./...
```

### Project Structure
```
cyverApiCli/
├── cmd/                    # Command definitions
│   ├── pentester/         # Pentester-specific commands
│   └── shared/            # Shared utilities
├── internal/              # Internal packages
│   ├── api/              # API client implementations
│   └── config/           # Configuration management
├── logger/               # Logging utilities
├── output/               # Output formatting
└── main.go              # Application entry point
```

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Ensure your credentials are correct
   - Check if 2FA is enabled and provide the code
   - Verify your account has the necessary permissions

2. **Configuration Issues**
   - Run `cyverApiCli config init` to recreate configuration
   - Check file permissions on `~/.cyverApiCli.yaml`
   - Verify API base URL is correct

3. **Token Expiration**
   - Use `cyverApiCli config refresh-token` to refresh
   - Re-authenticate with `cyverApiCli config re-auth`

4. **Output Format Issues**
   - Use `--output json` for complete JSON responses
   - Use `--output custom` for interactive field selection
   - Adjust `--max-columns` for table formatting

### Getting Help
```bash
# General help
cyverApiCli --help

# Command-specific help
cyverApiCli client --help
cyverApiCli pentester projects --help

# Verbose output for debugging
cyverApiCli -vvv [command]
```

## Error Handling

The CLI includes comprehensive error handling with:
- **Structured Error Types**: Standardized error codes and severity levels
- **Automatic Retry**: Built-in retry logic for transient failures
- **Input Validation**: Comprehensive validation with clear error messages
- **User-Friendly Messages**: Clear, actionable error messages for users
- **Logging Integration**: Detailed logging for debugging and monitoring

### Error Handling Examples

#### Basic Error Handling
```go
// Validate input parameters
if maxResultCount < 0 {
    shared.HandleError(cmd, errors.NewCyverError(errors.ErrCodeValidationFailed, "max-results must be non-negative", nil))
    return
}

// Handle API errors
projects, err := client.ClientOps.GetProjects(status, maxResultCount, skipCount, filter)
if err != nil {
    shared.HandleError(cmd, err)
    return
}
```

#### Input Validation
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

#### Output Format Validation
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

#### Error Types and Codes

The CLI uses structured error codes for better error handling:

**Validation Errors**:
- `VALIDATION_FAILED`: Input validation failed
- `INVALID_INPUT`: Invalid input format
- `MISSING_REQUIRED`: Required parameter missing

**API Errors**:
- `API_UNAUTHORIZED`: Authentication failed
- `API_FORBIDDEN`: Access denied
- `API_NOT_FOUND`: Resource not found
- `API_RATE_LIMITED`: Rate limit exceeded

**Configuration Errors**:
- `CONFIG_INVALID`: Invalid configuration
- `CONFIG_MISSING`: Configuration missing

**Internal Errors**:
- `INTERNAL_ERROR`: Unexpected internal error
- `NOT_IMPLEMENTED`: Feature not implemented
- `UNEXPECTED_TYPE`: Unexpected type error

#### Error Severity Levels

- **Low**: Warnings that don't prevent execution
- **Medium**: Errors that prevent command execution
- **High**: Critical errors that may affect system stability
- **Critical**: Fatal errors that require immediate attention

#### Best Practices

1. **Always use `shared.HandleError`** for consistent error handling
2. **Provide context** in error messages for better user experience
3. **Use appropriate error codes** for different error types
4. **Handle errors immediately** after they occur
5. **Validate input** before making API calls
6. **Test error scenarios** to ensure proper error handling

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

MIT License - see LICENSE file for details 