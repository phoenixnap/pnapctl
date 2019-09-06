package servers

import (
	"io"
	"io/ioutil"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
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
	Use:          "servers",
	Short:        "Retrieve one or all servers.",
	Aliases:      []string{"server"},
	SilenceUsage: true,
	Long: `
Retrieve one or all servers.

Prints brief or detailed information about the servers.
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
		return ctlerrors.GenericFailedRequestError("get server")
	}

	err = ctlerrors.
		Result("get server").
		IfOk("").
		IfNotFound("A server with the ID " + ID + " does not exist.").
		UseResponse(response)

	if err != nil {
		return err
	}

	return performServerGetRequest(response.Body, false)
}

func getAllServers() error {
	response, err := client.MainClient.PerformGet("servers")

	if err != nil {
		return ctlerrors.GenericFailedRequestError("get servers")
	}

	err = ctlerrors.
		Result("get servers").
		IfOk("").
		UseResponse(response)

	if err != nil {
		return err
	}

	return performServerGetRequest(response.Body, true)
}

func performServerGetRequest(responseBody io.Reader, multiple bool) error {
	body, err := ioutil.ReadAll(responseBody)

	if err != nil {
		return ctlerrors.GenericSuccessfulRequestError("ResponseBodyReadFailure", "get server")
	}

	if Full {
		if multiple {
			_, err = printer.MainPrinter.PrintOutput(body, &[]LongServer{})
		} else {
			_, err = printer.MainPrinter.PrintOutput(body, &LongServer{})
		}
	} else {
		if multiple {
			_, err = printer.MainPrinter.PrintOutput(body, &[]ShortServer{})
		} else {
			_, err = printer.MainPrinter.PrintOutput(body, &ShortServer{})
		}
	}

	if err != nil {
		return ctlerrors.GenericSuccessfulRequestError(err.Error(), "get server")
	}

	return nil
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVar(&ID, "id", "", "The ID of the server to retrieve")
}
