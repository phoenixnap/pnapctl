package publicnetwork

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
	utils.SetupOutputFlag(PatchPublicNetworkCmd)
	utils.SetupFilenameFlag(PatchPublicNetworkCmd, &Filename, utils.UPDATING)
}

var PatchPublicNetworkCmd = &cobra.Command{
	Use:          "public-network [ID]",
	Short:        "Patch a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch a public network.

Requires a file (yaml or json) containing the information needed to patch the public network.`,
	Example: `# Patch a public network using the contents of publicNetworkPatch.yaml as request body. 
pnapctl patch public-network <PUBLIC_NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# publicNetworkPatch.yaml
name: Network From CLI (Yaml)
description: This network was updated from the CLI using YAML`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchPublicNetwork(args[0])
	},
}

func patchPublicNetwork(id string) error {
	log.Info().Msgf("Patching Public Network with ID [%s].", id)

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
