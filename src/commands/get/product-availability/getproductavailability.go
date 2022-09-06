package productavailability

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get product-availability"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		return getProductAvailabilities()
	},
}

func getProductAvailabilities() error {
	queryParams, err := billingmodels.NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	if err != nil {
		return err
	}

	products, httpResponse, err := billing.Client.ProductAvailabilityGet(*queryParams)

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintProductAvailabilityListResponse(products, commandName)
	}
}

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
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&productCode, "code", []string{}, "Category to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().BoolVar(&showOnlyMinQuantityAvailable, "showOnlyMinQuantityAvailable", true, "Category to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&location, "location", []string{}, "Category to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().StringArrayVar(&solution, "solution", []string{}, "Category to filter product availabilities by.")
	GetProductAvailabilitiesCmd.Flags().Float32Var(&minQuantity, "minQuantity", 0.0, "Category to filter product availabilities by.")
}
