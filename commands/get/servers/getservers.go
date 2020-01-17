package servers

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get servers"

var Full bool
var ID string

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
pnapctl get servers -o json

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
	log.Debug("Getting servers...")

	path := "servers/" + serverID

	response, err := client.MainClient.PerformGet(path)

	if response == nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if response.StatusCode == 200 {
		return printer.PrintServerResponse(response.Body, serverID == "", Full, commandName)
	} else {
		return ctlerrors.HandleResponseError(response, commandName)
	}
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
