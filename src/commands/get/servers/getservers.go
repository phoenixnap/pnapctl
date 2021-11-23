package servers

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get servers"

var Full bool
var ID string
var tags []string

var GetServersCmd = &cobra.Command{
	Use:          "server [SERVER_ID]",
	Short:        "Retrieve one or all servers.",
	Aliases:      []string{"servers", "srv"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all servers.

Prints brief or detailed information about the servers.
By default, the data is printed in table format.

To print a single server, an ID needs to be passed as an argument.`,
	Example: `
# List all servers in json format.
pnapctl get servers --tag tagName.tagValue --tag tagName -o json

# List all details of a single server in yaml format.
pnapctl get servers NDIid939dfkoDd -o yaml --full`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getServers(ID)
		}
		return getServers("")
	},
}

func getServers(serverID string) error {
	var httpResponse *netHttp.Response
	var err error
	var server bmcapisdk.Server
	var servers []bmcapisdk.Server

	if serverID == "" {
		servers, httpResponse, err = bmcapi.Client.ServersGet(tags)
	} else {
		server, httpResponse, err = bmcapi.Client.ServerGetById(serverID)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		if serverID == "" {
			return printer.PrintServerListResponse(servers, Full, commandName)
		} else {
			return printer.PrintServerResponse(server, Full, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetServersCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
}
