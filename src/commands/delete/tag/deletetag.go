package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
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
		cmdname.SetCommandName(cmd)
		return deleteTag(args[0])
	},
}

func deleteTag(id string) error {
	result, httpResponse, err := tags.Client.TagDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		fmt.Println(result.Result, result.TagId)
		return nil
	}
}
