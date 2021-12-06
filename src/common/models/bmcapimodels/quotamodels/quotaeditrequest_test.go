package quotamodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestQuotaEditRequestToSdk(test_framework *testing.T) {
	quotaEditLimitRequest := GenerateQuotaEditLimitRequestCli()
	sdkModel := quotaEditLimitRequest.toSdk()

	assert.Equal(test_framework, quotaEditLimitRequest.Limit, sdkModel.Limit)
	assert.Equal(test_framework, quotaEditLimitRequest.Reason, sdkModel.Reason)
}
