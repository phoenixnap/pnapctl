package bmc

import (
	"os"

	"github.com/spf13/cobra"
	poweroff "phoenixnap.com/pnap-cli/pnapctl/bmc/power_off"
	poweron "phoenixnap.com/pnap-cli/pnapctl/bmc/power_on"
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
	BmcCmd.AddCommand(poweron.P_OnCmd)
}
