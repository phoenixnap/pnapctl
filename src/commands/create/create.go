package create

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/create/cluster"
	ip_blocks "phoenixnap.com/pnapctl/commands/create/ip-blocks"
	"phoenixnap.com/pnapctl/commands/create/privatenetwork"
	"phoenixnap.com/pnapctl/commands/create/publicnetwork"
	"phoenixnap.com/pnapctl/commands/create/reservation"
	"phoenixnap.com/pnapctl/commands/create/server"
	serveripblock "phoenixnap.com/pnapctl/commands/create/server/ipblocks"
	serverprivatenetwork "phoenixnap.com/pnapctl/commands/create/server/privatenetwork"
	serverpublicnetwork "phoenixnap.com/pnapctl/commands/create/server/publicnetwork"
	"phoenixnap.com/pnapctl/commands/create/sshkey"
	storagenetwork "phoenixnap.com/pnapctl/commands/create/storage-network"
	"phoenixnap.com/pnapctl/commands/create/tag"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long:  `Create a resource.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	CreateCmd.AddCommand(server.CreateServerCmd)
	CreateCmd.AddCommand(serverprivatenetwork.CreateServerPrivateNetworkCmd)
	CreateCmd.AddCommand(serverpublicnetwork.CreateServerPublicNetworkCmd)
	CreateCmd.AddCommand(serveripblock.CreateServerIpBlockCmd)
	CreateCmd.AddCommand(cluster.CreateClusterCmd)
	CreateCmd.AddCommand(tag.CreateTagCmd)
	CreateCmd.AddCommand(sshkey.CreateSshKeyCmd)
	CreateCmd.AddCommand(privatenetwork.CreatePrivateNetworkCmd)
	CreateCmd.AddCommand(publicnetwork.CreatePublicNetworkCmd)
	CreateCmd.AddCommand(ip_blocks.CreateIpBlockCmd)
	CreateCmd.AddCommand(reservation.CreateReservationCmd)
	CreateCmd.AddCommand(storagenetwork.CreateStorageNetworkCmd)
}
