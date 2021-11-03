package models

import (
	"fmt"
	"time"

	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type Quota struct {
	ID                           string                         `json:"id" yaml:"id"`
	Name                         string                         `json:"name" yaml:"name"`
	Description                  string                         `json:"description" yaml:"description"`
	Status                       string                         `json:"status" yaml:"status"`
	Limit                        int32                          `json:"limit" yaml:"limit"`
	Unit                         string                         `json:"unit" yaml:"unit"`
	Used                         int32                          `json:"used" yaml:"used"`
	QuotaEditLimitRequestDetails []QuotaEditLimitRequestDetails `json:"quotaEditLimitRequestDetails" yaml:"quotaEditLimitRequestDetails"`
}

type QuotaEditLimitRequestDetails struct {
	Limit       int32     `json:"limit" yaml:"limit"`
	Reason      string    `json:"reason" yaml:"reason"`
	RequestedOn time.Time `yajsonml:"requestedOn" yaml:"requestedOn"`
}

func QuotaSdkToDto(quota bmcapisdk.Quota) Quota {
	return Quota{
		ID:                           quota.Id,
		Name:                         quota.Name,
		Description:                  quota.Description,
		Status:                       quota.Status,
		Limit:                        quota.Limit,
		Unit:                         quota.Unit,
		Used:                         quota.Used,
		QuotaEditLimitRequestDetails: quotaEditLimitRequestDetailsListSdkToDto(quota.QuotaEditLimitRequestDetails),
	}
}

func quotaEditLimitRequestDetailsListSdkToDto(requestDetailsList []bmcapisdk.QuotaEditLimitRequestDetails) []QuotaEditLimitRequestDetails {
	if len(requestDetailsList) < 1 {
		return nil
	}

	var bmcRequestDetails []QuotaEditLimitRequestDetails

	for _, request := range requestDetailsList {
		bmcRequestDetails = append(bmcRequestDetails, quotaEditLimitRequestDetailsSdkToDto(request))
	}

	return bmcRequestDetails
}

func quotaEditLimitRequestDetailsSdkToDto(requestDetails bmcapisdk.QuotaEditLimitRequestDetails) QuotaEditLimitRequestDetails {
	return QuotaEditLimitRequestDetails{
		Limit:       requestDetails.Limit,
		Reason:      requestDetails.Reason,
		RequestedOn: requestDetails.RequestedOn,
	}
}

func QuotaEditLimitRequestDetailsToTableString(requestDetails []bmcapisdk.QuotaEditLimitRequestDetails) []string {
	var detailsAsStrings []string
	if len(requestDetails) < 1 {
		detailsAsStrings = []string{"N/A"}
	} else {
		dtoDetails := quotaEditLimitRequestDetailsListSdkToDto(requestDetails)
		for _, details := range dtoDetails {
			detailsAsStrings = append(detailsAsStrings, quotaEditLimitRequestDetailsToString(details))
		}
	}

	return detailsAsStrings
}

func quotaEditLimitRequestDetailsToString(requestDetails QuotaEditLimitRequestDetails) string {
	return "Limit: " + fmt.Sprint(requestDetails.Limit) +
		"\nReason: " + requestDetails.Reason +
		"\nRequestedOn: " + requestDetails.RequestedOn.String()
}
