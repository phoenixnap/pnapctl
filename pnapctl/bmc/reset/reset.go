package reset

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

var commandName = "reset"

// ResetCmd is the command for resetting a server.
var ResetCmd = &cobra.Command{
	Use:          "reset",
	Short:        "Resets a specific server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `
Resets a server.

Requires a file (yaml or json) containing the required information to reset the server.`,
	Example: `
# Reset a server
pnapctl bmc reset NDIid939dfkoDd --file=keys.yaml

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
			return ctlerrors.GenericNonRequestError(err.Error(), commandName)
		}

		response, err := client.PerformPost(resource, bytes.NewBuffer(structbyte))

		if err != nil {
			// Generic error with PerformPost
			return ctlerrors.GenericFailedRequestError(commandName)
		}

		return ctlerrors.Result(commandName).
			IfOk("Server reset request sent successfully.").
			IfNotFound("Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
	ResetCmd.Flags().StringVar(&Filename, "file", "", "File containing required information for reset")
	cobra.MarkFlagRequired(ResetCmd.Flags(), "file")
}
