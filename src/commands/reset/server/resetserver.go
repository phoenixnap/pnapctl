package server

import (
	"fmt"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/common/utils"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "reset server"

// ResetServerCmd is the command for resetting a server.
var ResetServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Resets a specific server.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Formats the device storage and re-installs the operating system.

Requires a file (yaml or json) containing the information needed to reset the server.`,
	Example: `# Reset a server with sshKeys as per serverReset.yaml.
pnapctl reset server <SERVER_ID> [--filename <FILE_PATH>]

# serverReset.yaml
installDefaultSshKeys: true
sshKeys:
	- ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSUGPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XAt3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/EnmZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbxNrRFi9wrf+M7Q== test1@test
	- ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test
sshKeyIds: 
	- 5fa54d1e91867c03a0a7b4a4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resetRequest, err := createResetRequest()

		if err != nil {
			return err
		}

		result, httpResponse, err := bmcapi.Client.ServerReset(args[0], *resetRequest)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			fmt.Println(result.Result)
			return err
		}
	},
}

func init() {
	//filename is optional
	ResetServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for reset")
}

func createResetRequest() (*bmcapisdk.ServerReset, error) {
	if len(Filename) < 1 {
		return &bmcapisdk.ServerReset{}, nil
	} else {
		return servermodels.CreateResetRequestFromFile(Filename, commandName)
	}
}
