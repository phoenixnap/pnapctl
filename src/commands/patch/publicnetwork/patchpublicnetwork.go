package publicnetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

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
		return nil
	},
}

func init() {
	utils.SetupOutputFlag(PatchPublicNetworkCmd)
}
