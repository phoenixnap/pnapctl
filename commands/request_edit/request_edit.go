package request_edit

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/request_edit/quotas"
)

var RequestEditCmd = &cobra.Command{
	Use:   "request-edit",
	Short: "Modify a resource.",
	Long:  `Modify a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	RequestEditCmd.AddCommand(quotas.RequestEditQuotaCmd)
}