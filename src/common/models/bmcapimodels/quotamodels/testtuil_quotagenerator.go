package quotamodels

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	generators "phoenixnap.com/pnapctl/testsupport/generators"
)

func GenerateQuotaSdk() bmcapisdk.Quota {
	return bmcapisdk.Quota{
		Id:                           generators.RandSeq(10),
		Name:                         generators.RandSeq(10),
		Description:                  generators.RandSeq(10),
		Status:                       generators.RandSeq(10),
		Limit:                        int32(rand.Int()),
		Unit:                         generators.RandSeq(10),
		Used:                         int32(rand.Int()),
		QuotaEditLimitRequestDetails: []bmcapisdk.QuotaEditLimitRequestDetails{},
	}
}

func GenerateQuotaCli() Quota {
	return Quota{
		ID:                           generators.RandSeq(10),
		Name:                         generators.RandSeq(10),
		Description:                  generators.RandSeq(10),
		Status:                       generators.RandSeq(10),
		Limit:                        int32(rand.Int()),
		Unit:                         generators.RandSeq(10),
		Used:                         int32(rand.Int()),
		QuotaEditLimitRequestDetails: []QuotaEditLimitRequestDetails{},
	}
}

func GenerateQuotaSdkList(n int) []bmcapisdk.Quota {
	var quotaList []bmcapisdk.Quota
	for i := 0; i < n; i++ {
		quotaList = append(quotaList, GenerateQuotaSdk())
	}
	return quotaList
}

func GenerateQuotaEditLimitRequestSdk() bmcapisdk.QuotaEditLimitRequest {
	return bmcapisdk.QuotaEditLimitRequest{
		Limit:  int32(rand.Int()),
		Reason: generators.RandSeq(10),
	}
}

func GenerateQuotaEditLimitRequestCli() QuotaEditRequest {
	return QuotaEditRequest{
		Limit:  int32(rand.Int()),
		Reason: generators.RandSeq(10),
	}
}

func GenerateQuotaEditLimitRequestDetailsSdk() bmcapisdk.QuotaEditLimitRequestDetails {
	return bmcapisdk.QuotaEditLimitRequestDetails{
		Limit:       int32(rand.Int()),
		Reason:      generators.RandSeq(10),
		RequestedOn: time.Now(),
	}
}
