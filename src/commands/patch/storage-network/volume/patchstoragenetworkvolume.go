package storagenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Filename, storageNetworkID, volumeID string
	Full bool
)

func init() {
	utils.SetupOutputFlag(PatchStorageNetworkVolumeCmd)
	utils.SetupFilenameFlag(PatchStorageNetworkVolumeCmd, &Filename, utils.UPDATING)
}

var PatchStorageNetworkVolumeCmd = &cobra.Command{
	Use:          "volume [storageNetworkID] [volumeID]",
	Short:        "Patch a storage network's volume details.",
	Args:         cobra.ExactArgs(2),
	Aliases:      []string{"sn-v"},
	SilenceUsage: true,
	Long: `Patch a storage network's volume details.
	
Requires a file (yaml or json) containing the information needed to patch the storage network's volume.`,
	Example: `# Patch a storage network's volume using the contents of storagenetworkvolumeupdate.yaml as request body.
pnapctl patch storage-network volume <storageNetworkID> <volumeID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storagenetworkvolumeupdate.yaml
name: name
description:description
capacityInGb: 2000
pathSuffix: /pathSuffix`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		storageNetworkID = args[0]
		volumeID = args[1]
		return patchStorageNetworkVolume()
	},
}

func patchStorageNetworkVolume() error {
	request, err := models.CreateRequestFromFile[networkstorageapi.VolumeUpdate](Filename)
	if err != nil {
		return err
	}

	sdkResponse, err := networkstorage.Client.NetworkStoragePatchVolumeById(storageNetworkID, volumeID, *request)

	if err != nil {
		return err
	} else {
		return printer.PrintVolumeResponse(sdkResponse, Full)
	}
}