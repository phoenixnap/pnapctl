package publicnetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	location string
)

func init() {
	utils.SetupOutputFlag(GetPublicNetworksCmd)

	GetPublicNetworksCmd.Flags().StringVar(&location, "location", "", "Filter by location")
}

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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) > 0 {
			return getPublicNetworkById(&args[0])
		}
		return getPublicNetworks()
	},
}

func getPublicNetworks() error {
	publicNetworks, err := networks.Client.PublicNetworksGet(location)

	if err != nil {
		return err
	} else {
		return printer.PrintPublicNetworkListResponse(publicNetworks)
	}
}

func getPublicNetworkById(id *string) error {
	publicNetwork, err := networks.Client.PublicNetworkGetById(*id)

	if err != nil {
		return err
	} else {
		return printer.PrintPublicNetworkResponse(publicNetwork)
	}
}
