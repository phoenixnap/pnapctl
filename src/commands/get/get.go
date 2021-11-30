package get

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/get/clusters"
	"phoenixnap.com/pnapctl/commands/get/events"
	"phoenixnap.com/pnapctl/commands/get/privatenetwork"
	"phoenixnap.com/pnapctl/commands/get/quotas"
	"phoenixnap.com/pnapctl/commands/get/servers"
	"phoenixnap.com/pnapctl/commands/get/sshkeys"
	"phoenixnap.com/pnapctl/commands/get/tags"
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
	GetCmd.AddCommand(tags.GetTagsCmd)
	GetCmd.AddCommand(sshkeys.GetSshKeysCmd)
	GetCmd.AddCommand(privatenetwork.GetPrivateNetworksCmd)
}
