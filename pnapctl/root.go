package pnapctl

import (
	"fmt"
	"os"

	"phoenixnap.com/pnap-cli/pnapctl/bmc"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pnapctl",
	Short: "Short Desc",
	Long:  "Longer Desc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command, setting flags appropriately.
// Called by main.main(), only needing to happen once.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		var _ = fmt.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	// add flags here when needed
	rootCmd.AddCommand(bmc.BmcCmd)
}
