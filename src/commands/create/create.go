package create

import (
	"github.com/spf13/cobra"
	"os"
	"phoenixnap.com/pnapctl/commands/create/cluster"
	"phoenixnap.com/pnapctl/commands/create/ip-blocks"
	"phoenixnap.com/pnapctl/commands/create/privatenetwork"
	"phoenixnap.com/pnapctl/commands/create/server"
	serveripblock "phoenixnap.com/pnapctl/commands/create/server/ipblocks"
	serverprivatenetwork "phoenixnap.com/pnapctl/commands/create/server/privatenetwork"
	"phoenixnap.com/pnapctl/commands/create/sshkey"
	"phoenixnap.com/pnapctl/commands/create/tag"
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
	CreateCmd.AddCommand(serverprivatenetwork.CreateServerPrivateNetworkCmd)
	CreateCmd.AddCommand(serveripblock.CreateServerIpBlockCmd)
	CreateCmd.AddCommand(cluster.CreateClusterCmd)
	CreateCmd.AddCommand(tag.CreateTagCmd)
	CreateCmd.AddCommand(sshkey.CreateSshKeyCmd)
	CreateCmd.AddCommand(privatenetwork.CreatePrivateNetworkCmd)
	CreateCmd.AddCommand(ip_blocks.CreateIpBlockCmd)
}
