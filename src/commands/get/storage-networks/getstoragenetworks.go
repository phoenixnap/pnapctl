package storagenetworks

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/get/storage-networks/volumes"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "get storage-networks"

var (
	ID string
)

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
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
		}
		return getStorageNetworks()
	},
}

func getStorageNetworks() error {
	var httpResponse *http.Response
	var err error
	var storagenetwork *networkstorageapi.StorageNetwork
	var storagenetworks []networkstorageapi.StorageNetwork

	if ID == "" {
		storagenetworks, httpResponse, err = networkstorage.Client.NetworkStorageGet()
	} else {
		storagenetwork, httpResponse, err = networkstorage.Client.NetworkStorageGetById(ID)
	}

	if generatedError := utils.CheckForErrors(httpResponse, err, commandName); *generatedError != nil {
		return *generatedError
	} else {
		if ID == "" {
			return printer.PrintStorageNetworkListResponse(storagenetworks, commandName)
		} else {
			return printer.PrintStorageNetworkResponse(storagenetwork, commandName)
		}
	}
}
