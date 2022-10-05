package storagenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/models/networkstoragemodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "create storage-network"

var (
	Filename string
)

func init() {
	utils.SetupFilenameFlag(CreateStorageNetworkCmd, &Filename, utils.CREATION)
}

var CreateStorageNetworkCmd = &cobra.Command{
	Use:          "storage-network",
	Short:        "Create a new storage network.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a storage network.
	
Requires a file (yaml or json) containing the information needed to create the storage network.`,
	Example: `# Create a storage network using the contents of storageNetworkCreate.yaml as request body.
pnapctl create storage-network --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storageNetworkCreate.yaml...`, // TODO Add yaml
	RunE: func(_ *cobra.Command, _ []string) error {
		return createStorageNetwork()
	},
}

func createStorageNetwork() error {
	request, err := networkstoragemodels.CreateStorageNetworkCreateFromFile(Filename, commandName)

	if err != nil {
		return err
	}

	sdkResponse, httpResponse, err := networkstorage.Client.NetworkStoragePost(*request)

	if generatedError := utils.CheckForErrors(httpResponse, err, commandName); *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintStorageNetworkResponse(sdkResponse, commandName)
	}
}
