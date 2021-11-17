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

var commandName = "edit tag"

// CreateServerCmd is the command for creating a server.
var EditTagCmd = &cobra.Command{
	Use:          "tag [TAG_ID]",
	Short:        "Submit a tag modification request.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Submit a tag modification request.

Requires a file (yaml or json) containing the information needed to submit the request.`,
	Example: `# modify an existing tag as per tagModificationRequest.yaml
pnapctl edit tag  --filename ~/tagModificationRequest.yaml

# tagModificationRequest.yaml
limit: 75
reason: "My current limit is not enough."`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagEdit, err := tagmodels.CreateTagUpdateFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		// httpResponse, err := bmcapi.Client.TagEditById(args[0], *tagEditRequest)
		tag, httpResponse, err := tags.Client.TagPatch(args[0], *tagEdit)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode != 202 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else {
			fmt.Println("Tag edit successful.")
			return printer.PrintTagResponse(tag, commandName)
		}
	},
}

func init() {
	EditTagCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	EditTagCmd.MarkFlagRequired("filename")
}
