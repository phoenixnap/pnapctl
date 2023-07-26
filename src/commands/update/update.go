package update

import (
	"os"

	"github.com/spf13/cobra"
	ip_blocks "phoenixnap.com/pnapctl/commands/update/ip-blocks"
	"phoenixnap.com/pnapctl/commands/update/privatenetwork"
	"phoenixnap.com/pnapctl/commands/update/sshkey"
	storagenetworks "phoenixnap.com/pnapctl/commands/update/storage-networks"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a resource.",
	Long:  `Update a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	UpdateCmd.AddCommand(sshkey.UpdateSshKeyCmd)
	UpdateCmd.AddCommand(privatenetwork.UpdatePrivateNetworkCmd)
	UpdateCmd.AddCommand(ip_blocks.UpdateIpBlockCmd)
	UpdateCmd.AddCommand(storagenetworks.UpdateStorageNetworkCmd)
}
