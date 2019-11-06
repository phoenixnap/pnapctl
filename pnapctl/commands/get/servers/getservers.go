package servers

import (
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/pnapctl/printer"

	"github.com/spf13/cobra"
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
			return getServer(ID)
		}
		return getAllServers()
	},
}

func getServer(serverID string) error {
	response, err := client.MainClient.PerformGet("servers/" + serverID)

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName)
	}

	err = ctlerrors.
		Result(commandName).
		IfOk("").
		IfNotFound("A server with the ID " + ID + " does not exist.").
		UseResponse(response)

	if err != nil {
		return err
	}

	return printer.PrintServerResponse(response.Body, false, Full, commandName)
}

func getAllServers() error {
	response, err := client.MainClient.PerformGet("servers")

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName)
	}

	err = ctlerrors.
		Result(commandName).
		UseResponse(response)

	if err != nil {
		return err
	}

	return printer.PrintServerResponse(response.Body, true, Full, commandName)
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
