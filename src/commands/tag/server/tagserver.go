package server

import (
	"net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

const commandName string = "tag server"

var Full bool

// TagServerCmd is the command for tagging a server.
var TagServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Tag a server.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Tag a server.

Requires a file (yaml or json) containing the information needed to tag the server.`,
	Example: `# Tag a server as per serverTag.yaml. 
pnapctl tag server --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverTag.yaml
- name: tagName
  value: tagValue
- name: tagName2
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagRequests, err := models.CreateRequestFromFile[[]bmcapisdk.TagAssignmentRequest](Filename, commandName)
		if err != nil {
			return err
		}

		serverResponse, httpResponse, err := performTagRequest(args[0], *tagRequests)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintServerResponse(serverResponse, Full, commandName)
		}
	},
}

func init() {
	TagServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	TagServerCmd.MarkFlagRequired("filename")
	TagServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	TagServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

func performTagRequest(serverId string, tagRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, *http.Response, error) {
	// An empty array must be used as a request body if file is empty
	if len(tagRequests) < 1 {
		return bmcapi.Client.ServerTag(serverId, []bmcapisdk.TagAssignmentRequest{})
	} else {
		return bmcapi.Client.ServerTag(serverId, tagRequests)
	}
}
