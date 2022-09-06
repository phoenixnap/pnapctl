package accountbillingconfiguration

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get account-billing-configuration"

var GetAccountBillingConfigurationCmd = &cobra.Command{
	Use:          "account-billing-configuration",
	Short:        "Retrieve your account billing configuration",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long:         `Retrieve your account billing configuration.`,
	Example: `
# Retrieve your account billing configuration
pnapctl get account-billing-configuration [--output=<OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getAccountBillingConfiguration()
	},
}

func getAccountBillingConfiguration() error {
	configurationDetails, httpResponse, err := billing.Client.AccountBillingConfigurationGet()

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintConfigurationDetailsResponse(configurationDetails, commandName)
	}
}

func init() {
	utils.SetupOutputFlag(GetAccountBillingConfigurationCmd)
}
