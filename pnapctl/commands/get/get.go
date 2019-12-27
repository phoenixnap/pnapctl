package get

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/commands/get/servers"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources.",
	Long:  `Display one or many resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	GetCmd.AddCommand(servers.GetServersCmd)
}