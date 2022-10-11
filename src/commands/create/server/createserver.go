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

var commandName = "create server"

var Full bool

func init() {
	CreateServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	CreateServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerCmd.MarkFlagRequired("filename")
}

// CreateServerCmd is the command for creating a server.
var CreateServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "Create a new server.",
	Args:         cobra.ExactArgs(0),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Create a new server.

Requires a file (yaml or json) containing the information needed to create the server.`,
	Example: `# Create a new server as described in serverCreate.yaml
pnapctl create server --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverCreate.yaml
hostname: "new-server"
description: "New server description"
os: "ubuntu/bionic"
type: "s1.c1.small"
location: "PHX"
sshKeys:
	- "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSUGPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XAt3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/EnmZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbxNrRFi9wrf+M7Q== test1@test"
	- "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test"`,
	RunE: func(_ *cobra.Command, _ []string) error {
		return createServer()
	},
}

func createServer() error {
	serverCreate, err := models.CreateRequestFromFile[bmcapisdk.ServerCreate](Filename, commandName)

	if err != nil {
		return err
	}

	// Create the server
	response, httpResponse, err := bmcapi.Client.ServersPost(*serverCreate)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintServerResponse(response, Full, commandName)
	}
}
