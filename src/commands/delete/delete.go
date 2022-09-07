package delete

import (
	"os"

	ip_blocks "phoenixnap.com/pnapctl/commands/delete/ip-blocks"
	"phoenixnap.com/pnapctl/commands/delete/publicnetwork"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/delete/cluster"
	"phoenixnap.com/pnapctl/commands/delete/server"
	"phoenixnap.com/pnapctl/commands/delete/sshkey"
	"phoenixnap.com/pnapctl/commands/delete/tag"

	"phoenixnap.com/pnapctl/commands/delete/privatenetwork"
	serveripblock "phoenixnap.com/pnapctl/commands/delete/server/ipblocks"
	serverprivatenetwork "phoenixnap.com/pnapctl/commands/delete/server/privatenetwork"
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
	DeleteCmd.AddCommand(serverprivatenetwork.DeleteServerPrivateNetworkCmd)
	DeleteCmd.AddCommand(serveripblock.DeleteServerIpBlockCmd)
	DeleteCmd.AddCommand(cluster.DeleteClusterCmd)
	DeleteCmd.AddCommand(tag.DeleteTagCmd)
	DeleteCmd.AddCommand(sshkey.DeleteSshKeyCmd)
	DeleteCmd.AddCommand(privatenetwork.DeletePrivateNetworkCmd)
	DeleteCmd.AddCommand(publicnetwork.DeletePublicNetworkCmd)
	DeleteCmd.AddCommand(ip_blocks.DeleteIpBlockCmd)
}
