package publicnetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "patch public-network"

var (
	Filename string
)

func init() {
	utils.SetupOutputFlag(PatchPublicNetworkCmd)

	PatchPublicNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for updating.")
	PatchPublicNetworkCmd.MarkFlagRequired("filename")
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
	RunE: func(_ *cobra.Command, args []string) error {
		return patchPublicNetwork(args[0])
	},
}

func patchPublicNetwork(id string) error {
	publicNetworkPatch, err := models.CreateRequestFromFile[networkapi.PublicNetworkModify](Filename, commandName)

	if err != nil {
		return err
	}

	response, httpResponse, err := networks.Client.PublicNetworkPatch(id, *publicNetworkPatch)

	if generatedError := utils.CheckForErrors(httpResponse, err, commandName); *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintPublicNetworkResponse(response, commandName)
	}
}
