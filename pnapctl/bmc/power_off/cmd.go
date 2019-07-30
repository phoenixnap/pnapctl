package power_off

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var P_OffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Powers off a specific server.",
	Long:  "Powers off a specific server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		os.Exit(0)
	},
}

func init() {
}
