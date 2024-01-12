package servers

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
    "github.com/rs/zerolog/log"
)

var Full bool
var tags []string

func init() {
	utils.SetupOutputFlag(GetServersCmd)
	utils.SetupFullFlag(GetServersCmd, &Full, "server")

	GetServersCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
}

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
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getServersById(args[0])
		}
		return getServers()
	},
}

func getServers() error {
	servers, err := bmcapi.Client.ServersGet(tags)

	if err != nil {
		log.Error().Err(err)
		return err
	} else {
		log.Info().Msg("Retrieving list of servers")
		return printer.PrintServerListResponse(servers, Full)
	}
}

func getServersById(serverID string) error {
	server, err := bmcapi.Client.ServerGetById(serverID)

	if err != nil {
		return err
	} else {
		log.Info().Msgf("Retrieving server with [ID] %s", serverID)
		return printer.PrintServerResponse(server, Full)
	}
}
