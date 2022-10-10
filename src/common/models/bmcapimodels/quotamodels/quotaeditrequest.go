package quotamodels

import bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"

type QuotaEditRequest struct {
	Limit  int32  `json:"limit" yaml:"limit"`
	Reason string `json:"reason" yaml:"reason"`
}

func (quotaEditRequest QuotaEditRequest) toSdk() *bmcapisdk.QuotaEditLimitRequest {
	return &bmcapisdk.QuotaEditLimitRequest{
		Limit:  quotaEditRequest.Limit,
		Reason: quotaEditRequest.Reason,
	}
}
