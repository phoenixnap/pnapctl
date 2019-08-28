package reboot

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var RebootCmd = &cobra.Command{
	Use:           "reboot",
	Short:         "Reboots a specific server.",
	Long:          "Reboots a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			return errors.New("args")
		}

		resource := "servers/" + args[0] + "/actions/reboot"
		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			fmt.Println("Error while powering off server:", err)
			return errors.New("client-fail")
		}

		return ctlerrors.Result().
			IfOk("Rebooted successfully").
			IfNotFound("Error: Server with ID " + args[0] + " not found.").
			UseResponse(response)
	},
}

func init() {}
