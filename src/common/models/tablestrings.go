package models

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/phoenixnap/go-sdk-bmc/ipapi"
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/phoenixnap/go-sdk-bmc/tagapi"
)

// auditapi
func UserInfoToTableString(sdk *auditapi.UserInfo) string {
	return ""
}

// billingapi
func LocationAvailabilityDetailsToTableString(sdk billingapi.LocationAvailabilityDetail) string {
	return ""
}

func ThresholdConfigurationToTableString(sdk *billingapi.ThresholdConfigurationDetails) string {
	return fmt.Sprintf("%f", sdk.GetThresholdAmount())
}

func PricingPlanToTableString(sdk billingapi.PricingPlan) string {
	return ""
}

// bmcapi
func QuotaEditLimitRequestDetailsToTableString(sdk []bmcapi.QuotaEditLimitRequestDetails) []string {
	return []string{}
}

func TagsToTableStrings(sdk []bmcapi.TagAssignment) []string {
	return []string{}
}

func OsConfigurationToTableString(sdk *bmcapi.OsConfiguration) string {
	return ""
}

func NetworkConfigurationToTableString(sdk bmcapi.NetworkConfiguration) string {
	return ""
}

// ipapi
func TagAssignmentToTableString(sdk *ipapi.TagAssignment) string {
	return ""
}

// networkapi
func NetworkMembershipToTableString(sdk networkapi.NetworkMembership) string {
	return ""
}

func PublicNetworkIpBlockToTableString(sdk networkapi.PublicNetworkIpBlock) string {
	return ""
}

func PrivateNetworkServerToTableString(sdk *networkapi.PrivateNetworkServer) string {
	return ""
}

// ranchersolutionapi
func NodePoolsToTableStrings(sdk []ranchersolutionapi.NodePool) []string {
	return nil
}

func ClusterConfigurationToTableString(sdk *ranchersolutionapi.ClusterConfiguration) string {
	return ""
}

func ClusterMetadataToTableString(sdk *ranchersolutionapi.ClusterMetadata) string {
	return ""
}

// tagapi
func ResourceAssignmentToTableString(sdk tagapi.ResourceAssignment) string {
	return ""
}
