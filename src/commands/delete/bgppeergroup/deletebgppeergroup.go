package bgppeergroup

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeleteBgpPeerGroupCmd = &cobra.Command{
	Use:          "bgp-peer-group [ID]",
	Short:        "Deletes a BGP peer group.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long:         `Delete a BGP peer group.`,
	Example: `# Delete a BGP peer group
pnapctl delete BGP peer group <ID>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteBgpPeerGroup(args[0])
	},
}

func deleteBgpPeerGroup(id string) error {
	log.Info().Msgf("Deleting BGP peer group with ID [%s].", id)

	response, err := networks.Client.BgpPeerGroupDeleteById(id)

	if err != nil {
		return err
	} else {
		return printer.PrintBgpPeerGroupResponse(response)
	}
}
