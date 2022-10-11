package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/tags"
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
	RunE: func(_ *cobra.Command, args []string) error {
		return deleteTag(args[0])
	},
}

func deleteTag(id string) error {
	result, httpResponse, err := tags.Client.TagDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result.Result, result.TagId)
		return nil
	}
}
