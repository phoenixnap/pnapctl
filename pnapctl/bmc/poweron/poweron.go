package poweron

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var P_OnCmd = &cobra.Command{
	Use:           "power-on",
	Short:         "Powers on a specific server.",
	Long:          "Powers on a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-on':", len(args), "passed")
			return errors.New("args")
		}

		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering on server:", err)
			return errors.New("client-fail")
		}

		return ctlerrors.Result().
			IfOk("Powered on successfully.").
			IfNotFound("Error: Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
}
