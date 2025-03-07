package bgppeergroup

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	location string
)

func init() {
	utils.SetupOutputFlag(GetBgpPeerGroupsCmd)

	GetBgpPeerGroupsCmd.Flags().StringVar(&location, "location", "", "Filter by location")
}

var GetBgpPeerGroupsCmd = &cobra.Command{
	Use:          "bgp-peer-group [PUBLIC_NETWORK_ID]",
	Short:        "Retrieve one or all BGP peer groups.",
	Aliases:      []string{"bgp-peer-groups"},
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	Long: `Retrieve one or all BGP peer groups.

Prints detailed information about the BGP peer groups.
By default, the data is printed in table format.

To print a specific BGP peer group, an ID needs to be passed as an argument.`,
	Example: `
# List all BGP peer groups.
pnapctl get bgp-peer-groups [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific public network.
pnapctl get bgp-peer-groups <BGP_PEER_GROUP_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) > 0 {
			return getBgpPeerGroupById(&args[0])
		}
		return getBgpPeerGroups()
	},
}

func getBgpPeerGroups() error {
	log.Info().Msg("Retrieving list of BGP peer groups...")

	bgpPeerGroups, err := networks.Client.BgpPeerGroupsGet(location)

	if err != nil {
		return err
	} else {
		return printer.PrintBgpPeerGroupListResponse(bgpPeerGroups)
	}
}

func getBgpPeerGroupById(id *string) error {
	log.Info().Msgf("Retrieving BGP peer group with ID [%s].", *id)

	bgpPeerGroup, err := networks.Client.BgpPeerGroupGetById(*id)

	if err != nil {
		return err
	} else {
		return printer.PrintBgpPeerGroupResponse(bgpPeerGroup)
	}
}
