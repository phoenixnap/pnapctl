package ipblock

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

var CreatePublicNetworkIpBlockCmd = &cobra.Command{
	Use:          "ip-block [NETWORK_ID]",
	Short:        "Create an ip-block on a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create an ip-block on a public network..

Requires a file (yaml or json) containing the information needed to create an ip-block.`,
	Example: `# Create an ip-block using the contents of publicNetworkIpBlockCreate.yaml as request body. 
pnapctl create public-network ip-block <NETWORK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkIpBlockCreate.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	utils.SetupOutputFlag(CreatePublicNetworkIpBlockCmd)
}
