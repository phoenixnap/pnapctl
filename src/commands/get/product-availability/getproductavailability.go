package productavailability

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	productCategory              []string
	productCode                  []string
	showOnlyMinQuantityAvailable bool
	location                     []string
	solution                     []string
	minQuantity                  float32
)

func init() {
	utils.SetupOutputFlag(GetProductAvailabilitiesCmd)

	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&productCategory, "category", []string{}, "Category to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&productCode, "code", []string{}, "Code to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().BoolVar(&showOnlyMinQuantityAvailable, "showOnlyMinQuantityAvailable", true, "Whether to show only min quantity available. Defaults to true.")
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&location, "location", []string{}, "Location to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&solution, "solution", []string{}, "Solution to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().Float32Var(&minQuantity, "minQuantity", 0.0, "Minimum quantity to filter product availabilities by.")
}

var GetProductAvailabilitiesCmd = &cobra.Command{
	Use:          "product-availabilities",
	Short:        "Retrieve product availabilities",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long:         `Retrieve one or all reservations.`,
	Example: `
# Retrieve all product-availabilities
pnapctl get product-availabilities 
	[--output=<OUTPUT_TYPE>] 
	[--category=<CATEGORY>] 
	[--code=<CODE>] 
	[--showOnlyMinQuantityAvailable=false] 
	[--location=<LOCATION>] 
	[--solution=<SOLUTION>] 
	[--minQuantity=<MIN_QUANTITY>]`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return getProductAvailabilities()
	},
}

func getProductAvailabilities() error {
	log.Info().Msg("Retrieving list of product availabilities...")

	products, err := billing.Client.ProductAvailabilityGet(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	if err != nil {
		return err
	} else {
		return printer.PrintProductAvailabilityListResponse(products)
	}
}
