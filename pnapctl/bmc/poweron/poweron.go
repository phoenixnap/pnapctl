package poweron

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

const commandName string = "power-on"

var P_OnCmd = &cobra.Command{
	Use:          "power-on",
	Short:        "Powers on a specific server.",
	Long:         "Powers on a specific server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			return ctlerrors.GenericFailedRequestError(commandName)
		}

		return ctlerrors.Result(commandName).
			IfOk("Powered on successfully.").
			IfNotFound("Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
}
