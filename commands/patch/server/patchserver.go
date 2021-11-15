package server

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

const commandName string = "patch server"

var Full bool

// PatchServerCmd is the command for patching a server.
var PatchServerCmd = &cobra.Command{
	Use:          "server [SERVER_ID]",
	Short:        "Patch a server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch a server.

Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPatch.yaml as request body. 
pnapctl patch server x78sdkjds879sd7cx8 --filename ~/serverPatch.yaml

# serverPatch.yaml
hostname: "patched-server"
description: "My custom server edit"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		patchRequest, err := bmcapimodels.PatchServerRequestFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		serverResponse, httpResponse, err := bmcapi.Client.ServerPatch(args[0], *patchRequest)

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
	PatchServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	PatchServerCmd.MarkFlagRequired("filename")
	PatchServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	PatchServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
