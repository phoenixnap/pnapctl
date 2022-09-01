package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ConfigurationDetails struct {
	ThresholdConfiguration *ThresholdConfigurationDetails `json:"thresholdConfiguration,omitempty" yaml:"thresholdConfiguration,omitempty"`
}

type ThresholdConfigurationDetails struct {
	ThresholdAmount float32 `json:"thresholdAmount" yaml:"thresholdAmount"`
}

func ConfigurationDetailsFromSdk(sdk billingapi.ConfigurationDetails) ConfigurationDetails {
	return ConfigurationDetails{
		ThresholdConfiguration: thresholdConfigurationDetailsFromSdk(sdk.ThresholdConfiguration),
	}
}

func thresholdConfigurationDetailsFromSdk(sdk *billingapi.ThresholdConfigurationDetails) *ThresholdConfigurationDetails {
	if sdk == nil {
		return nil
	}

	return &ThresholdConfigurationDetails{
		ThresholdAmount: sdk.ThresholdAmount,
	}
}
