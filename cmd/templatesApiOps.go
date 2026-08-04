package cmd

import "github.com/spf13/cobra"

// Templates command group
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage templates operations",
	Long:  `Manage templates and related operations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)
}
