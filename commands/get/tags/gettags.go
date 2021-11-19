package tags

import (
	netHttp "net/http"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	tagclient "phoenixnap.com/pnap-cli/common/client/tags"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"

	log "github.com/sirupsen/logrus"
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

To print a single tag, an ID needs to be passed as an argument.`,
	Example: `
# List all tags in json format.
pnapctl get tags -o json

# List a single tag in yaml format.
pnapctl get tag NDIid939dfkoDd -o yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getTags(ID)
		}
		return getTags("")
	},
}

func getTags(tagID string) error {
	log.Debug("Getting tags...")

	var httpResponse *netHttp.Response
	var err error
	var tag tagapisdk.Tag
	var tags []tagapisdk.Tag

	if tagID == "" {
		tags, httpResponse, err = tagclient.Client.TagsGet(Name)
	} else {
		tag, httpResponse, err = tagclient.Client.TagGetById(tagID)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		if tagID == "" {
			return printer.PrintTagListResponse(tags, commandName)
		} else {
			return printer.PrintTagResponse(tag, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetTagsCmd.Flags().StringVar(&Name, "name", "", "Name of the tag")
	GetTagsCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
