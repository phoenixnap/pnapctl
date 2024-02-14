package pay

import (
	"os"
	"phoenixnap.com/pnapctl/commands/pay/invoicing"
	"github.com/spf13/cobra"
)

var PayCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay a resource.",
	Long:  `Pay a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PayCmd.AddCommand(invoicing.PayInvoiceCmd)
}
