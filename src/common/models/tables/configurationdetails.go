package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"phoenixnap.com/pnapctl/common/models"
)

type ConfigurationDetailsTable struct {
	ThresholdConfiguration string `header:"Threshold Configuration"`
}

func ConfigurationDetailsTableFromSdk(sdk billingapi.ConfigurationDetails) ConfigurationDetailsTable {
	return ConfigurationDetailsTable{
		ThresholdConfiguration: models.ThresholdConfigurationToTableString(sdk.ThresholdConfiguration),
	}
}
