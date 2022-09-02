package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

type ConfigurationDetailsTable struct {
	ThresholdConfiguration string `header:"Threshold Configuration"`
}

func ConfigurationDetailsTableFromSdk(sdk billingapi.ConfigurationDetails) ConfigurationDetailsTable {
	return ConfigurationDetailsTable{
		ThresholdConfiguration: billingmodels.ThresholdConfigurationToTableString(sdk.ThresholdConfiguration),
	}
}
