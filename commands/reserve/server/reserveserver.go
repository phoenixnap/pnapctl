package server

import (
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"

	"github.com/spf13/cobra"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "reserve server"

var Full bool

// ResetServerCmd is the command for resetting a server.
var ReserveServerCmd = &cobra.Command{
	Use:   "server SERVER_ID",
	Short: "Reserve a specific server.",
	Long: `Reserve a specific server for future use.
Pricing Model is to be passed within a YAML or JSON file.`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Example: `# Reserve a server
pnapctl reserve server 5da891e90ab0c59bd28e34ad --filename serverReserve.yaml

# serverReserve.yaml
pricingModel: ONE_MONTH_RESERVATION`,
	RunE: func(cmd *cobra.Command, args []string) error {
		reserveRequest, err := bmcapimodels.CreateReserveRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		serverResponse, httpResponse, err := bmcapi.Client.ServerReserve(args[0], *reserveRequest)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
			return printer.PrintServerResponse(serverResponse, Full, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	ReserveServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	ReserveServerCmd.MarkFlagRequired("filename")
	ReserveServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	ReserveServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
