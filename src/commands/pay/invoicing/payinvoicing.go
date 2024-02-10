package invoicing

// import (
// 	"fmt"

// 	invoicing "phoenixnap.com/pnapctl/common/client/invoicing"
// 	"github.com/spf13/cobra"
// 	"github.com/rs/zerolog/log"
// 	"phoenixnap.com/pnapctl/common/utils/cmdname"
// )

// var ID string


// func init() {
// 	PayInvoiceCmd.PersistentFlags().StringVar(&ID, "id", "", "Id of the invoice")

// }

// var PayInvoiceCmd = &cobra.Command{
// 	Use:          "invoice INVOICE_ID",
// 	Short:        "Pay invoice.",
// 	Long:         "Pay invoice.",
// 	Example:      "pnapctl pay invoice <INVOICE_ID>",
// 	Args:         cobra.ExactArgs(1),
// 	Aliases:      []string{"inv"},
// 	SilenceUsage: true,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		cmdname.SetCommandName(cmd)
// 		return payInvoice(args[0])
// 	},
// }

// func payInvoice(id string) error {
// 	log.Info().Msgf("Paying Invoice with ID [%s].", id)

// 	result, err := invoicing.Client.InvoicesInvoiceIdPayPost(id)
// 	if err != nil {
// 		return err
// 	} else {
// 		fmt.Println(result)
// 		return nil
// 	}
// }