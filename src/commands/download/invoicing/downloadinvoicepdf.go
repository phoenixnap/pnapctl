package invoicing

import (
	//"fmt"

	invoicing "phoenixnap.com/pnapctl/common/client/invoicing"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/common/fileprocessor"

)

// Destination for downloading the invoice
var Destination string

func init() {
	utils.SetupDestinationFlag(DownloadInvoiceCmd, &Destination, utils.DOWNLOAD)
}

var DownloadInvoiceCmd = &cobra.Command{
	Use:          "invoice INVOICE_ID",
	Short:        "Download invoice.",
	Long:         "Download invoice.",
	Example:      "pnapctl download invoice <INVOICE_ID>",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"inv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return downloadInvoice(args[0])
	},
}

func downloadInvoice(id string) error {
	log.Info().Msgf("Downloading Invoice with ID [%s].", id)

	result, err := invoicing.Client.InvoicesInvoiceIdGeneratePdfPost(id)
	if err != nil {
		return err
	} else {
		fileprocessor.SaveFile(Destination, result)
		return nil
	}
}