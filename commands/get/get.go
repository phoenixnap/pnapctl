package get

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/get/clusters"
	"phoenixnap.com/pnap-cli/commands/get/events"
	"phoenixnap.com/pnap-cli/commands/get/quotas"
	"phoenixnap.com/pnap-cli/commands/get/servers"
	"phoenixnap.com/pnap-cli/commands/get/sshkeys"
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
	GetCmd.AddCommand(clusters.GetClustersCmd)
	GetCmd.AddCommand(quotas.GetQuotasCmd)
	GetCmd.AddCommand(events.GetEventsCmd)
	GetCmd.AddCommand(sshkeys.GetSshKeysCmd)
}
