package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yourusername/cyverApiCli/cmd/client"
	"github.com/yourusername/cyverApiCli/cmd/shared"
)

// Client command group
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Manage client operations",
	Long:  `Manage client information, projects, findings, assets, and users.`,
}

func init() {
	// Set up shared package with config loader and verbose level
	shared.SetConfigLoader(&ConfigLoader{})
	shared.SetVerboseLevel(verboseLevel)

	// Initialize all client commands using the new compartmentalized structure
	client.InitClientCommands(clientCmd)

	// Add client command to root
	rootCmd.AddCommand(clientCmd)
}
