package storagenetworks

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/get/storage-networks/volumes"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

const commandName = "get storage-networks"

func init() {
	utils.SetupOutputFlag(GetStorageNetworksCmd)
	GetStorageNetworksCmd.AddCommand(volumes.GetStorageNetworkVolumesCmd)
}

var GetStorageNetworksCmd = &cobra.Command{
	Use:          "storage-network [ID]",
	Short:        "Retrieve one or all storage networks.",
	Aliases:      []string{"storage-networks"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all storage networks.
	
Prints information about the storage networks.
By default, the data is printed in table format.

To print a specific storage network, an ID needs to be passed as argument.`,
	Example: `
# List all storage networks.
pnapctl get storage-networks [--output <OUTPUT_TYPE>]

# List a specific storage network.
pnapctl get storage-network <ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getStorageNetworksById(args[0])
		}
		return getStorageNetworks()
	},
}

func getStorageNetworks() error {
	storagenetworks, httpResponse, err := networkstorage.Client.NetworkStorageGet()

	if generatedError := utils.CheckForErrors(httpResponse, err); generatedError != nil {
		return generatedError
	} else {
		return printer.PrintStorageNetworkListResponse(storagenetworks)
	}
}

func getStorageNetworksById(id string) error {
	storagenetwork, httpResponse, err := networkstorage.Client.NetworkStorageGetById(id)

	if generatedError := utils.CheckForErrors(httpResponse, err); generatedError != nil {
		return generatedError
	} else {
		return printer.PrintStorageNetworkResponse(storagenetwork)
	}
}
