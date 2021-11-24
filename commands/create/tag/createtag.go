package tag

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/tags"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tagmodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create tag"

// CreateTagCmd is the command for creating a tag.
var CreateTagCmd = &cobra.Command{
	Use:          "tag",
	Short:        "Create a new tag.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new tag.

Requires a file (yaml or json) containing the information needed to create the tag.`,
	Example: `# create a new tag as described in tag.yaml
pnapctl create tag --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

#tagCreate.yaml
name: TagName
description: The description of the tag.
isBillingTag: false
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagCreate, err := tagmodels.CreateTagCreateFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the tag
		response, httpResponse, err := tags.Client.TagPost(*tagCreate)

		if httpResponse != nil && httpResponse.StatusCode != 201 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			return printer.PrintTagResponse(response, commandName)
		}
	},
}

func init() {
	CreateTagCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateTagCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateTagCmd.MarkFlagRequired("filename")
}
