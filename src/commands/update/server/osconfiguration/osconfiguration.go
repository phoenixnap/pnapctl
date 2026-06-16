package os_configuration

import (
	"os"

	"github.com/spf13/cobra"

	ipxe "phoenixnap.com/pnapctl/commands/update/server/osconfiguration/ipxe"
)

// UpdateServerOsConfigurationCmd is the parent command for server OS configuration updates.
var UpdateServerOsConfigurationCmd = &cobra.Command{
	Use:   "osconfiguration",
	Short: "Update a server's OS configuration.",
	Long:  `Update a server's OS configuration.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	UpdateServerOsConfigurationCmd.AddCommand(ipxe.PutServerIpxeCmd)
}
