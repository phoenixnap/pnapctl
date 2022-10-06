package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestToQuotaTable(test_framework *testing.T) {
	quota := generators.GenerateQuotaSdk()
	table := ToQuotaTable(quota)

	assertQuotasEqual(test_framework, quota, table)
}

func assertQuotasEqual(test_framework *testing.T, quota bmcapisdk.Quota, table Quota) {
	assert.Equal(test_framework, quota.Id, table.Id)
	assert.Equal(test_framework, quota.Name, table.Name)
	assert.Equal(test_framework, quota.Description, table.Description)
	assert.Equal(test_framework, quota.Status, table.Status)
	assert.Equal(test_framework, quota.Limit, table.Limit)
	assert.Equal(test_framework, quota.Unit, table.Unit)
	assert.Equal(test_framework, quota.Used, table.Used)
	assert.Equal(test_framework, iterutils.MapRef(quota.QuotaEditLimitRequestDetails, models.QuotaEditLimitRequestDetailsToTableString), table.QuotaEditLimitRequestDetails)
}
