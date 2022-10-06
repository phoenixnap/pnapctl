package tag

import (
	"github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "patch tag"

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
		tagEdit, err := models.CreateRequestFromFile[tagapi.TagUpdate](Filename, commandName)
		if err != nil {
			return err
		}

		tag, httpResponse, err := tags.Client.TagPatch(args[0], *tagEdit)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintTagResponse(tag, commandName)
		}
	},
}

func init() {
	PatchTagCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	PatchTagCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for modification")
	PatchTagCmd.MarkFlagRequired("filename")
}
