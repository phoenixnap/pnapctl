package storagenetwork

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "patch storage-network"

var (
	Filename, ID string
)

func init() {
	utils.SetupOutputFlag(PatchStorageNetworkCmd)
	utils.SetupFilenameFlag(PatchStorageNetworkCmd, &Filename, utils.UPDATING)
}

var PatchStorageNetworkCmd = &cobra.Command{
	Use:          "storage-network [ID]",
	Short:        "Patch a storage network.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Patch a storage network.
	
Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a storage network using the contents of storageNetworkPatch.yaml as request body.
pnapctl patch storage-network <ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storageNetworkPatch.yaml...`, // TODO: Update YAML
	RunE: func(_ *cobra.Command, args []string) error {
		ID = args[0]
		return patchStorageNetwork()
	},
}

// TODO: Remove
func dummy() (*networkstorageapi.StorageNetworkUpdate, error) { return nil, nil }

func patchStorageNetwork() error {
	request, err := dummy()
	if err != nil {
		return err
	}

	sdkResponse, httpResponse, err := networkstorage.Client.NetworkStoragePatch(ID, *request)

	// for silencing usage errors
	fmt.Println(sdkResponse)

	if generatedError := utils.CheckForErrors(httpResponse, err, commandName); *generatedError != nil {
		return *generatedError
	} else {
		return nil // TODO add printer
	}
}
