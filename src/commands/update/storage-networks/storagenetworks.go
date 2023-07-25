package storagenetworks

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/update/storage-networks/volumes"
)

func init() {
	UpdateStorageNetworkCmd.AddCommand(volumes.UpdateStorageNetworkVolumeCmd)
}

var UpdateStorageNetworkCmd = &cobra.Command{
	Use:   "storage-network",
	Short: "Update a storage network.",
	Long:  `Update a storage network.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}
