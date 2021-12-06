package quotamodels

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func TestQuotaSdkToDto(test_framework *testing.T) {
	sdkModel := GenerateQuotaSdk()
	quota := QuotaFromSdk(sdkModel)

	assert.Equal(test_framework, sdkModel.Id, quota.ID)
	assert.Equal(test_framework, sdkModel.Name, quota.Name)
	assert.Equal(test_framework, sdkModel.Description, quota.Description)
	assert.Equal(test_framework, sdkModel.Status, quota.Status)
	assert.Equal(test_framework, sdkModel.Limit, quota.Limit)
	assert.Equal(test_framework, sdkModel.Unit, quota.Unit)
	assert.Equal(test_framework, sdkModel.Used, quota.Used)
	assertEqualQuotaEditLimitRequestDetails(test_framework, sdkModel.QuotaEditLimitRequestDetails, quota.QuotaEditLimitRequestDetails)
}

func TestQuotaEditLimitRequestDetailsToTableString_noRequests(test_framework *testing.T) {
	list := []bmcapisdk.QuotaEditLimitRequestDetails{}

	result := QuotaEditLimitRequestDetailsToTableString(list)

	assert.Equal(test_framework, len(result), 1)
	assert.Equal(test_framework, result[0], "N/A")
}

func TestQuotaEditLimitRequestDetailsToTableString_witRequests(test_framework *testing.T) {
	sdkModel_1 := GenerateQuotaEditLimitRequestDetailsSdk()
	sdkModel_2 := GenerateQuotaEditLimitRequestDetailsSdk()

	list := []bmcapisdk.QuotaEditLimitRequestDetails{
		sdkModel_1, sdkModel_2,
	}

	result := QuotaEditLimitRequestDetailsToTableString(list)

	assert.Equal(test_framework, len(result), 2)
	assert.Equal(test_framework, result[0], generateResultString(sdkModel_1))
	assert.Equal(test_framework, result[1], generateResultString(sdkModel_2))
}

func assertEqualQuotaEditLimitRequestDetails(test_framework *testing.T, sdkList []bmcapisdk.QuotaEditLimitRequestDetails, dtoList []QuotaEditLimitRequestDetails) {

	for i, sdkModel := range sdkList {
		assert.Equal(test_framework, sdkModel.Limit, dtoList[i].Limit)
		assert.Equal(test_framework, sdkModel.Reason, dtoList[i].Reason)
		assert.Equal(test_framework, sdkModel.RequestedOn, dtoList[i].RequestedOn)
	}
}

func generateResultString(requestDetails bmcapisdk.QuotaEditLimitRequestDetails) string {
	return fmt.Sprintf("Limit: %d\nReason: %s\nRequestedOn: %s", requestDetails.Limit, requestDetails.Reason, requestDetails.RequestedOn)
}
