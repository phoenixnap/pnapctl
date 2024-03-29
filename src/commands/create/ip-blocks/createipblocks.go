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

// Filename is the filename from which to retrieve the request body
var Filename string

var Full bool

func init() {
	utils.SetupOutputFlag(CreateIpBlockCmd)
	utils.SetupFullFlag(CreateIpBlockCmd, &Full, "ip-block")
	utils.SetupFilenameFlag(CreateIpBlockCmd, &Filename, utils.CREATION)
}

// CreateIpBlockCmd is the command for creating an ip block.
var CreateIpBlockCmd = &cobra.Command{
	Use:          "ip-block",
	Short:        "Create a new ip-block.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new ip-block.

Requires a file (yaml or json) containing the information needed to create the ip-block.`,
	Example: `# Create a new ip-block as described in ipblockcreate.yaml
pnapctl create ip-block --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# ipblockcreate.yaml
cidrBlockSize: /28
location: PHX`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createIpBlock()
	},
}

func createIpBlock() error {
	log.Info().Msg("Creating new Ip Block...")

	ipBlockCreate, err := models.CreateRequestFromFile[ipapi.IpBlockCreate](Filename)

	if err != nil {
		return err
	}

	// Create the ssh key
	response, err := ip.Client.IpBlockPost(*ipBlockCreate)
	if err != nil {
		return err
	} else {
		return printer.PrintIpBlockResponse(response, Full)
	}
}
