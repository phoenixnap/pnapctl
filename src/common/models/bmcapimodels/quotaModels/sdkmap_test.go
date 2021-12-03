package quotaModels

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func TestQuotaEditRequestToSdk(test_framework *testing.T) {
	quotaEditLimitRequest := GenerateQuotaEditLimitRequestCli()
	sdkModel := quotaEditLimitRequest.toSdk()

	assert.Equal(test_framework, quotaEditLimitRequest.Limit, sdkModel.Limit)
	assert.Equal(test_framework, quotaEditLimitRequest.Reason, sdkModel.Reason)
}

func TestQuotaSdkToDto(test_framework *testing.T) {
	sdkModel := GenerateQuotaSdk()
	quotaEditLimitRequest := QuotaSdkToDto(sdkModel)

	assert.Equal(test_framework, sdkModel.Id, quotaEditLimitRequest.ID)
	assert.Equal(test_framework, sdkModel.Name, quotaEditLimitRequest.Name)
	assert.Equal(test_framework, sdkModel.Description, quotaEditLimitRequest.Description)
	assert.Equal(test_framework, sdkModel.Status, quotaEditLimitRequest.Status)
	assert.Equal(test_framework, sdkModel.Limit, quotaEditLimitRequest.Limit)
	assert.Equal(test_framework, sdkModel.Unit, quotaEditLimitRequest.Unit)
	assert.Equal(test_framework, sdkModel.Used, quotaEditLimitRequest.Used)
	assertEqualQuotaEditLimitRequestDetails(test_framework, sdkModel.QuotaEditLimitRequestDetails, quotaEditLimitRequest.QuotaEditLimitRequestDetails)
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
