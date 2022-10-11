package server

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "deprovision server"

func init() {
	DeprovisionServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for deprovision")
	DeprovisionServerCmd.MarkFlagRequired("filename")
}

// DeprovisionServerCmd
var DeprovisionServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Deprovision a server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Deprovision a server.

Requires a file (yaml or json) containing the information needed to deprovision a server.`,
	Example: `# Deprovision a server as per serverdeprovision.yaml
pnapctl deprovision server <SERVER_ID> --filename <FILE_PATH>

# serverdeprovision.yaml
deleteIpBlocks: false`,
	RunE: func(_ *cobra.Command, args []string) error {
		return deprovisionServer(args[0])
	},
}

func deprovisionServer(id string) error {
	relinquishIpBlockRequest, err := models.CreateRequestFromFile[bmcapisdk.RelinquishIpBlock](Filename, commandName)
	if err != nil {
		return err
	}

	result, httpResponse, err := bmcapi.Client.ServerDeprovision(id, *relinquishIpBlockRequest)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result)
		return nil
	}
}
