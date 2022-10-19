package tags

import (
	tagclient "phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	"github.com/spf13/cobra"
)

var Name string

func init() {
	utils.SetupOutputFlag(GetTagsCmd)

	GetTagsCmd.Flags().StringVar(&Name, "name", "", "Name of the tag")
}

var GetTagsCmd = &cobra.Command{
	Use:          "tag [TAG_ID]",
	Short:        "Retrieve one or all tags.",
	Aliases:      []string{"tags"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all tags.
	
Prints information about the tags.
By default, the data is printed in table format.

To print a specific tag, an ID needs to be passed as an argument.`,
	Example: `
# List all tags.
pnapctl get tags [--output <OUTPUT_TYPE>]

# List a specific tag.
pnapctl get tag <TAG_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getTagById(args[0])
		}
		return getTags()
	},
}

func getTags() error {
	tags, err := tagclient.Client.TagsGet(Name)

	if err != nil {
		return err
	} else {
		return printer.PrintTagListResponse(tags)
	}
}

func getTagById(tagID string) error {
	tag, err := tagclient.Client.TagGetById(tagID)

	if err != nil {
		return err
	} else {
		return printer.PrintTagResponse(tag)
	}
}
