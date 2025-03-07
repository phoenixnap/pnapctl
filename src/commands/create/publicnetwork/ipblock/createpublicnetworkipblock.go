package ipblock

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/rs/zerolog/log"
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
	utils.SetupOutputFlag(CreatePublicNetworkIpBlockCmd)
	utils.SetupFilenameFlag(CreatePublicNetworkIpBlockCmd, &Filename, utils.CREATION)
}

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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return createPublicNetworkIpBlock(args[0])
	},
}

func createPublicNetworkIpBlock(id string) error {
	log.Info().Msgf("Creating new Ip Block on Public Network with ID [%s].", id)

	ipBlock, err := models.CreateRequestFromFile[networkapi.PublicNetworkIpBlockCreate](Filename)

	if err != nil {
		return err
	}

	response, err := networks.Client.PublicNetworkIpBlockPost(id, *ipBlock)

	if err != nil {
		return err
	} else {
		return printer.PrintPublicNetworkIpBlockResponse(response)
	}
}
