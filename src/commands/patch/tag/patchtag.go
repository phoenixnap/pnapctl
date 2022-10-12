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
	utils.SetupOutputFlag(PatchTagCmd)
	utils.SetupFilenameFlag(PatchTagCmd, &Filename, utils.UPDATING)
}

// PatchTagCmd is the command for creating a server.
var PatchTagCmd = &cobra.Command{
	Use:          "tag TAG_ID",
	Short:        "Patch/Update a tag.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch/Update a tag.

Requires a file (yaml or json) containing the information needed to patch the tag.`,
	Example: `# Modify an existing tag as per tagPatch.yaml
pnapctl patch tag <TAG_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# tagPatch.yaml
name: Tag Name
description: The description of the tag.
isBillingTag: false`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchTag(args[0])
	},
}

func patchTag(id string) error {
	tagEdit, err := models.CreateRequestFromFile[tagapi.TagUpdate](Filename)
	if err != nil {
		return err
	}

	tag, httpResponse, err := tags.Client.TagPatch(id, *tagEdit)
	var generatedError = utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintTagResponse(tag)
	}
}
