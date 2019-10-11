package delete

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

const commandName = "delete"

var DeleteCmd = &cobra.Command{
	Use:          "delete",
	Short:        "Deletes a specific server.",
	Long:         "Deletes a specific server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resource = "servers/" + args[0]
		var response, err = client.MainClient.PerformDelete(resource)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName)
		}

		return ctlerrors.Result(commandName).
			IfOk("Deleted successfully.").
			IfNotFound("Server with ID " + args[0] + " not found").
			UseResponse(response)
	},
}
