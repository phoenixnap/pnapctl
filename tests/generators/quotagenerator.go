package generators

import (
	"math/rand"

	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

func GenerateQuota() bmcapisdk.Quota {
	return bmcapisdk.Quota{
		Id:                           randSeq(10),
		Name:                         randSeq(10),
		Description:                  randSeq(10),
		Status:                       randSeq(10),
		Limit:                        int32(rand.Int()),
		Unit:                         randSeq(10),
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
		Reason: randSeq(10),
	}
}
