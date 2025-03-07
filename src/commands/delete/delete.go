package delete

import (
	"os"

	"phoenixnap.com/pnapctl/commands/delete/bgppeergroup"
	ip_blocks "phoenixnap.com/pnapctl/commands/delete/ip-blocks"
	"phoenixnap.com/pnapctl/commands/delete/publicnetwork"
	storagenetwork "phoenixnap.com/pnapctl/commands/delete/storage-network"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/delete/cluster"
	"phoenixnap.com/pnapctl/commands/delete/server"
	"phoenixnap.com/pnapctl/commands/delete/sshkey"
	"phoenixnap.com/pnapctl/commands/delete/tag"

	"phoenixnap.com/pnapctl/commands/delete/privatenetwork"
	serveripblock "phoenixnap.com/pnapctl/commands/delete/server/ipblocks"
	serverprivatenetwork "phoenixnap.com/pnapctl/commands/delete/server/privatenetwork"
	serverpublicnetwork "phoenixnap.com/pnapctl/commands/delete/server/publicnetwork"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource.",
	Long:  `Delete a resource.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DeleteCmd.AddCommand(server.DeleteServerCmd)
	DeleteCmd.AddCommand(serverprivatenetwork.DeleteServerPrivateNetworkCmd)
	DeleteCmd.AddCommand(serverpublicnetwork.DeleteServerPublicNetworkCmd)
	DeleteCmd.AddCommand(serveripblock.DeleteServerIpBlockCmd)
	DeleteCmd.AddCommand(cluster.DeleteClusterCmd)
	DeleteCmd.AddCommand(tag.DeleteTagCmd)
	DeleteCmd.AddCommand(sshkey.DeleteSshKeyCmd)
	DeleteCmd.AddCommand(privatenetwork.DeletePrivateNetworkCmd)
	DeleteCmd.AddCommand(publicnetwork.DeletePublicNetworkCmd)
	DeleteCmd.AddCommand(bgppeergroup.DeleteBgpPeerGroupCmd)
	DeleteCmd.AddCommand(ip_blocks.DeleteIpBlockCmd)
	DeleteCmd.AddCommand(storagenetwork.DeleteStorageNetworkCmd)
}
