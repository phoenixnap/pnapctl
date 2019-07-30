package pnapctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pnapctl",
	Short: "Short Desc",
	Long:  "Longer Desc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var Client = NewHttpClient("http://localhost:8080", 10)

// Execute adds all child commands to the root command, setting flags appropriately.
// Called by main.main(), only needing to happen once.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		var _ = fmt.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	// add flags here when needed
}
