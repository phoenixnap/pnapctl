package ipblock

import (
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Filename string

var Full bool

const commandName = "patch ip-block"

func init() {
	utils.SetupOutputFlag(PatchIpBlockCmd)
	utils.SetupFullFlag(PatchIpBlockCmd, &Full, "ip-block")
	utils.SetupFilenameFlag(PatchIpBlockCmd, &Filename, utils.UPDATING)
}

var PatchIpBlockCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Updates a specific ip-block.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch an existing ip-block.

Requires a file (yaml or json) containing the information needed to update the ip-block.`,
	Example: `# Update an existing ip-block with request body as described in ipblockpatch.yaml
	pnapctl patch ip-block <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]
	
	# ipblockpatch.yaml
	description: ip block description`,

	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return patchIpBlock(args[0])
	},
}

func patchIpBlock(id string) error {
	ipBlockPatch, err := models.CreateRequestFromFile[ipapi.IpBlockPatch](Filename)

	if err != nil {
		return err
	}

	response, httpResponse, err := ip.Client.IpBlocksIpBlockIdPatch(id, *ipBlockPatch)

	var generatedError = utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintIpBlockResponse(response, Full)
	}
}
