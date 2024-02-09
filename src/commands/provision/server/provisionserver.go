package server

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupFilenameFlag(ProvisionServerCmd, &Filename, utils.PROVISION)
}

// ProvisionServerCmd
var ProvisionServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Provision a server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Deprovision a server.

Requires a file (yaml or json) containing the information needed to provision a server.`,
	Example: `# Provision a server as per serverprovision.yaml
pnapctl deprovision server <SERVER_ID> --filename <FILE_PATH>

# serverprovision.yaml
hostname: my-server-1
description: 'My custom server #1'
os: ubuntu/bionic
installDefaultSshKeys: true
sshKeys:
  - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDF9LdAFElNCi7JoWh6KUcchrJ2Gac1aqGRPpdZNowObpRtmiRCecAMb7bUgNAaNfcmwiQi7tos9TlnFgprIcfMWb8MSs3ABYHmBgqEEt3RWYf0fAc9CsIpJdMCUG28TPGTlRXCEUVNKgLMdcseAlJoGp1CgbHWIN65fB3he3kAZcfpPn5mapV0tsl2p+ZyuAGRYdn5dJv2RZDHUZBkOeUobwsij+weHCKAFmKQKtCP7ybgVHaQjAPrj8MGnk1jBbjDt5ws+Be+9JNjQJee9zCKbAOsIo3i+GcUIkrw5jxPU/RTGlWBcemPaKHdciSzGcjWboapzIy49qypQhZe1U75
    user@my_ip
sshKeyIds:
  - 5fa54d1e91867c03a0a7b4a4
networkType: PUBLIC_AND_PRIVATE
networkConfiguration:
  gatewayAddress: 182.16.0.145
  privateNetworkConfiguration:
    configurationType: USER_DEFINED
    privateNetworks:
      - id: 60f81608e2f4665962b214db
        ips:
          - 10.0.0.11
          - 10.0.0.12
        dhcp: false
  publicNetworkConfiguration:
    publicNetworks:
      - id: 60473c2509268bc77fd06d29
        ips:
          - 182.16.0.146
          - 182.16.0.147
  ipBlocksConfiguration:
    configurationType: USER_DEFINED
    ipBlocks:
      - id: 60473a6115e34466c9f8f083`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return provisionServer(args[0])
	},
}

func provisionServer(id string) error {
	log.Info().Msgf("Provisioning Server with ID [%s].", id)

	var request *bmcapisdk.ServerProvision = &bmcapisdk.ServerProvision{}
	var err error

	if Filename != "" {
		request, err = models.CreateRequestFromFile[bmcapisdk.ServerProvision](Filename)
		if err != nil {
			return err
		}
	}

	result, err := bmcapi.Client.ServerProvision(id, *request)
	if err != nil {
		return err
	} else {
		fmt.Println(result)
		return nil
	}
}
