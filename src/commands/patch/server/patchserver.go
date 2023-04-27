package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	privatenetwork "phoenixnap.com/pnapctl/commands/patch/server/private-network"
	publicnetwork "phoenixnap.com/pnapctl/commands/patch/server/public-network"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var Full bool

func init() {
	utils.SetupOutputFlag(PatchServerCmd)
	utils.SetupFullFlag(PatchServerCmd, &Full, "server")
	utils.SetupFilenameFlag(PatchServerCmd, &Filename, utils.UPDATING)

	PatchServerCmd.AddCommand(privatenetwork.PatchServerPrivateNetworkCmd)
	PatchServerCmd.AddCommand(publicnetwork.PatchServerPublicNetworkCmd)
}

// PatchServerCmd is the command for patching a server.
var PatchServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Patch a server.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Patch a server.

Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPatch.yaml as request body. 
pnapctl patch server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPatch.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchServer(args[0])
	},
}

func patchServer(id string) error {
	patchRequest, err := models.CreateRequestFromFile[bmcapisdk.ServerPatch](Filename)
	if err != nil {
		return err
	}

	serverResponse, err := bmcapi.Client.ServerPatch(id, *patchRequest)
	if err != nil {
		return err
	} else {
		return printer.PrintServerResponse(serverResponse, Full)
	}
}
