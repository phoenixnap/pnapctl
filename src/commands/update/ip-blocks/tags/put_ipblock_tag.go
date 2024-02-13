package ip_blocks

import (
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Filename string

var Full bool

func init() {
	utils.SetupOutputFlag(PutIpBlockTagCmd)
	utils.SetupFullFlag(PutIpBlockTagCmd, &Full, "ip-block")
	utils.SetupFilenameFlag(PutIpBlockTagCmd, &Filename, utils.UPDATING)
}

var PutIpBlockTagCmd = &cobra.Command{
	// TODO tag? not tags?
	Use:          "tag IP_BLOCK_ID",
	Short:        "Updates an ip block's tags.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update an existing ip-block's tag.

Requires a file (yaml or json) containing the information needed to update the ip-block's tags.
	`,
	Example: `# Update a tag on an existing ip-block with request body as described in ipblockputtag.yaml
pnapctl update ip-block tag <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# ipblockputtag.yaml
---
- name: ip block tag name
  value: ip block tag value`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updateTagsOnIpBlock(args[0])
	},
}

func updateTagsOnIpBlock(id string) error {
	log.Info().Msgf("Updating tags for Ip Block with ID [%s].", id)

	ipBlockPutTag, err := models.CreateRequestFromFile[[]ipapi.TagAssignmentRequest](Filename)

	if err != nil {
		return err
	}

	response, err := ip.Client.IpBlocksIpBlockIdTagsPut(id, *ipBlockPutTag)

	if err != nil {
		return err
	} else {
		return printer.PrintIpBlockResponse(response, Full)
	}
}
