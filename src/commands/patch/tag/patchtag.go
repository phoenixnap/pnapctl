package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/tags"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tagmodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "patch tag"

// CreateServerCmd is the command for creating a server.
var PatchTagCmd = &cobra.Command{
	Use:          "tag TAG_ID",
	Short:        "Patch/Update a tag.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch/Update a tag.

Requires a file (yaml or json) containing the information needed to patch the tag.`,
	Example: `# modify an existing tag as per tagPatch.yaml
pnapctl patch tag <TAG_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# tagPatch.yaml
name: Tag Name
description: The description of the tag.
isBillingTag: false`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagEdit, err := tagmodels.CreateTagUpdateFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		tag, httpResponse, err := tags.Client.TagPatch(args[0], *tagEdit)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else {
			fmt.Println("Tag edit successful.")
			return printer.PrintTagResponse(tag, commandName)
		}
	},
}

func init() {
	PatchTagCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	PatchTagCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for modification")
	PatchTagCmd.MarkFlagRequired("filename")
}
