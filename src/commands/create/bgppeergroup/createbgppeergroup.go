package bgppeergroup

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Filename string
)

func init() {
	utils.SetupOutputFlag(CreateBgpPeerGroupCmd)
	utils.SetupFilenameFlag(CreateBgpPeerGroupCmd, &Filename, utils.CREATION)
}

var CreateBgpPeerGroupCmd = &cobra.Command{
	Use:          "bgp-peer-group",
	Short:        "Create a new BGP peer group.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a BGP peer group.

Requires a file (yaml or json) containing the information needed to create the BGP peer group.`,
	Example: `# Create a public network using the contents of bgpPeerGroupCreate.yaml as request body. 
pnapctl create bgp-peer-group --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# bgpPeerGroupCreate.yaml
location: "PHX"
asn: 98239
password: "password"
advertisedRoutes: "DEFAULT"`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createBgpPeerGroup()
	},
}

func createBgpPeerGroup() error {
	log.Info().Msg("Creating new BGP peer group...")

	bgpPeerGroupCreate, err := models.CreateRequestFromFile[networkapi.BgpPeerGroupCreate](Filename)

	if err != nil {
		return err
	}

	response, err := networks.Client.BgpPeerGroupsPost(*bgpPeerGroupCreate)

	if err != nil {
		return err
	} else {
		return printer.PrintBgpPeerGroupResponse(response)
	}
}
