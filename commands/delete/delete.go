package delete

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/delete/server"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource.",
	Long:  `Delete a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DeleteCmd.AddCommand(server.DeleteServerCmd)
}
