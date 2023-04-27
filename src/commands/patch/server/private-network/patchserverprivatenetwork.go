package privatenetwork

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupOutputFlag(PatchServerPrivateNetworkCmd)
	utils.SetupFilenameFlag(PatchServerPrivateNetworkCmd, &Filename, utils.UPDATING)
}

var PatchServerPrivateNetworkCmd = &cobra.Command{
	Use:          "private-network SERVER_ID NETWORK_ID",
	Short:        "Patch a server's private network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Patch a server's private network.
	
Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPrivateNetworkPatch.yaml as the request body.
pnapctl patch server private-network <SERVER_ID> <NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPrivateNetworkPatch.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchServerPrivateNetwork(args[0], args[1])
	},
}

func patchServerPrivateNetwork(serverId string, networkId string) error {
	patchRequest, err := models.CreateRequestFromFile[bmcapisdk.ServerNetworkUpdate](Filename)
	if err != nil {
		return err
	}

	serverPrivateNetworkResponse, err := bmcapi.Client.ServerPrivateNetworkPatch(serverId, networkId, *patchRequest)
	if err != nil {
		return err
	} else {
		return printer.PrintServerPrivateNetwork(serverPrivateNetworkResponse)
	}
}