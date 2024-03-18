package storagenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Filename, storageNetworkID string
	Full                       bool
)

func init() {
	utils.SetupOutputFlag(CreateStorageNetworkVolumeCmd)
	utils.SetupFilenameFlag(CreateStorageNetworkVolumeCmd, &Filename, utils.CREATION)
}

var CreateStorageNetworkVolumeCmd = &cobra.Command{
	Use:          "volume [storageNetworkID]",
	Short:        "Create a storage network's volume.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"sn-v"},
	SilenceUsage: true,
	Long: `Create a storage network's volume.
	
Requires a file (yaml or json) containing the information needed to create a storage network's volume.`,
	Example: `# Create a storage network's volume using the contents of storagenetworkvolumecreate.yaml as request body.
pnapctl create storage-network volume <storageNetworkID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storagenetworkvolumecreate.yaml
name: name
description:description
capacityInGb: 2000
pathSuffix: /pathSuffix`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		storageNetworkID = args[0]
		return createStorageNetworkVolume()
	},
}

func createStorageNetworkVolume() error {
	log.Info().Msgf("Creating new Volume for Storage Network with ID [%s].", storageNetworkID)

	request, err := models.CreateRequestFromFile[networkstorageapi.VolumeCreate](Filename)
	if err != nil {
		return err
	}

	sdkResponse, err := networkstorage.Client.NetworkStoragePostVolume(storageNetworkID, *request)

	if err != nil {
		return err
	} else {
		return printer.PrintVolumeResponse(sdkResponse, Full)
	}
}
