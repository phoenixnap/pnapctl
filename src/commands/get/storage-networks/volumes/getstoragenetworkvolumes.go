package volumes

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

const commandName = "get storage-network volumes"

var (
	Full bool
)

func init() {
	utils.SetupOutputFlag(GetStorageNetworkVolumesCmd)
	utils.SetupFullFlag(GetStorageNetworkVolumesCmd, &Full, "volume")
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
pnapctl get volumes [--full] [--output <OUTPUT_TYPE>]

# List a specific volume.
pnapctl get volume <ID> [--full] [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 2 {
			return getVolumeById(args[0], args[1])
		}
		return getVolumes(args[0])
	},
}

func getVolumes(storageId string) error {
	volumes, httpResponse, err := networkstorage.Client.NetworkStorageGetVolumes(storageId)

	generatedError := utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintVolumeListResponse(volumes, Full)
	}
}

func getVolumeById(storageId, volumeId string) error {
	volume, httpResponse, err := networkstorage.Client.NetworkStorageGetVolumeById(storageId, volumeId)

	generatedError := utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintVolumeResponse(volume, Full)
	}
}
