package download

import (
	"os"
	"phoenixnap.com/pnapctl/commands/download/invoicing"

	"github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a resource.",
	Long:  `Download a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DownloadCmd.AddCommand(invoicing.DownloadInvoiceCmd)
}
