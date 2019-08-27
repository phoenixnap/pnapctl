package servers

import (
	"errors"
	"fmt"
	"io/ioutil"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/printer"

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

var Full bool
var ID string

var GetServersCmd = &cobra.Command{
	Use:           "servers",
	Short:         "Retrieve one or all servers.",
	Aliases:       []string{"server"},
	SilenceErrors: true,
	SilenceUsage:  true,
	Long: `
Retrieve one or all servers.

Prints the most important information about the servers.
The format they are printed in is a table by default.

To print a single server, an ID needs to be passed as an argument.`,
	Example: `
# List all servers in json format.
pnapctl get servers -o json

# List a single server in yaml format.
pnapctl get servers --id=NDIid939dfkoDd -o yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ID != "" {
			return getServer(ID)
		}
		return getAllServers()
	},
}

func getServer(serverID string) error {
	response, err := client.MainClient.PerformGet("servers/" + serverID)

	if err != nil {
		fmt.Println("Error while requesting a server:", err)
		return errors.New("get-fail")
	}

	if response.StatusCode == 404 {
		fmt.Println("A server with the ID", ID, "does not exist.")
		return errors.New("404")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading server from response:", err)
		return errors.New("read-fail")
	}

	if Full {
		_, err = printer.MainPrinter.PrintOutput(body, &LongServer{})
	} else {
		_, err = printer.MainPrinter.PrintOutput(body, &ShortServer{})
	}

	if err != nil {
		fmt.Println("Error while printing output:", err)
		return err
	}

	return nil
}

func getAllServers() error {
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

	if Full {
		_, err = printer.MainPrinter.PrintOutput(body, &[]LongServer{})
	} else {
		_, err = printer.MainPrinter.PrintOutput(body, &[]ShortServer{})
	}

	if err != nil {
		fmt.Println("Error while printing output:", err)
		return err
	}

	return nil
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVar(&ID, "id", "", "The ID of the server to retrieve")
}
