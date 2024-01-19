package storagenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	volume "phoenixnap.com/pnapctl/commands/create/storage-network/volume"
)

var (
	Filename string
)

func init() {
	utils.SetupFilenameFlag(CreateStorageNetworkCmd, &Filename, utils.CREATION)
	CreateStorageNetworkCmd.AddCommand(volume.CreateStorageNetworkVolumeCmd)
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

# storageNetworkCreate.yaml
name: "CreatedSN"
description: "Description"
location: "PHX"
volumes:
  - name: "VolumeName"
    description: "VDescription"
    pathSuffix: "/cliyaml"
    capacityInGb: 1000`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createStorageNetwork()
	},
}

func createStorageNetwork() error {
	log.Info().Msg("Creating new Storage Network...")

	request, err := models.CreateRequestFromFile[networkstorageapi.StorageNetworkCreate](Filename)

	if err != nil {
		return err
	}

	sdkResponse, err := networkstorage.Client.NetworkStoragePost(*request)

	if err != nil {
		return err
	} else {
		return printer.PrintStorageNetworkResponse(sdkResponse)
	}
}
