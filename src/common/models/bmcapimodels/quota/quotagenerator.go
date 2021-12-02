package quota

import (
	"math/rand"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	generators "phoenixnap.com/pnapctl/tests/generators"
)

func GenerateQuota() bmcapisdk.Quota {
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

func GenerateQuotas(n int) []bmcapisdk.Quota {
	var quotaList []bmcapisdk.Quota
	for i := 0; i < n; i++ {
		quotaList = append(quotaList, GenerateQuota())
	}
	return quotaList
}

func GenerateQuotaEditLimitRequest() bmcapisdk.QuotaEditLimitRequest {
	return bmcapisdk.QuotaEditLimitRequest{
		Limit:  int32(rand.Int()),
		Reason: generators.RandSeq(10),
	}
}
