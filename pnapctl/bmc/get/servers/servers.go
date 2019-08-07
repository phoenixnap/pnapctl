package servers

import (
	"errors"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.MainClient.PerformGet("servers")

		if err != nil {
			fmt.Println("Error while requesting servers:", err)
			return errors.New("get-fail")
		}

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Error while reading servers from response:", err)
			return errors.New("read-fail")
		}

		if full {
			err := printer.PrintOutput(body, &[]LongServer{})
		} else {
			err := printer.PrintOutput(body, &[]ShortServer{})
		}

		if err != nil {
			fmt.Println("Error while printing output:", err)
			return err
		}

		return nil
	},
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&full, "full", false, "Shows all server details")
}
