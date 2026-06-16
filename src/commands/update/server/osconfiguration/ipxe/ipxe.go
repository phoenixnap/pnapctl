package ipxe

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

func init() {
	utils.SetupOutputFlag(PutServerIpxeCmd)
	utils.SetupFilenameFlag(PutServerIpxeCmd, &Filename, utils.UPDATING)
}

// PutServerIpxeCmd is the command for updating a server's iPXE OS configuration.
var PutServerIpxeCmd = &cobra.Command{
	Use:          "ipxe SERVER_ID",
	Short:        "Update a server's iPXE OS configuration.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update a server's iPXE OS configuration.

Requires a file (yaml or json) containing the information needed to modify the server's iPXE configuration.`,
	Example: `# Update a server's iPXE configuration as per serverIpxeUpdate.yaml
pnapctl update server os-configuration ipxe <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverIpxeUpdate.yaml
url: https://example.com/boot.ipxe
nativeVlanConfiguration:
  vlanId: 10
  staticDhcpAddressV4: 185.74.213.56`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updateServerIpxe(args[0])
	},
}

func updateServerIpxe(id string) error {
	log.Info().Msgf("Updating iPXE OS configuration for Server with ID [%s].", id)

	serverIpxeUpdate, err := models.CreateRequestFromFile[bmcapisdk.OsConfigurationIPXE](Filename)

	if err != nil {
		return err
	}

	// update the server's iPXE configuration
	response, err := bmcapi.Client.ServerOsConfigurationIpxePut(id, *serverIpxeUpdate)
	if err != nil {
		return err
	} else {
		return printer.PrintServerIpxe(response)
	}
}
