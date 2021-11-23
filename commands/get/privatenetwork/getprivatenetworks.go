package privatenetwork

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get private-networks"

var Full bool
var ID string
var tags []string

var GetPrivateNetworksCmd = &cobra.Command{
	Use:          "private-network [SERVER_ID]",
	Short:        "Retrieve one or all private networks.",
	Aliases:      []string{"private-networks"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all private networks.

Prints brief or detailed information about the private networks.
By default, the data is printed in table format.

To print a single private network, an ID needs to be passed as an argument.`,
	Example: `
# List all private networks in json format.
pnapctl get private-networks --tag tagName.tagValue --tag tagName -o json

# List all details of a single private network in yaml format.
pnapctl get private-networks NDIid939dfkoDd -o yaml --full`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getPrivateNetworks(ID)
		}
		return getPrivateNetworks("")
	},
}

func getPrivateNetworks(privateNetworkID string) error {
	log.Debug("Getting private networks...")

	var httpResponse *netHttp.Response
	var err error
	var private-network bmcapisdk.PrivateNetwork
	var private-networks []bmcapisdk.PrivateNetwork

	if privateNetworkID == "" {
		private-networks, httpResponse, err = bmcapi.Client.PrivateNetworksGet(tags)
	} else {
		private-network, httpResponse, err = bmcapi.Client.PrivateNetworkGetById(privateNetworkID)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		if privateNetworkID == "" {
			return printer.PrintPrivateNetworkListResponse(private-networks, Full, commandName)
		} else {
			return printer.PrintPrivateNetworkResponse(private-network, Full, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetPrivateNetworksCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all private network details")
	GetPrivateNetworksCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetPrivateNetworksCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
}
