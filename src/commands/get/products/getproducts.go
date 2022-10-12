package products

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	qp "phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	ProductCode     string
	ProductCategory string
	SkuCode         string
	Location        string
)

func init() {
	utils.SetupOutputFlag(GetProductsCmd)

	GetProductsCmd.PersistentFlags().StringVar(&ProductCode, "product-code", "", "Product code to filter products by.")
	GetProductsCmd.PersistentFlags().StringVar(&ProductCategory, "category", "", "Product category to filter products by.")
	GetProductsCmd.PersistentFlags().StringVar(&SkuCode, "sku-code", "", "Sku code to filter products by.")
	GetProductsCmd.PersistentFlags().StringVar(&Location, "location", "", "Location to filter products by.")
}

var GetProductsCmd = &cobra.Command{
	Use:          "product",
	Short:        "Retrieves all products.",
	Aliases:      []string{"products"},
	SilenceUsage: true,
	Long: `Retrieve all products.

Prints all information about products.
By default, the data is printed in table format.`,
	Example: ``,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return getProducts()
	},
}

func getProducts() error {
	queryParams := qp.NewProductsGetQueryParams(ProductCode, ProductCategory, SkuCode, Location)

	products, httpResponse, err := billing.Client.ProductsGet(queryParams)

	if err := utils.CheckErrs(httpResponse, err); err != nil {
		return err
	} else {
		return printer.PrintProductListResponse(products)
	}
}
