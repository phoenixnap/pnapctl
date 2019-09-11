package shutdown

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var ShutdownCmd = &cobra.Command{
	Use:          "shutdown",
	Short:        "Shuts down a specific server.",
	Long:         "Shuts down a specific server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resource = "servers/" + args[0] + "/actions/shutdown"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			return ctlerrors.GenericFailedRequestError("shutdown")
		}

		return ctlerrors.Result("shutdown").
			IfOk("Shutdown successfully.").
			IfNotFound("Server with ID " + args[0] + " not found.").
			UseResponse(response)
	},
}
