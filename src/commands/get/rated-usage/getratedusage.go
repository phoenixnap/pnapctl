package rated_usage

import (
	"errors"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"

	"github.com/spf13/cobra"
)

const commandName string = "get rated-usage"

var GetRatedUsageCmd = &cobra.Command{
	Use:          "rated-usage",
	Short:        "Retrieve all rated-usages for the given time period.",
	Aliases:      []string{"rated-usages"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve all rated-usages for the given time period.

Prints all information about the rated-usages for the given time period.
By default, the data is printed in table format.

Every record corresponds to a charge. All date & times are in UTC.`,
	Example: `
# List all rated usages.
pnapctl get rated-usages [--output <OUTPUT_TYPE>]

# List all rated usages.
pnapctl get rated-usages [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getRatedUsage()
	},
}

func getRatedUsage() error {
	// var httpResponse *netHttp.Response
	// var err error

	// Only here for import.
	var ratedUsageRecords []billingapisdk.RatedUsageRecord
	return errors.New(ratedUsageRecords[0].Id)

	// if ipBlockId == "" {
	// 	ipBlocks, httpResponse, err = ip.Client.IpBlocksGet()
	// } else {
	// 	ipBlock, httpResponse, err = ip.Client.IpBlocksGetById(ipBlockId)
	// }

	// var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	// if *generatedError != nil {
	// 	return *generatedError
	// } else {
	// 	if ipBlockId == "" {
	// 		return printer.PrintIpBlockListResponse(ipBlocks, commandName)
	// 	} else {
	// 		return printer.PrintIpBlockResponse(ipBlock, commandName)
	// 	}
	// }
}

var FromYearMonth string
var ToYearMonth string
var ProductCategory string

func init() {
	GetRatedUsageCmd.PersistentFlags().StringVarP(&FromYearMonth, "fromYearMonth", "from", "", "From year month (inclusive) to filter rated usage records by.")
	GetRatedUsageCmd.PersistentFlags().StringVarP(&ToYearMonth, "toYearMonth", "to", "", "To year month (inclusive) to filter rated usage records by.")
	GetRatedUsageCmd.PersistentFlags().StringVarP(&ProductCategory, "productCategory", "category", "", "The product category to filter by.")

	GetRatedUsageCmd.MarkPersistentFlagRequired("fromYearMonth")
	GetRatedUsageCmd.MarkPersistentFlagRequired("toYearMonth")
}
