package tables

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type ConfigurationDetailsTable struct {
	ThresholdConfiguration string //wrapping a float32
}

func AccountBillingConfigurationTableFromSdk(sdk *billingapi.ConfigurationDetails) *ConfigurationDetailsTable {
	if sdk == nil {
		return nil
	}

	return &ConfigurationDetailsTable{}
}
