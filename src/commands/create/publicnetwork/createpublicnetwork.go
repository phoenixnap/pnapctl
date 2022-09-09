package publicnetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/create/publicnetwork/ipblock"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "create public-network"

var CreatePublicNetworkCmd = &cobra.Command{
	Use:          "public-network",
	Short:        "Create a public network.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a public network.

Requires a file (yaml or json) containing the information needed to create the public network.`,
	Example: `# Create a public network using the contents of publicNetworkCreate.yaml as request body. 
pnapctl create public-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkCreate.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(_ *cobra.Command, _ []string) error {
		publicNetworkCreate, err := networkmodels.CreatePublicNetworkCreateFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := networks.Client.PublicNetworksPost(*publicNetworkCreate)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintPublicNetworkResponse(response, commandName)
		}
	},
}

var (
	Filename string
)

func init() {
	utils.SetupOutputFlag(CreatePublicNetworkCmd)

	CreatePublicNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation.")
	CreatePublicNetworkCmd.MarkFlagRequired("filename")

	CreatePublicNetworkCmd.AddCommand(ipblock.CreatePublicNetworkIpBlockCmd)
}
