package storagenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	volume "phoenixnap.com/pnapctl/commands/patch/storage-network/volume"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Filename, ID string
)

func init() {
	utils.SetupOutputFlag(PatchStorageNetworkCmd)
	utils.SetupFilenameFlag(PatchStorageNetworkCmd, &Filename, utils.UPDATING)
	PatchStorageNetworkCmd.AddCommand(volume.PatchStorageNetworkVolumeCmd)
}

var PatchStorageNetworkCmd = &cobra.Command{
	Use:          "storage-network [ID]",
	Short:        "Patch a storage network.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Patch a storage network.
	
Requires a file (yaml or json) containing the information needed to patch the storage network.`,
	Example: `# Patch a storage network using the contents of storageNetworkPatch.yaml as request body.
pnapctl patch storage-network <ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storageNetworkPatch.yaml
name: "UpdatedSN"
description: "Description"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		ID = args[0]
		return patchStorageNetwork()
	},
}

func patchStorageNetwork() error {
	log.Info().Msgf("Patching Storage Network with ID [%s].", ID)

	request, err := models.CreateRequestFromFile[networkstorageapi.StorageNetworkUpdate](Filename)
	if err != nil {
		return err
	}

	sdkResponse, err := networkstorage.Client.NetworkStoragePatch(ID, *request)

	if err != nil {
		return err
	} else {
		return printer.PrintStorageNetworkResponse(sdkResponse)
	}
}
