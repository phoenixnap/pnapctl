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
	utils.SetupOutputFlag(UpdateServerOsConfigurationIpxeCmd)
	utils.SetupFilenameFlag(UpdateServerOsConfigurationIpxeCmd, &Filename, utils.UPDATING)
}

// UpdateServerOsConfigurationIpxeCmd is the command for updating a server's iPXE OS configuration.
var UpdateServerOsConfigurationIpxeCmd = &cobra.Command{
	Use:          "ipxe SERVER_ID",
	Short:        "Update the iPXE OS configuration of a server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update the iPXE OS configuration of a server.

Requires a file (yaml or json) containing the information needed to modify the server's iPXE configuration.`,
	Example: `# Update the iPXE OS configuration of a server as per serverOsConfigurationIpxeUpdate.yaml
pnapctl update server os-configuration ipxe <SERVER_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# serverOsConfigurationIpxeUpdate.yaml
url: https://example.com/boot.ipxe
nativeVlanConfiguration:
  vlanId: 10
  staticDhcpAddressV4: 185.74.213.56`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updateServerOsConfigurationIpxe(args[0])
	},
}

func updateServerOsConfigurationIpxe(id string) error {
	log.Info().Msgf("Updating iPXE OS configuration for Server with ID [%s].", id)

	serverOsConfigurationIpxeUpdate, err := models.CreateRequestFromFile[bmcapisdk.OsConfigurationIPXE](Filename)

	if err != nil {
		return err
	}

	// update the server's iPXE OS configuration
	response, err := bmcapi.Client.ServerOsConfigurationIpxePut(id, *serverOsConfigurationIpxeUpdate)
	if err != nil {
		return err
	} else {
		return printer.PrintServerOsConfigurationIpxe(response)
	}
}
