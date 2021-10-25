package server

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	serverCreateModel "phoenixnap.com/pnap-cli/common/models"
	serverModel "phoenixnap.com/pnap-cli/common/models"
	"phoenixnap.com/pnap-cli/common/printer"
)

// CreateServerCmd is the command for creating a server.
var CreateServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "Create a new server.",
	Args:         cobra.ExactArgs(0),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Create a new server.

Requires a file (yaml or json) containing the information needed to create the server.`,
	Example: `# create a new server as described in server.yaml
pnapctl create server --filename ~/server.yaml

# server.yaml
hostname: "new-server"
description: "New server description"
os: "ubuntu/bionic"
type: "s1.c1.small"
location: "PHX"
sshKeys:
	- "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSUGPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XAt3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/EnmZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbxNrRFi9wrf+M7Q== test1@test"
	- "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverCreate, err := serverModel.CreateServerRequestFromFile()

		if err != nil {
			return err
		}

		// Create the server
		response, r, err := bmcapi.Client.ServersPost(*serverCreate)

		if err != nil {
			// TODO - Validate way of processing errors.
			return ctlerrors.GenericFailedRequestError(err, serverCreateModel.CommandName, ctlerrors.ErrorSendingRequest)
		} else if r.StatusCode == 200 {
			return printer.PrintServerResponse(response, false, serverCreateModel.CommandName)
		} else {
			return ctlerrors.HandleBMCError(r, serverCreateModel.CommandName)
		}
	},
}

func init() {
	CreateServerCmd.PersistentFlags().BoolVar(&serverCreateModel.Full, "full", false, "Shows all server details")
	CreateServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerCmd.Flags().StringVarP(&serverCreateModel.Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerCmd.MarkFlagRequired("filename")
}
