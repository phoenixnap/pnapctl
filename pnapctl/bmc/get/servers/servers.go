package servers

import (
	"fmt"
	"io/ioutil"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/printer"
	"phoenixnap.com/pnap-cli/pnapctl/client"

	"github.com/spf13/cobra"
)

type ShortServer struct {
	ID          string `header:"id"`
	Status      string `header:"status"`
	Name        string `header:"name"`
	Description string `header:"description"`
}

type LongServer struct {
	ID          string `header:"id"`
	Status      string `header:"status"`
	Name        string `header:"name"`
	Description string `header:"description"`
	Os          string `header:"os"`
	Type        string `header:"type"`
	Location    string `header:"location"`
	CPU         string `header:"cpu"`
	RAM         string `header:"ram"`
	Storage     string `header:"storage"`
}

var full bool

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
		response, err := client.MainClient.PerformGet("servers")

		if err != nil {
			fmt.Println("Error while requesting servers:", err)
			return
		}

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Error while getting servers:", err)
			return
		}

		if full {
			printer.PrintOutput(body, &[]LongServer{})
		} else {
			printer.PrintOutput(body, &[]ShortServer{})
		}
	},
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&full, "full", false, "Shows all server details")
}
