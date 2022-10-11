package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"

	"github.com/spf13/cobra"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "reserve server"

var Full bool

func init() {
	ReserveServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	ReserveServerCmd.MarkFlagRequired("filename")
	ReserveServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	ReserveServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

// ResetServerCmd is the command for resetting a server.
var ReserveServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Reserve a specific server.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Reserve a specific server for future use.

Requires a file (yaml or json) containing the information needed to reserve the specific server.`,
	Example: `# Reserve a specific server with pricing model described in serverReserve.yaml
pnapctl reserve server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverReserve.yaml
pricingModel: ONE_MONTH_RESERVATION`,
	RunE: func(_ *cobra.Command, args []string) error {
		return reserveServer(args[0])
	},
}

func reserveServer(id string) error {
	reserveRequest, err := models.CreateRequestFromFile[bmcapisdk.ServerReserve](Filename, commandName)

	if err != nil {
		return err
	}

	serverResponse, httpResponse, err := bmcapi.Client.ServerReserve(id, *reserveRequest)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintServerResponse(serverResponse, Full, commandName)
	}
}
