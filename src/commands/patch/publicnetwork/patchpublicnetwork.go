package publicnetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Filename string
)

func init() {
	utils.SetupOutputFlag(PatchPublicNetworkCmd)
	utils.SetupFilenameFlag(PatchPublicNetworkCmd, &Filename, utils.UPDATING)
}

var PatchPublicNetworkCmd = &cobra.Command{
	Use:          "public-network [ID]",
	Short:        "Patch a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch a public network.

Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPatch.yaml as request body. 
pnapctl patch server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPatch.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchPublicNetwork(args[0])
	},
}

func patchPublicNetwork(id string) error {
	publicNetworkPatch, err := models.CreateRequestFromFile[networkapi.PublicNetworkModify](Filename)

	if err != nil {
		return err
	}

	response, err := networks.Client.PublicNetworkPatch(id, *publicNetworkPatch)

	if err != nil {
		return err
	} else {
		return printer.PrintPublicNetworkResponse(response)
	}
}
