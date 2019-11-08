package create

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/commands/create/server"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long:  `Create a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	CreateCmd.AddCommand(server.CreateServerCmd)
}
