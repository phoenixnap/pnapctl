package generators

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateQuotaSdk() bmcapisdk.Quota {
	return bmcapisdk.Quota{
		Id:                           testutil.RandSeq(10),
		Name:                         testutil.RandSeq(10),
		Description:                  testutil.RandSeq(10),
		Status:                       testutil.RandSeq(10),
		Limit:                        int32(rand.Int()),
		Unit:                         testutil.RandSeq(10),
		Used:                         int32(rand.Int()),
		QuotaEditLimitRequestDetails: []bmcapisdk.QuotaEditLimitRequestDetails{},
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
		Reason: testutil.RandSeq(10),
	}
}

func GenerateQuotaEditLimitRequestDetailsSdk() bmcapisdk.QuotaEditLimitRequestDetails {
	return bmcapisdk.QuotaEditLimitRequestDetails{
		Limit:       int32(rand.Int()),
		Reason:      testutil.RandSeq(10),
		RequestedOn: time.Now(),
	}
}
