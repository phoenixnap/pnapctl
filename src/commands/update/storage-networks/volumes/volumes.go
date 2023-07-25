package volumes

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/update/storage-networks/volumes/tags"
)

func init() {
	UpdateStorageNetworkVolumeCmd.AddCommand(tags.UpdateStorageNetworkVolumeTagsCmd)
}

var UpdateStorageNetworkVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Update a volume.",
	Long:  `Update a volume.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}
