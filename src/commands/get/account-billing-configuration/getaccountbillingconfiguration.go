package accountbillingconfiguration

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

func init() {
	utils.SetupOutputFlag(GetAccountBillingConfigurationCmd)
}

var GetAccountBillingConfigurationCmd = &cobra.Command{
	Use:          "account-billing-configuration",
	Short:        "Retrieve your account billing configuration",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long:         `Retrieve your account billing configuration.`,
	Example: `
# Retrieve your account billing configuration
pnapctl get account-billing-configuration [--output=<OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return getAccountBillingConfiguration()
	},
}

func getAccountBillingConfiguration() error {
	log.Info().Msg("Retrieving account billing configuration...")

	configurationDetails, err := billing.Client.AccountBillingConfigurationGet()

	if err != nil {
		return err
	} else {
		log.Info().Msg("Retrieving configuration details...")
		return printer.PrintConfigurationDetailsResponse(configurationDetails)
	}
}
