package transactions

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	payments "phoenixnap.com/pnapctl/common/client/payments"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Limit int
var Offset int
var SortDirection string
var SortField string
var From string
var To string

func init() {
	utils.SetupOutputFlag(GetTransactionsCmd)

	GetTransactionsCmd.PersistentFlags().IntVar(&Limit, "limit", 0, "A 'from' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetTransactionsCmd.PersistentFlags().IntVar(&Offset, "offset", 0, "A 'to' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetTransactionsCmd.PersistentFlags().StringVar(&SortDirection, "sortDirection", 0, "Limit the number of records returned.")
	GetTransactionsCmd.PersistentFlags().StringVar(&SortField, "sortField", "", "Ordering of the event's time. Must be 'ASC' or 'DESC'")
	GetTransactionsCmd.PersistentFlags().StringVar(&From, "from", "", "The username that did the actions.")
	GetTransactionsCmd.PersistentFlags().StringVar(&To, "to", "", "The HTTP verb corresponding to the action. Must be 'POST', 'PUT', 'PATCH', 'DELETE'")
}

var GetTransactionsCmd = &cobra.Command{
	Use:          "transaction [TRANSACTION_ID]",
	Short:        "Retrieve one or all transactions for your account.",
	Aliases:      []string{"transactions"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all transactions for your account.

Prints all information about the transactions assigned to your account.
By default, the data is printed in json format.

Table format isn't supported for this command.

To print a specific transaction, a transaction ID needs to be passed as an argument.`,
	Example: `
# List all transactions in.
pnapctl get transactions [--limit <LIMIT>] [--offset <OFFSET>] [--sortdirection <SORTDIRECTION>] [--sortfield <SORTFIELD>] [--from <FROM>] [--to <TO>] [--output <OUTPUT_TYPE>]

# List a specific transaction.
pnapctl get transactions <TRANSACTION_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		printer.OutputFormat = "json"
		if len(args) >= 1 {
			return getTransactionById(args[0])
		}
		return getTransactions()
	},
}

func getTransactions() error {
	log.Info().Msg("Retrieving list of Transactions...")
	results, err := payments.Client.TransactionsGet(Limit, Offset, SortDirection, SortField, From, To)

	if err != nil {
		return err
	} else {
		return printer.MainPrinter.PrintOutput(results)
	}
}

func getTransactionById(id string) error {
	log.Info().Msgf("Retrieving Transaction with ID [%s].", id)
	results, err := payments.Client.TransactionGetById(id)

	if err != nil {
		return err
	} else {
		return printer.MainPrinter.PrintOutput(results)
	}
}
