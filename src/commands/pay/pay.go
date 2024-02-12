package pay
// SDK is failing to parse the response because API returns empty string instead of empty object. Commented because it is currenlty being investigated.

// import (
// 	"os"
// 	"phoenixnap.com/pnapctl/commands/pay/invoicing"

// 	"github.com/spf13/cobra"
// )

// var PayCmd = &cobra.Command{
// 	Use:   "pay",
// 	Short: "Pay an invoice.",
// 	Long:  `Pay an invoice.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		cmd.Help()
// 		os.Exit(0)
// 	},
// }

// func init() {
// 	PayCmd.AddCommand(invoicing.PayInvoiceCmd)
// }
