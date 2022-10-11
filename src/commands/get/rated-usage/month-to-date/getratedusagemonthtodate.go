package month_to_date

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	qp "phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

const commandName string = "get rated-usage month-to-date"

var (
	Full            bool
	ProductCategory string
)

func init() {
	utils.SetupOutputFlag(GetRatedUsageMonthToDateCmd)
	utils.SetupFullFlag(GetRatedUsageMonthToDateCmd, &Full, "rated-usage")

	GetRatedUsageMonthToDateCmd.PersistentFlags().StringVar(&ProductCategory, "category", "", "The product category to filter by.")
}

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
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return getRatedUsageMonthToDate()
	},
}

func getRatedUsageMonthToDate() error {
	queryParams, err := qp.NewRatedUsageGetMonthToDateQueryParams(ProductCategory)
	if err != nil {
		return err
	}

	ratedUsageRecords, httpResponse, err := billing.Client.RatedUsageMonthToDateGet(*queryParams)

	generatedError := utils.CheckForErrors(httpResponse, err)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintRatedUsageListResponse(ratedUsageRecords, Full)
	}
}
