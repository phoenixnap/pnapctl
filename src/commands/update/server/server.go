package server

import (
	"os"

	"github.com/spf13/cobra"

	osconfiguration "phoenixnap.com/pnapctl/commands/update/server/osconfiguration"
)

// UpdateServerCmd is the parent command for server update operations.
var UpdateServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Update a server.",
	Long:  `Update a server.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	UpdateServerCmd.AddCommand(osconfiguration.UpdateServerOsConfigurationCmd)
}
