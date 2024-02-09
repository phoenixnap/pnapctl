package provision

import (
	"github.com/spf13/cobra"
	"os"
)

var ProvisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision a resource",
	Long:  `Provision a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ProvisionCmd.AddCommand()
}
