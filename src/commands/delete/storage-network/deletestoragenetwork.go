package storagenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	ID string
)

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
	return networkstorage.Client.NetworkStorageDelete(ID)
}
