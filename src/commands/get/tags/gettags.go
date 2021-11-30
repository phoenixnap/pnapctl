package tags

import (
	netHttp "net/http"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	tagclient "phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"

	"github.com/spf13/cobra"
)

const commandName string = "get tags"

var ID string

var Name string

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
		if len(args) >= 1 {
			ID = args[0]
			return getTags(ID)
		}
		return getTags("")
	},
}

func getTags(tagID string) error {
	var httpResponse *netHttp.Response
	var err error
	var tag tagapisdk.Tag
	var tags []tagapisdk.Tag

	if tagID == "" {
		tags, httpResponse, err = tagclient.Client.TagsGet(Name)
	} else {
		tag, httpResponse, err = tagclient.Client.TagGetById(tagID)
	}

	if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else {
		if tagID == "" {
			return printer.PrintTagListResponse(tags, commandName)
		} else {
			return printer.PrintTagResponse(tag, commandName)
		}
	}
}

func init() {
	GetTagsCmd.Flags().StringVar(&Name, "name", "", "Name of the tag")
	GetTagsCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
