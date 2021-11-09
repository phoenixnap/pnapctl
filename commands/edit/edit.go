package edit

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/edit/quotas"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Modify a resource.",
	Long:  `Modify a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	EditCmd.AddCommand(quotas.EditQuotaCmd)
}
