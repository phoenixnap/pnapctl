package month_to_date

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "get rated-usage month-to-date"

var GetRatedUsageMonthToDateCmd = &cobra.Command{
	Use:          "month-to-date",
	Short:        "Retrieve all rated-usages for the current calendar month.",
	SilenceUsage: true,
	Long: `Retrieve all rated-usages for the current calendar month.
	
Prints all information about the rated-usages for the current month.
By default, the data is printed in a table format.

Every record corresponds to a charge. All dates & times are in UTC.`,
	Example: `
# List all rated-usages	
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getRatedUsageMonthToDate()
	},
}

func getRatedUsageMonthToDate() error {
	queryParams, err := billingmodels.NewRatedUsageGetMonthToDateQueryParams(ProductCategory)
	if err != nil {
		return err
	}

	ratedUsageRecords, httpResponse, err := billing.Client.RatedUsageMonthToDateGet(*queryParams)

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintRatedUsageListResponse(ratedUsageRecords, Full, commandName)
	}
}

var Full bool
var ProductCategory string

func init() {
	GetRatedUsageMonthToDateCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetRatedUsageMonthToDateCmd.PersistentFlags().StringVar(&ProductCategory, "category", "", "The product category to filter by.")
}
