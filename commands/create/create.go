package create

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/create/server"
	privatenetwork "phoenixnap.com/pnap-cli/commands/create/server/private_network"
	sshkey "phoenixnap.com/pnap-cli/commands/create/ssh_key"
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
	CreateCmd.AddCommand(privatenetwork.CreateServerPrivateNetworkCmd)
	CreateCmd.AddCommand(sshkey.CreateSshKeyCmd)
}
