package tables

import (
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestToConfigurationDetailsTable_NotNil(t *testing.T) {
	sdk := generators.Generate[billingapi.ConfigurationDetails]()
	tbl := ConfigurationDetailsTableFromSdk(sdk)

	assertConfigurationTablesEqual(t, sdk, tbl)
}

func assertConfigurationTablesEqual(t *testing.T, sdk billingapi.ConfigurationDetails, tbl ConfigurationDetailsTable) {
	assert.Equal(t, fmt.Sprintf("%f", sdk.ThresholdConfiguration.ThresholdAmount), tbl.ThresholdConfiguration)
}
