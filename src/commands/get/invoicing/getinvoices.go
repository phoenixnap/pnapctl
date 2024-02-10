package invoicing

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	invoicing "phoenixnap.com/pnapctl/common/client/invoicing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Number string
var Status string
var SentOnFrom string
var SentOnTo string
var Limit int
var Offset int
var SortField string
var SortDirection string

func init() {
	utils.SetupOutputFlag(GetInvoicingCmd)
	GetInvoicingCmd.PersistentFlags().StringVar(&Number, "number", "", "A user-friendly reference number assigned to the invoice.")
	GetInvoicingCmd.PersistentFlags().StringVar(&Status, "status", "", "Payment status of the invoice.")
	GetInvoicingCmd.PersistentFlags().StringVar(&SentOnFrom, "sentOnFrom", "", "Minimum value to filter invoices by sent on date.")
	GetInvoicingCmd.PersistentFlags().StringVar(&SentOnTo, "sentOnTo", "", "Maximum value to filter invoices by sent on date.")
	GetInvoicingCmd.PersistentFlags().IntVar(&Limit, "limit", 0, "The limit of the number of results returned. The number of records returned may be smaller than the limit.")
	GetInvoicingCmd.PersistentFlags().IntVar(&Offset, "offset", 0, "The number of items to skip in the results.")
	GetInvoicingCmd.PersistentFlags().StringVar(&SortField, "sortField", "", "If a sortField is requested, pagination will be done after sorting. Default sorting is by number.")
	GetInvoicingCmd.PersistentFlags().StringVar(&SortDirection, "sortDirection", "", "Sort Given Field depending on the desired direction. Default sorting is descending.")
}

var GetInvoicingCmd = &cobra.Command{
	Use:          "invoice [INVOICE_ID]",
	Short:        "Retrieve one or all invoices for your account.",
	Aliases:      []string{"invoices"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all invoices for your account.
Prints all information about the invoices assigned to your account.
By default, the data is printed in json format.
Table format isn't supported for this command.
To print a specific invoice, an invoice ID needs to be passed as an argument.`,
	Example: `
# List all invoices in.
pnapctl get invoices [--number <NUMBER>] [--status <STATUS>] [--sentOnFrom <SENT_ON_FROM>] [--sentOnTO <SENT_ON_TO>] [--limit <LIMIT>] [--offset <OFFSET>] [--sortField <SORT_FIELD>] [--SortDirection <SORT_DIRECTION>] [--output <OUTPUT_TYPE>]

# List a specific invoice.
pnapctl get invoice <INVOICE_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		printer.OutputFormat = "json"
		if len(args) >= 1 {
			return getInvoicesById(args[0])
		}
		return getInvoices()
	},
}

func getInvoices() error {
	log.Info().Msg("Retrieving list of Invoices...")

	results, err := invoicing.Client.InvoicesGet(Number, Status, SentOnFrom, SentOnTo, Limit, Offset, SortField, SortDirection)

	if err != nil {
		return err
	} else {
		return printer.MainPrinter.PrintOutput(results)
	}
}

func getInvoicesById(id string) error {
	log.Info().Msgf("Retrieving Invoice with ID [%s].", id)

	results, err := invoicing.Client.InvoicesInvoiceIdGet(id)

	if err != nil {
		return err
	} else {
		return printer.MainPrinter.PrintOutput(results)
	}
}