package tag

import (
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupOutputFlag(CreateTagCmd)
	utils.SetupFilenameFlag(CreateTagCmd, &Filename, utils.CREATION)
}

// CreateTagCmd is the command for creating a tag.
var CreateTagCmd = &cobra.Command{
	Use:          "tag",
	Short:        "Create a new tag.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new tag.

Requires a file (yaml or json) containing the information needed to create the tag.`,
	Example: `# Create a new tag as described in tagCreate.yaml
pnapctl create tag --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# tagCreate.yaml
name: TagName
description: The description of the tag.
isBillingTag: false
`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createTag()
	},
}

func createTag() error {
	tagCreate, err := models.CreateRequestFromFile[tagapi.TagCreate](Filename)

	if err != nil {
		return err
	}

	// Create the tag
	response, err := tags.Client.TagPost(*tagCreate)
	if err != nil {
		return err
	} else {
		return printer.PrintTagResponse(response)
	}
}
