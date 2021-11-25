package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/tags"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
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

		if err != nil {
			return err
		} else if httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println(result.Result, result.TagId)
		return nil
	},
}
