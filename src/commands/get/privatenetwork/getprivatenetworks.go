package privatenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var location string

func init() {
	GetPrivateNetworksCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetPrivateNetworksCmd.PersistentFlags().StringVar(&location, "location", "", "Filter by location")
}

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
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getPrivateNetworkById(args[0])
		}
		return getPrivateNetworks()
	},
}

func getPrivateNetworks() error {
	privateNetworks, httpResponse, err := networks.Client.PrivateNetworksGet(location)

	if httpResponse != nil && httpResponse.StatusCode != 200 {
		return ctlerrors.HandleBMCError(httpResponse)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)
	} else {
		return printer.PrintPrivateNetworkListResponse(privateNetworks)
	}
}

func getPrivateNetworkById(privateNetworkID string) error {
	privateNetwork, httpResponse, err := networks.Client.PrivateNetworkGetById(privateNetworkID)

	if httpResponse != nil && httpResponse.StatusCode != 200 {
		return ctlerrors.HandleBMCError(httpResponse)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)
	} else {
		return printer.PrintPrivateNetworkResponse(privateNetwork)
	}
}
