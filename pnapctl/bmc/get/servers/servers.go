package servers

import (
	"encoding/json"
	"errors"
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

const commandName string = "get servers"

var Full bool
var ID string

var GetServersCmd = &cobra.Command{
	Use:          "servers",
	Short:        "Retrieve one or all servers.",
	Aliases:      []string{"server"},
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long: `
Retrieve one or all servers.

Prints brief or detailed information about the servers.
The format they are printed in is a table by default.

To print a single server, an ID needs to be passed as an argument.`,
	Example: `
# List all servers in json format.
pnapctl bmc get servers -o json

# List a single server in yaml format.
pnapctl bmc get servers --id=NDIid939dfkoDd -o yaml`,
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
		return ctlerrors.GenericFailedRequestError(err, commandName)
	}

	err = ctlerrors.
		Result(commandName).
		IfOk("").
		IfNotFound("A server with the ID " + ID + " does not exist.").
		UseResponse(response)

	if err != nil {
		return err
	}

	return printGetServerResponse(response.Body, false)
}

func getAllServers() error {
	response, err := client.MainClient.PerformGet("servers")

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName)
	}

	err = ctlerrors.
		Result(commandName).
		IfOk("").
		UseResponse(response)

	if err != nil {
		return err
	}

	return printGetServerResponse(response.Body, true)
}

func printGetServerResponse(responseBody io.Reader, multiple bool) error {
	body, err := ioutil.ReadAll(responseBody)

	if err != nil {
		return ctlerrors.GenericNonRequestError(ctlerrors.ResponseBodyReadFailure, commandName)
	}

	if Full {
		if multiple {
			construct := &[]LongServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = printer.MainPrinter.PrintOutput(construct, len(*construct) == 0)
			}
		} else {
			construct := &LongServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = printer.MainPrinter.PrintOutput(construct, false)
			}
		}
	} else {
		if multiple {
			construct := &[]ShortServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = printer.MainPrinter.PrintOutput(construct, len(*construct) == 0)
			}
		} else {
			construct := &ShortServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = printer.MainPrinter.PrintOutput(construct, false)
			}
		}
	}

	// This err is the one outputted within the PrintOutput
	if err != nil {
		return ctlerrors.GenericNonRequestError(err.Error(), commandName)
	}

	return nil
}

// unmarshall will unmarshall a Json byte stream into the provided construct.
func unmarshall(body []byte, construct interface{}) error {
	err := json.Unmarshal(body, &construct)
	if err != nil {
		return errors.New(ctlerrors.UnmarshallingErrorBody)
	}
	return nil
}

func init() {
	GetServersCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetServersCmd.PersistentFlags().StringVar(&ID, "id", "", "The ID of the server to retrieve")
}
