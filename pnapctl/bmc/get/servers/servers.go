package servers

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/printer"
)

var GetServersCmd = &cobra.Command{
	Use:   "servers",
	Short: "Retrieve one or more servers.",
	Long: `
Retrieve one or more servers.

Prints a table of the most important information about the servers.`,
	Example: `
# List all servers in json format.
pnapctl get servers -o json`,
	Run: func(cmd *cobra.Command, args []string) {
		//do something
		fmt.Println("The output format is currently:", printer.OutputFormat)
	},
}

func init() {}
