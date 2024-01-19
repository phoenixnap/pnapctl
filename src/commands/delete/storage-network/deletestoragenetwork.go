package storagenetwork

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	volume "phoenixnap.com/pnapctl/commands/delete/storage-network/volume"
)

var (
	ID string
)

func init() {
	DeleteStorageNetworkCmd.AddCommand(volume.DeleteStorageNetworkVolumeCmd)
}

var DeleteStorageNetworkCmd = &cobra.Command{
	Use:          "storage-network [ID]",
	Short:        "Deletes a specific storage network.",
	Long:         "Deletes a specific storage network.",
	Example:      `pnapctl delete storage-network <ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		ID = args[0]
		return deleteStorageNetwork()
	},
}

func deleteStorageNetwork() error {
	log.Info().Msgf("Deleting Storage Network with ID [%s].", ID)

	return networkstorage.Client.NetworkStorageDelete(ID)
}
