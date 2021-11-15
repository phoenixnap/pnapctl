package server

import (
	"net/http"

	"github.com/spf13/cobra"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

const commandName string = "tag server"

var Full bool

// TagServerCmd is the command for tagging a server.
var TagServerCmd = &cobra.Command{
	Use:          "server [SERVER_ID]",
	Short:        "Tag a server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Tag a server.

Requires a file (yaml or json) containing the information needed to tag the server.`,
	Example: `# Tag a server using the contents of serverTag.yaml as request body. 
pnapctl tag server x78sdkjds879sd7cx8 --filename ~/serverTag.yaml

# serverTag.json
[
	{
		name: tagName
		value: tagValue
	},
	{
		name: tagName2
	}
]

#serverTag.yaml
- name: tagName
  value: tagValue
- name: tagName2
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagRequests, err := bmcapimodels.TagServerRequestFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		serverResponse, httpResponse, err := performTagRequest(args[0], tagRequests)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
			return printer.PrintServerResponse(serverResponse, Full, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	TagServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	TagServerCmd.MarkFlagRequired("filename")
	TagServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	TagServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

func performTagRequest(serverId string, tagRequests *[]bmcapisdk.TagAssignmentRequest) (bmcapisdk.Server, *http.Response, error) {
	// An empty array must be used as a request body if file is empty
	if len(*tagRequests) < 1 {
		return bmcapi.Client.ServerTag(serverId, []bmcapisdk.TagAssignmentRequest{})
	} else {
		return bmcapi.Client.ServerTag(serverId, *tagRequests)
	}
}
