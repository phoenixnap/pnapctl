package privatenetwork

import (
	netHttp "net/http"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/printer"
)

const commandName string = "get private-network"

var ID string
var location string

var GetPrivateNetworksCmd = &cobra.Command{
	Use:          "private-network [PRIVATE_NETWORK_ID]",
	Short:        "Retrieve one or all private networks.",
	Aliases:      []string{"private-networks"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all private networks.

Prints detailed information about the private networks.
By default, the data is printed in table format.

To print a specific private network, an ID needs to be passed as an argument.`,
	Example: `
# List all private networks.
pnapctl get private-networks [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific private network.
pnapctl get private-networks <PRIVATE_NETWORK_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getPrivateNetworks(ID)
		}
		return getPrivateNetworks("")
	},
}

func getPrivateNetworks(privateNetworkID string) error {
	var httpResponse *netHttp.Response
	var err error
	var privateNetwork *networkapisdk.PrivateNetwork
	var privateNetworks []networkapisdk.PrivateNetwork

	if privateNetworkID == "" {
		privateNetworks, httpResponse, err = networks.Client.PrivateNetworksGet(location)
	} else {
		privateNetwork, httpResponse, err = networks.Client.PrivateNetworkGetById(privateNetworkID)
	}

	if httpResponse != nil && httpResponse.StatusCode != 200 {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else {
		if privateNetworkID == "" {
			return printer.PrintPrivateNetworkListResponse(privateNetworks, commandName)
		} else {
			return printer.PrintPrivateNetworkResponse(*privateNetwork, commandName)
		}
	}
}

func init() {
	GetPrivateNetworksCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetPrivateNetworksCmd.PersistentFlags().StringVar(&location, "location", "", "Filter by location")
}
