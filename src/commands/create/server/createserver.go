package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var (
	Filename string
	force    bool
)

var Full bool

func init() {
	utils.SetupOutputFlag(CreateServerCmd)
	utils.SetupFullFlag(CreateServerCmd, &Full, "server")
	utils.SetupFilenameFlag(CreateServerCmd, &Filename, utils.CREATION)

	CreateServerCmd.Flags().BoolVar(&force, "force", false, "Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.")
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
pnapctl create server --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>] [--force=false]

# serverCreate.yaml
hostname: "new-server"
description: "New server description"
os: "ubuntu/bionic"
type: "s1.c1.small"
location: "PHX"
sshKeys:
	- "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSUGPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XAt3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/EnmZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbxNrRFi9wrf+M7Q== test1@test"
	- "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test"`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createServer()
	},
}

func createServer() error {
	log.Info().Msg("Creating new Server...")

	serverCreate, err := models.CreateRequestFromFile[bmcapisdk.ServerCreate](Filename)

	if err != nil {
		return err
	}

	// Create the server
	response, err := bmcapi.Client.ServersPost(*serverCreate, force)
	if err != nil {
		return err
	} else {
		return printer.PrintServerResponse(response, Full)
	}
}
