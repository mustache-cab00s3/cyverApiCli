package client

import (
	"github.com/spf13/cobra"
)

// Command group variables for compartmentalization
var (
	projectsCmd *cobra.Command
	findingsCmd *cobra.Command
	assetsCmd   *cobra.Command
	usersCmd    *cobra.Command
)

// InitClientCommands initializes all client commands and adds them to the main client command
func InitClientCommands(clientCmd *cobra.Command) {
	// Initialize all command groups
	initProjectsCommands()
	initFindingsCommands()
	initAssetsCommands()
	initUsersCommands()

	// Add all command groups to the main client command
	clientCmd.AddCommand(projectsCmd)
	clientCmd.AddCommand(findingsCmd)
	clientCmd.AddCommand(assetsCmd)
	clientCmd.AddCommand(usersCmd)

	// Initialize continuous projects commands (handled separately)
	InitContinuousProjectsCommands(clientCmd)
}

// initProjectsCommands initializes projects-related commands
func initProjectsCommands() {
	// Create projects command group
	projectsCmd = &cobra.Command{
		Use:   "projects",
		Short: "Manage projects",
		Long:  `Manage client projects including listing, getting details, and requesting new projects.`,
	}

	// Add project commands to projects command group
	projectsCmd.AddCommand(getProjectsCmd)
	projectsCmd.AddCommand(getProjectByIDCmd)
	projectsCmd.AddCommand(getProjectRequestFormsCmd)
	projectsCmd.AddCommand(requestProjectCmd)
}

// initFindingsCommands initializes findings-related commands
func initFindingsCommands() {
	// Create findings command group
	findingsCmd = &cobra.Command{
		Use:   "findings",
		Short: "Manage findings",
		Long:  `Manage client findings including listing, getting details, and updating status.`,
	}

	// Add findings commands to findings command group
	findingsCmd.AddCommand(getFindingsCmd)
	findingsCmd.AddCommand(getFindingByIDCmd)
	findingsCmd.AddCommand(setFindingStatusCmd)
}

// initAssetsCommands initializes assets-related commands
func initAssetsCommands() {
	// Create assets command group
	assetsCmd = &cobra.Command{
		Use:   "assets",
		Short: "Manage assets",
		Long:  `Manage client assets including listing, creating, updating, and deleting assets.`,
	}

	// Add assets commands to assets command group
	assetsCmd.AddCommand(getAssetsCmd)
	assetsCmd.AddCommand(createAssetCmd)
	assetsCmd.AddCommand(deleteAssetCmd)
	assetsCmd.AddCommand(updateAssetCmd)
}

// initUsersCommands initializes users-related commands
func initUsersCommands() {
	// Create users command group
	usersCmd = &cobra.Command{
		Use:   "users",
		Short: "Manage users",
		Long:  `Manage client users including listing and creating users.`,
	}

	// Add users commands to users command group
	usersCmd.AddCommand(getUsersCmd)
	usersCmd.AddCommand(createUserCmd)
}
