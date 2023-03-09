package storagenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	storageNetworkID, volumeID string
)

var DeleteStorageNetworkVolumeCmd = &cobra.Command{
	Use:          "volume [storageNetworkID] [volumeID]",
	Short:        "Delete a storage network's volume.",
	Args:         cobra.ExactArgs(2),
	Aliases:      []string{"sn-v"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		storageNetworkID = args[0]
		volumeID = args[1]
		return deleteStorageNetworkVolume()
	},
}

func deleteStorageNetworkVolume() error {
	return networkstorage.Client.NetworkStorageDeleteVolume(storageNetworkID, volumeID)
}