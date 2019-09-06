package reboot

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

const commandName string = "reboot"

var RebootCmd = &cobra.Command{
	Use:          "reboot",
	Short:        "Reboots a specific server.",
	Long:         "Reboots a specific server.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return ctlerrors.InvalidNumberOfArgs(1, len(args), commandName)
		}

		resource := "servers/" + args[0] + "/actions/reboot"
		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			return ctlerrors.GenericFailedRequestError(commandName)
		}

		return ctlerrors.Result(commandName).
			IfOk("Rebooted successfully").
			IfNotFound("Error: Server with ID " + args[0] + " not found.").
			UseResponse(response)
	},
}

func init() {}
