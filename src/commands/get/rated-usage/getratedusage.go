package rated_usage

import (
	month_to_date "phoenixnap.com/pnapctl/commands/get/rated-usage/month-to-date"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"

	"github.com/spf13/cobra"
)

const commandName string = "get rated-usage"

var GetRatedUsageCmd = &cobra.Command{
	Use:          "rated-usage",
	Short:        "Retrieve all rated-usages for the given time period.",
	Aliases:      []string{"rated-usages"},
	SilenceUsage: true,
	Long: `Retrieve all rated-usages for the given time period.

Prints all information about the rated-usages for the given time period.
By default, the data is printed in table format.

Every record corresponds to a charge. All date & times are in UTC.
Note: "from" and "to" are required and need to be in a valid YYYY/MM format.`,
	Example: `
# List all rated usages.
pnapctl get rated-usages --from=2020/10 --to=2021/11 [--category <CATEGORY>] [--output <OUTPUT_TYPE>]
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getRatedUsage()
	},
}

func getRatedUsage() error {
	queryParams, err := billingmodels.NewRatedUsageGetQueryParams(FromYearMonth, ToYearMonth, ProductCategory)
	if err != nil {
		return err
	}

	ratedUsageRecords, httpResponse, err := billing.Client.RatedUsageGet(*queryParams)

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintRatedUsageListResponse(ratedUsageRecords, Full, commandName)
	}
}

var (
	Full            bool
	FromYearMonth   string
	ToYearMonth     string
	ProductCategory string
)

func init() {
	GetRatedUsageCmd.AddCommand(month_to_date.GetRatedUsageMonthToDateCmd)

	utils.SetupOutputFlag(GetRatedUsageCmd)
	utils.SetupFullFlag(GetRatedUsageCmd, &Full, "rated usage")

	GetRatedUsageCmd.Flags().StringVar(&FromYearMonth, "from", "", "From year month (inclusive) to filter rated usage records by.")
	GetRatedUsageCmd.Flags().StringVar(&ToYearMonth, "to", "", "To year month (inclusive) to filter rated usage records by.")
	GetRatedUsageCmd.PersistentFlags().StringVar(&ProductCategory, "category", "", "The product category to filter by.")

	GetRatedUsageCmd.MarkFlagRequired("from")
	GetRatedUsageCmd.MarkFlagRequired("to")
}
