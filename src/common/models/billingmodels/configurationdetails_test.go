package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestConfigurationDetailsFromSdk_NotNil(t *testing.T) {
	sdk := *GenerateConfigurationDetails()
	cli := ConfigurationDetailsFromSdk(sdk)

	assertEqualConfigurationDetails(t, sdk, cli)
}

func assertEqualConfigurationDetails(t *testing.T, sdk billingapi.ConfigurationDetails, cli ConfigurationDetails) {
	assert.Equal(t, sdk.ThresholdConfiguration.ThresholdAmount, cli.ThresholdConfiguration.ThresholdAmount)
}
