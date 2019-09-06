package poweron

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var P_OnCmd = &cobra.Command{
	Use:          "power-on",
	Short:        "Powers on a specific server.",
	Long:         "Powers on a specific server.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			return ctlerrors.InvalidNumberOfArgs(1, len(args), "power-on")
		}

		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			return ctlerrors.GenericFailedRequestError("power-on")
		}

		return ctlerrors.Result("power-on").
			IfOk("Powered on successfully.").
			IfNotFound("Error: Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
}
