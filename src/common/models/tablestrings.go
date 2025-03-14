package models

import (
	"fmt"
	"strings"

	"github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
)

// utils
func toTableString[T, O any](ifPresent func(T) O) func(*T) O {
	return func(item *T) O {
		if item == nil {
			var o O
			return o
		}
		return ifPresent(*item)
	}
}

func processNil[T any](item *T) T {
	if item == nil {
		var t T
		return t
	}
	return *item
}

// auditapi
var UserInfoToTableString = toTableString(func(sdk auditapi.UserInfo) string {
	return fmt.Sprintf("Account:(%s)\nClientId:(%s)\nUsername:%s", sdk.AccountId, processNil(sdk.ClientId), sdk.Username)
})

// billingapi
var LocationAvailabilityDetailsToTableString = toTableString(func(sdk billingapi.LocationAvailabilityDetail) string {
	return fmt.Sprintf("%s - (Req: %f/Available: %f)", sdk.Location, sdk.MinQuantityRequested, sdk.AvailableQuantity)
})

var ThresholdConfigurationToTableString = toTableString(func(sdk billingapi.ThresholdConfigurationDetails) string {
	return fmt.Sprintf("%f", sdk.GetThresholdAmount())
})

var PricingPlanToTableString = toTableString(func(sdk billingapi.PricingPlan) string {
	return fmt.Sprintf("Sku: %s\nPrice: %f\nPrice Unit: %s", sdk.Sku, sdk.Price, sdk.PriceUnit)
})

// bmcapi
var QuotaEditLimitRequestDetailsToTableString = toTableString(func(sdk bmcapi.QuotaEditLimitRequestDetails) string {
	return fmt.Sprintf("Limit: %d\nReason: %s\nRequestedOn: %s", sdk.Limit, sdk.Reason, sdk.RequestedOn)
})

var TagsToTableString = toTableString(func(sdk bmcapi.TagAssignment) string {
	var tagValue string
	if sdk.Value != nil {
		tagValue = ": " + *sdk.Value
	}
	return fmt.Sprintf("(%s) %s%s", sdk.Id, sdk.Name, tagValue)
})

var OsConfigurationToTableString = toTableString(func(sdk bmcapi.OsConfiguration) (out string) {
	if sdk.RootPassword != nil {
		out = "Password: " + *sdk.RootPassword
	}
	return
})

var NetworkConfigurationToTableString = toTableString(func(sdk bmcapi.NetworkConfiguration) string {
	if sdk.PrivateNetworkConfiguration == nil {
		return "Public"
	} else {
		return "Private"
	}
})

var StorageConfigurationToTableString = toTableString(func(sdk bmcapi.StorageConfiguration) string {
	return fmt.Sprintf("Raid: %s\nSize: %d",
		*sdk.RootPartition.Raid, sdk.RootPartition.Size)
})

// ipapi
var TagAssignmentToTableString = toTableString(func(sdk ipapi.TagAssignment) string {
	return fmt.Sprintf("ID: %s\nName: %s\nValue: %s\nIsBillingTag: %t\nCreated By: %s",
		sdk.Id, sdk.Name, processNil(sdk.Value), sdk.IsBillingTag, processNil(sdk.CreatedBy))
})

// networkapi
var NetworkMembershipToTableString = toTableString(func(sdk networkapi.NetworkMembership) string {
	return fmt.Sprintf("%s(%s)\n%v", sdk.ResourceType, sdk.ResourceId, sdk.Ips)
})

var PublicNetworkIpBlockToTableString = toTableString(func(sdk networkapi.PublicNetworkIpBlock) string {
	return fmt.Sprintf("ID: %s", sdk.Id)
})

var PrivateNetworkServerToTableString = toTableString(func(sdk networkapi.PrivateNetworkServer) string {
	return fmt.Sprintf("ID: %s\nIps: %v\n", sdk.Id, sdk.Ips)
})

var BgpIpv4PrefixToTableString = toTableString(func(sdk networkapi.BgpIPv4Prefix) string {
	return fmt.Sprintf("IPv4 Allocation Id: %s\nCidr: %s\nStatus: %s\nIn Use: %v", sdk.Ipv4AllocationId, sdk.Cidr, sdk.Status, sdk.InUse)
})

var AsnDetailsToTableString = toTableString(func(sdk networkapi.AsnDetails) string {
	return fmt.Sprintf("Asn: %v\nVerification Status: %s\nVerification Reason: %s", sdk.Asn, sdk.VerificationStatus, processNil(sdk.VerificationReason))
})

// ranchersolutionapi
var NodePoolToTableString = toTableString(func(sdk ranchersolutionapi.NodePool) string {
	return fmt.Sprintf("%s - %d nodes", processNil(sdk.Name), processNil(sdk.NodeCount))
})

var ClusterConfigurationToTableString = toTableString(func(sdk ranchersolutionapi.RancherClusterConfig) string {
	return fmt.Sprintf("Token: %s, Domain: %s", processNil(sdk.Token), processNil(sdk.ClusterDomain))
})

var ClusterMetadataToTableString = toTableString(func(sdk ranchersolutionapi.RancherServerMetadata) string {
	var username, password, url string
	if sdk.Username != nil {
		username = "User: " + *sdk.Username + "\n"
	}
	if sdk.Password != nil {
		password = "Pass: " + *sdk.Password + "\n"
	}
	if sdk.Url != nil {
		url = "Url: " + *sdk.Url + "\n"
	}
	return fmt.Sprintf("%s%s%s", username, password, url)
})

// tagapi
var ResourceAssignmentToTableString = toTableString(func(sdk tagapi.ResourceAssignment) string {
	var value string
	if sdk.Value == nil {
		value = "N/A"
	} else {
		value = *sdk.Value
	}
	return fmt.Sprintf("%s: %s", sdk.ResourceName, value)
})

// networkstorageapi
var VolumeToTableString = toTableString(func(sdk networkstorageapi.Volume) string {
	return fmt.Sprintf("(%s) %s [%dGb]", sdk.GetId(), sdk.GetName(), sdk.GetCapacityInGb())
})

var PermissionsToTableString = toTableString(func(sdk networkstorageapi.Permissions) string {
	var permissions string
	if sdk.Nfs != nil {
		permissions = permissions + fmt.Sprintf("NFS: %s", NfsPermissionsToTableString(sdk.Nfs))
	}
	return permissions
})

var NfsPermissionsToTableString = toTableString(func(sdk networkstorageapi.NfsPermissions) string {
	return strings.Join([]string{
		fmt.Sprintf("ReadWrite: %v", sdk.ReadWrite),
		fmt.Sprintf("ReadOnly: %v", sdk.ReadOnly),
		fmt.Sprintf("RootSquash: %v", sdk.RootSquash),
		fmt.Sprintf("NoSquash: %v", sdk.NoSquash),
		fmt.Sprintf("AllSquash: %v", sdk.AllSquash),
	}, "\n")
})

var StorageNetworkTagAssignmentToTableString = toTableString(func(sdk networkstorageapi.TagAssignment) string {
	return fmt.Sprintf("%v: %v", sdk.Name, sdk.Value)
})

// locationapi
var ProductCategoryToTableString = toTableString(func(sdk locationapi.ProductCategory) string {
	return strings.Join([]string{
		fmt.Sprintf("Product: %v", sdk.ProductCategory),
		fmt.Sprintf("Description: %v", sdk.ProductCategoryDescription),
	}, "\n")
})
