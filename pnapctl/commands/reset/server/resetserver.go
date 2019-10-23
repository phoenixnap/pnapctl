package server

import (
	"bytes"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	files "phoenixnap.com/pnap-cli/pnapctl/fileprocessor"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"

	"github.com/spf13/cobra"
)

// Performs a Post request with a body containing a ServerReset struct
//		Field in YAML/JSON must be "sshKeys"
// 		Receives: 200, 400, 404, 500.

// ServerReset is the struct used as the body of the request
// to the "reset" endpoint.
type ServerReset struct {
	SSHKeys []string `json:"sshKeys" yaml:"sshKeys"`
}

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "reset server"

// ResetServerCmd is the command for resetting a server.
var ResetServerCmd = &cobra.Command{
	Use:   "server SERVER_ID",
	Short: "Resets a specific server.",
	Long: `Formats the device storage and re-installs the operating system.
Since SSH keys are not stored, they need to be passed as parameters within a YAML or JSON file.`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Example: `# Reset a server
pnapctl reset server 5da891e90ab0c59bd28e34ad --filename keys.yaml

# keys.yaml
sshKeys:
	- "dkleDileD93lD8a3L"
	- "dkleEILDD93lD8a3L"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		files.ExpandPath(&Filename)
		var resource = "servers/" + args[0] + "/actions/reset"

		data, err := files.ReadFile(Filename)

		if files.IsNotExist(err) {
			return ctlerrors.FileNotExistError(Filename)
		} else if err != nil {
			return ctlerrors.GenericNonRequestError(err.Error(), commandName)
		}

		// Marshal file into JSON using the struct
		var serverReset ServerReset

		structbyte, err := files.UnmarshalToJson(data, &serverReset)

		if err != nil {
			return ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingInFileProcessor, commandName)
		}

		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer(structbyte))

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName)
		}

		return ctlerrors.Result(commandName).
			IfOk("Server reset request sent successfully.").
			IfNotFound("Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
	ResetServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for reset")
	ResetServerCmd.MarkFlagRequired("filename")
}
