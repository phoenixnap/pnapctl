package poweroff

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var P_OffCmd = &cobra.Command{
	Use:           "power-off",
	Short:         "Powers off a specific server.",
	Long:          "Powers off a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			return errors.New("args")
		}

		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering off server:", err)
			return errors.New("client-fail")
		}

		return ctlerrors.Result().
			IfOk("Powered off successfully.").
			IfNotFound("Error: Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}

func init() {
}
