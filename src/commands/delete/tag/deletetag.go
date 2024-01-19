package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

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
	log.Info().Msgf("Deleting Tag with ID [%s].", id)

	result, err := tags.Client.TagDelete(id)
	if err != nil {
		return err
	} else {
		fmt.Println(result.Result, result.TagId)
		return nil
	}
}
