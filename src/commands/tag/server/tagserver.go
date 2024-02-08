package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var Full bool

func init() {
	utils.SetupOutputFlag(TagServerCmd)
	utils.SetupFullFlag(TagServerCmd, &Full, "server")
	utils.SetupFilenameFlag(TagServerCmd, &Filename, utils.TAGGING)
}

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
		cmdname.SetCommandName(cmd)
		return tagServer(args[0])
	},
}

// TODO Look into this weird part.
func performTagRequest(serverId string, tagRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, error) {
	// An empty array must be used as a request body if file is empty
	if len(tagRequests) < 1 {
		return bmcapi.Client.ServerTag(serverId, []bmcapisdk.TagAssignmentRequest{})
	} else {
		return bmcapi.Client.ServerTag(serverId, tagRequests)
	}
}

func tagServer(id string) error {
	log.Info().Msgf("Tagging Server with ID [%s].", id)

	tagRequests, err := models.CreateRequestFromFile[[]bmcapisdk.TagAssignmentRequest](Filename)
	if err != nil {
		return err
	}

	serverResponse, err := performTagRequest(id, *tagRequests)
	if err != nil {
		return err
	} else {
		return printer.PrintServerResponse(serverResponse, Full)
	}
}
