package reboot

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var RebootCmd = &cobra.Command{
	Use:          "reboot",
	Short:        "Reboots a specific server.",
	Long:         "Reboots a specific server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		resource := "servers/" + args[0] + "/actions/reboot"
		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			return ctlerrors.GenericFailedRequestError("reboot")
		}

		return ctlerrors.Result("reboot").
			IfOk("Rebooted successfully").
			IfNotFound("Server with ID " + args[0] + " not found.").
			UseResponse(response)
	},
}

func init() {}
