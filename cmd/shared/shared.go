package shared

import (
	"github.com/spf13/viper"
	"github.com/yourusername/cyverApiCli/internal/api/versions"
	"github.com/yourusername/cyverApiCli/internal/api/versions/v2_2"
	"github.com/yourusername/cyverApiCli/internal/errors"
	log "github.com/yourusername/cyverApiCli/logger"
	"github.com/yourusername/cyverApiCli/output"
)

// Global variables that need to be accessible across packages
var (
	VerboseLevel int
)

// SetVerboseLevel sets the global verbose level
func SetVerboseLevel(level int) {
	VerboseLevel = level
}

// GetVerboseLevel returns the current verbose level
func GetVerboseLevel() int {
	return VerboseLevel
}

// VersionedApiClient interface for different API client versions
type VersionedApiClient interface {
	GetPentesterOps() interface{}
}

// CreateVersionedApiClient creates a versioned API client with the given configuration
func CreateVersionedApiClient(apiKey, baseURL, apiVersionString string) (interface{}, error) {
	// Validate inputs
	if apiKey == "" {
		// Check if we have an access token as an alternative to API key
		accessToken := viper.GetString("token.access_token")
		if accessToken == "" {
			return nil, errors.NewCyverError(errors.ErrCodeConfigInvalid, "API key cannot be empty", nil)
		}
		// If we have an access token, we can proceed without an API key
	}
	if baseURL == "" {
		return nil, errors.NewCyverError(errors.ErrCodeConfigInvalid, "base URL cannot be empty", nil)
	}
	if apiVersionString == "" {
		return nil, errors.NewCyverError(errors.ErrCodeConfigInvalid, "API version cannot be empty", nil)
	}

	// Use the client factory from internal/api/versions
	genericClient, err := versions.NewClient(versions.APIVersion(apiVersionString), baseURL, apiKey)
	if err != nil {
		log.GetLogger(VerboseLevel).Error("Error creating API client for version", "apiVersionString", apiVersionString, "error", err)
		return nil, errors.WrapError(err, errors.ErrCodeConfigInvalid, "failed to create API client")
	}

	switch apiVersionString {
	case "v2.2", "latest":
		v2_2Client, ok := genericClient.(*v2_2.Client)
		if !ok {
			log.GetLogger(VerboseLevel).Error("API client for version 'v2.2' is not of expected type *v2_2.Client. Got", "genericClient", genericClient)
			return nil, errors.NewCyverError(errors.ErrCodeUnexpectedType, "API client type mismatch", nil)
		}
		// Set the verbose level for the v2.2 client
		v2_2.SetVerboseLevel(VerboseLevel)
		return v2_2Client, nil
	default:
		log.GetLogger(VerboseLevel).Error("Unsupported API version", "Supported Version", "v2.2, latest", "apiVersionString", apiVersionString)
		return nil, errors.NewCyverError(errors.ErrCodeConfigInvalid, "unsupported API version", nil)
	}
}

// PrintJSONResponse prints a JSON response
func PrintJSONResponse(data interface{}) error {
	return output.PrintJSONResponse(data, VerboseLevel)
}

// PrintCustomTable prints a custom table with specified number of columns
func PrintCustomTable(data interface{}, maxColumns int) error {
	return output.PrintCustomTable(data, maxColumns, VerboseLevel)
}

// PrintPentesterInfoTable prints pentester information in table format
func PrintPentesterInfoTable(data interface{}) error {
	return output.PrintPentesterInfoTable(data, VerboseLevel)
}

// PrintClientsTable prints clients in table format
func PrintClientsTable(data interface{}) error {
	return output.PrintClientsTable(data, VerboseLevel)
}

// PrintPentestersTable prints pentesters in table format
func PrintPentestersTable(data interface{}) error {
	return output.PrintPentestersTable(data, VerboseLevel)
}

// PrintSimpleProjectsList prints a simple projects list
func PrintSimpleProjectsList(data interface{}) error {
	return output.PrintSimpleProjectsList(data, VerboseLevel)
}

// PrintSimpleProjectsTable prints a simple projects table
func PrintSimpleProjectsTable(data interface{}) error {
	return output.PrintSimpleProjectsTable(data, VerboseLevel)
}

// PrintProjectTable prints project information in table format
func PrintProjectTable(data interface{}) error {
	return output.PrintProjectTable(data, VerboseLevel)
}

// PrintSimpleFindingsList prints a simple findings list
func PrintSimpleFindingsList(data interface{}) error {
	return output.PrintSimpleFindingsList(data, VerboseLevel)
}

// PrintSimpleFindingsTable prints a simple findings table
func PrintSimpleFindingsTable(data interface{}) error {
	return output.PrintSimpleFindingsTable(data, VerboseLevel)
}

// PrintFindingTable prints finding information in table format
func PrintFindingTable(data interface{}) error {
	return output.PrintFindingTable(data, VerboseLevel)
}

// PrintChecklistsTable prints checklists in table format
func PrintChecklistsTable(data interface{}) error {
	return output.PrintChecklistsTable(data, VerboseLevel)
}

// PrintComplianceNormsTable prints compliance norms in table format
func PrintComplianceNormsTable(data interface{}) error {
	return output.PrintComplianceNormsTable(data, VerboseLevel)
}

// PrintReportVersionsTable prints report versions in table format
func PrintReportVersionsTable(data interface{}) error {
	return output.PrintReportVersionsTable(data, VerboseLevel)
}

// PrintReportTable prints report information in table format
func PrintReportTable(data interface{}) error {
	return output.PrintReportTable(data, VerboseLevel)
}

// LogError logs an error with the current verbose level
func LogError(message string, args ...interface{}) {
	log.GetLogger(VerboseLevel).Error(message, args...)
}

// LogInfo logs an info message with the current verbose level
func LogInfo(message string, args ...interface{}) {
	log.GetLogger(VerboseLevel).Info(message, args...)
}

// LogDebug logs a debug message with the current verbose level
func LogDebug(message string, args ...interface{}) {
	log.GetLogger(VerboseLevel).Debug(message, args...)
}

// ConfigLoader interface for loading configuration
type ConfigLoader interface {
	LoadConfig() (apiKey, baseURL, apiVersion string, err error)
}

// Global config loader - will be set by the main package
var configLoader ConfigLoader

// SetConfigLoader sets the global config loader
func SetConfigLoader(loader ConfigLoader) {
	configLoader = loader
}

// GetVersionedApiClient gets a versioned API client using the global config loader
func GetVersionedApiClient() interface{} {
	if configLoader == nil {
		LogError("Error: config loader not set")
		return nil
	}

	apiKey, baseURL, apiVersion, err := configLoader.LoadConfig()
	if err != nil {
		LogError("Error: failed to load config", "error", err)
		return nil
	}

	client, err := CreateVersionedApiClient(apiKey, baseURL, apiVersion)
	if err != nil {
		LogError("Error: failed to create versioned API client", "error", err)
		return nil
	}
	return client
}

// GetLogger returns a logger with the current verbose level
func GetLogger() *log.Logger {
	return log.GetLogger(VerboseLevel)
}
