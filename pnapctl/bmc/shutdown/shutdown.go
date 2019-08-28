package shutdown

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var ShutdownCmd = &cobra.Command{
	Use:           "shutdown",
	Short:         "Shuts down a specific server.",
	Long:          "Shuts down a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println("Only 1 argument can be passed for 'shutdown' :", len(args), "passed")
			return errors.New("args")
		}

		var resource = "servers/" + args[0] + "/actions/shutdown"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while shutting down server:", err)
			return errors.New("client-fail")
		}

		return ctlerrors.Result().
			IfOk("Shutdown successfully.").
			IfNotFound("Error: Server with ID " + args[0] + " not found.").
			UseResponse(response)
	},
}
