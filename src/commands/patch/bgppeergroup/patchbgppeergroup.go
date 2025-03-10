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
	utils.SetupOutputFlag(PatchBgpPeerGroupCmd)
	utils.SetupFilenameFlag(PatchBgpPeerGroupCmd, &Filename, utils.UPDATING)
}

var PatchBgpPeerGroupCmd = &cobra.Command{
	Use:          "bgp-peer-group [ID]",
	Short:        "Patch a BGP peer group.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch a BGP peer group.

Requires a file (yaml or json) containing the information needed to patch the BGP peer group.`,
	Example: `# Patch a BGP peer group using the contents of bgpPeerGroupPatch.yaml as request body. 
pnapctl patch bgp-peer-group <BGP_PEER_GROUP_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# bgpPeerGroupPatch.yaml
asn: 98239
password: "password"
advertisedRoutes: "/route"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchBgpPeerGroup(args[0])
	},
}

func patchBgpPeerGroup(id string) error {
	log.Info().Msgf("Patching BGP peer group with ID [%s].", id)

	bgpPeerGroupPatch, err := models.CreateRequestFromFile[networkapi.BgpPeerGroupPatch](Filename)

	if err != nil {
		return err
	}

	response, err := networks.Client.BgpPeerGroupPatchById(id, *bgpPeerGroupPatch)

	if err != nil {
		return err
	} else {
		return printer.PrintBgpPeerGroupResponse(response)
	}
}
