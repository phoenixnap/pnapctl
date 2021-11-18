package get

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/get/clusters"
	"phoenixnap.com/pnap-cli/commands/get/quotas"
	"phoenixnap.com/pnap-cli/commands/get/servers"
	sshkey "phoenixnap.com/pnap-cli/commands/get/ssh_keys"
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
	GetCmd.AddCommand(sshkey.GetSshKeysCmd)
}
