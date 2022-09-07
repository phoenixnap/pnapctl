package publicnetwork

import (
	"github.com/spf13/cobra"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	utils.SetupOutputFlag(CreatePublicNetworkCmd)
}
