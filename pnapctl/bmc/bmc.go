package bmc

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/get"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweroff"
)

var BmcCmd = &cobra.Command{
	Use:   "bmc",
	Short: "Bare Metal Cloud - Short",
	Long:  "Bare Metal Cloud - Long",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	BmcCmd.AddCommand(poweroff.P_OffCmd)
	BmcCmd.AddCommand(get.GetCmd)
}
