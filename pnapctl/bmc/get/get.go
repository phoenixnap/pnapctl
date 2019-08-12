package get

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/pnapctl/printer"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `
Display one or many resources

Prints a table of the most important information about the specified resources.`,
	Example: `
# List all servers in json format.
pnapctl get servers -o json`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	GetCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetCmd.AddCommand(servers.GetServersCmd)
}
