package publicnetwork

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get public-networks"

var GetPublicNetworksCmd = &cobra.Command{
	Use:          "public-network [PUBLIC_NETWORK_ID]",
	Short:        "Retrieve one or all public networks.",
	Aliases:      []string{"public-networks"},
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	Long: `Retrieve one or all public networks.

Prints detailed information about the public networks.
By default, the data is printed in table format.

To print a specific public network, an ID needs to be passed as an argument.`,
	Example: `
# List all public networks.
pnapctl get public-networks [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific public network.
pnapctl get public-networks <PUBLIC_NETWORK_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) > 0 {
			return getPublicNetworks(&args[0])
		}
		return getPublicNetworks(nil)
	},
}

func getPublicNetworks(id *string) error {
	var httpResponse *http.Response
	var err error
	var publicNetwork *networkapi.PublicNetwork
	var publicNetworks []networkapi.PublicNetwork

	queryParams, err := networkmodels.NewPublicNetworksGetQueryParams(location)

	if err != nil {
		return err
	}

	if id == nil {
		publicNetworks, httpResponse, err = networks.Client.PublicNetworksGet(*queryParams)
	} else {
		publicNetwork, httpResponse, err = networks.Client.PublicNetworkGetById(*id)
	}

	if generatedError := utils.CheckForErrors(httpResponse, err, commandName); *generatedError != nil {
		return *generatedError
	} else {
		if id == nil {
			return printer.PrintPublicNetworkListResponse(publicNetworks, commandName)
		} else {
			return printer.PrintPublicNetworkResponse(publicNetwork, commandName)
		}
	}
}

var (
	location string
)

func init() {
	utils.SetupOutputFlag(GetPublicNetworksCmd)

	GetPublicNetworksCmd.Flags().StringVar(&location, "location", "", "Filter by location")
}
