package publicnetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/create/publicnetwork/ipblock"
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
	utils.SetupOutputFlag(CreatePublicNetworkCmd)
	utils.SetupFilenameFlag(CreatePublicNetworkCmd, &Filename, utils.CREATION)

	CreatePublicNetworkCmd.AddCommand(ipblock.CreatePublicNetworkIpBlockCmd)
}

var CreatePublicNetworkCmd = &cobra.Command{
	Use:          "public-network",
	Short:        "Create a new public network.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a public network.

Requires a file (yaml or json) containing the information needed to create the public network.`,
	Example: `# Create a public network using the contents of publicNetworkCreate.yaml as request body. 
pnapctl create public-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkCreate.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createPublicNetwork()
	},
}

func createPublicNetwork() error {
	publicNetworkCreate, err := models.CreateRequestFromFile[networkapi.PublicNetworkCreate](Filename)

	if err != nil {
		return err
	}

	response, httpResponse, err := networks.Client.PublicNetworksPost(*publicNetworkCreate)

	var generatedError = utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintPublicNetworkResponse(response)
	}
}
