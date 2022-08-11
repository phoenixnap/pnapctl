package servers

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
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

To print a specific server, an ID needs to be passed as an argument.`,
	Example: `
# List all servers.
pnapctl get servers [--tag <TagName>.<TagValue>] [--tag <TagName>] [--full] [--output <OUTPUT_TYPE>]

# List all specific server.
pnapctl get servers <SERVER_ID> [--full] [--output <OUTPUT_TYPE>]`,
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
	var server *bmcapisdk.Server
	var servers []bmcapisdk.Server

	if serverID == "" {
		servers, httpResponse, err = bmcapi.Client.ServersGet(tags)
	} else {
		server, httpResponse, err = bmcapi.Client.ServerGetById(serverID)
	}

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		if serverID == "" {
			return printer.PrintServerListResponse(servers, Full, commandName)
		} else {
			return printer.PrintServerResponse(server, Full, commandName)
		}
	}
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetServersCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
}
