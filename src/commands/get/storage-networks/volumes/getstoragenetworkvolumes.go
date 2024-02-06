package volumes

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Full bool
	Tags []string
)

func init() {
	utils.SetupOutputFlag(GetStorageNetworkVolumesCmd)
	utils.SetupFullFlag(GetStorageNetworkVolumesCmd, &Full, "volume")

	GetStorageNetworkVolumesCmd.PersistentFlags().StringArrayVar(&Tags, "tag", []string{}, "Tags to filter by.")
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
	log.Info().Msgf("Retrieving list of Volumes for Storage Network with ID [%s].", storageId)

	volumes, err := networkstorage.Client.NetworkStorageGetVolumes(storageId, Tags)

	if err != nil {
		return err
	} else {
		return printer.PrintVolumeListResponse(volumes, Full)
	}
}

func getVolumeById(storageId, volumeId string) error {
	log.Info().Msgf("Retrieving Volume with ID [%s] for Storage Network with ID [%s].", volumeId, storageId)

	volume, err := networkstorage.Client.NetworkStorageGetVolumeById(storageId, volumeId)

	if err != nil {
		return err
	} else {
		return printer.PrintVolumeResponse(volume, Full)
	}
}
