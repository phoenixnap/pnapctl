package bmc

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl"
)

var bmcCmd = &cobra.Command{
	Use:   "bmc",
	Short: "Bare Metal Cloud - Short",
	Long:  "Bare Metal Cloud - Long",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	pnapctl.RootCmd.AddCommand(bmcCmd)
}
