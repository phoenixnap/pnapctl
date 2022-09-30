package volumes

import (
	"fmt"
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "get storage-network volumes"

var (
	STORAGE_ID, VOLUME_ID string
)

func init() {
	utils.SetupOutputFlag(GetStorageNetworkVolumesCmd)
}

var GetStorageNetworkVolumesCmd = &cobra.Command{
	Use:          "volume [ID]",
	Short:        "Retrieve one or all volumes.",
	Aliases:      []string{"volumes"},
	SilenceUsage: true,
	Args:         cobra.RangeArgs(1, 2),
	Long: `Retrieve one or all volumes.
	
Prints information about the volumes.
By default, the data is printed in table format.

To print a specific volume, an ID needs to be passed as argument.`,
	Example: `
# List all volumes.
pnapctl get volumes [--output <OUTPUT_TYPE>]

# List a specific volume.
pnapctl get volume <ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(_ *cobra.Command, args []string) error {
		STORAGE_ID = args[0]
		if len(args) >= 2 {
			VOLUME_ID = args[1]
		}
		return getStorageNetworkVolumes()
	},
}

func getStorageNetworkVolumes() error {
	var httpResponse *http.Response
	var err error
	var volume *networkstorageapi.Volume
	var volumes []networkstorageapi.Volume

	if VOLUME_ID == "" {
		volumes, httpResponse, err = networkstorage.Client.NetworkStorageGetVolumes(STORAGE_ID)
	} else {
		volume, httpResponse, err = networkstorage.Client.NetworkStorageGetVolumeById(STORAGE_ID, VOLUME_ID)
	}

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	// to silence usage errors
	fmt.Println(volume, volumes)

	if *generatedError != nil {
		return *generatedError
	} else {
		if VOLUME_ID == "" {
			return nil // TODO add printer
		} else {
			return nil // TODO add printer
		}
	}
}
