package server

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupFilenameFlag(DeprovisionServerCmd, &Filename, utils.DEPROVISION)
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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deprovisionServer(args[0])
	},
}

func deprovisionServer(id string) error {
	log.Info().Msgf("Deprovisioning Server with ID [%s].", id)

	relinquishIpBlockRequest, err := models.CreateRequestFromFile[bmcapisdk.RelinquishIpBlock](Filename)
	if err != nil {
		return err
	}

	result, err := bmcapi.Client.ServerDeprovision(id, *relinquishIpBlockRequest)
	if err != nil {
		return err
	} else {
		fmt.Println(result)
		return nil
	}
}
