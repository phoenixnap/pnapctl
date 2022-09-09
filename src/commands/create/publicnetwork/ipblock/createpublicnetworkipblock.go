package ipblock

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "create public-network ip-block"

var CreatePublicNetworkIpBlockCmd = &cobra.Command{
	Use:          "ip-block [NETWORK_ID]",
	Short:        "Create an ip-block on a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create an ip-block on a public network.

Requires a file (yaml or json) containing the information needed to create an ip-block.`,
	Example: `# Create an ip-block using the contents of publicNetworkIpBlockCreate.yaml as request body. 
pnapctl create public-network ip-block <NETWORK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkIpBlockCreate.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(_ *cobra.Command, args []string) error {
		ipBlock, err := networkmodels.CreatePublicNetworkIpBlockFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := networks.Client.PublicNetworkIpBlockPost(args[0], *ipBlock)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintPublicNetworkIpBlockResponse(response, commandName)
		}
	},
}

var (
	Filename string
)

func init() {
	utils.SetupOutputFlag(CreatePublicNetworkIpBlockCmd)
}
