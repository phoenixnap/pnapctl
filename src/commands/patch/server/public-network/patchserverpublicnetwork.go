package publicnetwork

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var force bool

func init() {
	utils.SetupOutputFlag(PatchServerPublicNetworkCmd)
	utils.SetupFilenameFlag(PatchServerPublicNetworkCmd, &Filename, utils.UPDATING)
	PatchServerPublicNetworkCmd.PersistentFlags().BoolVar(&force, "force", false, "Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.")
}

var PatchServerPublicNetworkCmd = &cobra.Command{
	Use:          "public-network SERVER_ID NETWORK_ID",
	Short:        "Patch a server's public network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Patch a server's public network.
	
Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPublicNetworkPatch.yaml as the request body.
pnapctl patch server public-network <SERVER_ID> <NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>] [--force]

# serverPublicNetworkPatch.yaml
ips:
  - "10.0.0.0"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchServerPublicNetwork(args[0], args[1])
	},
}

func patchServerPublicNetwork(serverId string, networkId string) error {
	log.Info().Msgf("Patching Public Network with ID [%s] for Server with ID [%s].", networkId, serverId)

	patchRequest, err := models.CreateRequestFromFile[bmcapisdk.ServerNetworkUpdate](Filename)
	if err != nil {
		return err
	}

	serverPublicNetworkResponse, err := bmcapi.Client.ServerPublicNetworkPatch(serverId, networkId, *patchRequest, force)
	if err != nil {
		return err
	} else {
		return printer.PrintServerPublicNetwork(serverPublicNetworkResponse)
	}
}
