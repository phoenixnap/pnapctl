package tags

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
	Filename string
	Full     bool
)

func init() {
	utils.SetupOutputFlag(UpdateStorageNetworkVolumeTagsCmd)
	utils.SetupFilenameFlag(UpdateStorageNetworkVolumeTagsCmd, &Filename, utils.UPDATING)
}

var UpdateStorageNetworkVolumeTagsCmd = &cobra.Command{
	Use:          "tags",
	Short:        "Update the tags of a storage network volume.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Update the tags of a storage network volume.
	
Requires a file (yaml or json) containing the information needed to update the tags of a storage network volume.`,
	Example: `# Update the tags of a storage network volume as per storageNetworkVolumeTagsUpdate.yaml
pnapctl update storage-network volume tags <STORAGE_NETWORK_ID> <VOLUME_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# storageNetworkVolumeTagsUpdate.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updateStorageNetworkVolumeTags(args[0], args[1])
	},
}

func updateStorageNetworkVolumeTags(storageNetworkId, volumeId string) error {
	log.Info().Msgf("Updating tags for Volume with ID [%s].", volumeId)

	storageNetworkVolumeTagsUpdate, err := models.CreateRequestFromFile[[]networkstorageapi.TagAssignmentRequest](Filename)
	if err != nil {
		return err
	}

	// update the storage network volume's tags
	response, err := networkstorage.Client.NetworkStorageVolumePutTags(storageNetworkId, volumeId, *storageNetworkVolumeTagsUpdate)

	if err != nil {
		return err
	} else {
		return printer.PrintVolumeResponse(response, Full)
	}
}
