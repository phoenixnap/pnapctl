package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "delete tag"

var DeleteTagCmd = &cobra.Command{
	Use:          "tag TAG_ID",
	Short:        "Deletes a specific tag.",
	Long:         "Deletes a specific tag.",
	Example:      `pnapctl delete tag <TAG_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := tags.Client.TagDelete(args[0])

		if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			fmt.Println(result.Result, result.TagId)
			return nil
		}
	},
}
