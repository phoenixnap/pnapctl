package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

const commandName string = "patch server"

var Full bool

func init() {
	PatchServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	PatchServerCmd.MarkFlagRequired("filename")
	PatchServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	PatchServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

// PatchServerCmd is the command for patching a server.
var PatchServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Patch a server.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Patch a server.

Requires a file (yaml or json) containing the information needed to patch the server.`,
	Example: `# Patch a server using the contents of serverPatch.yaml as request body. 
pnapctl patch server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPatch.yaml
hostname: patched-server
description: My custom server edit`,
	RunE: func(_ *cobra.Command, args []string) error {
		return patchServer(args[0])
	},
}

func patchServer(id string) error {
	patchRequest, err := models.CreateRequestFromFile[bmcapisdk.ServerPatch](Filename, commandName)
	if err != nil {
		return err
	}

	serverResponse, httpResponse, err := bmcapi.Client.ServerPatch(id, *patchRequest)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintServerResponse(serverResponse, Full, commandName)
	}
}
