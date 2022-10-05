package networkstoragemodels

import (
	"fmt"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type StorageNetwork struct {
	Id          *string                   `json:"id,omitempty" yaml:"id,omitempty"`
	Name        *string                   `json:"name,omitempty" yaml:"name,omitempty"`
	Description *string                   `json:"description,omitempty" yaml:"description,omitempty"`
	Status      *networkstorageapi.Status `json:"status,omitempty" yaml:"status,omitempty"`
	Location    *string                   `json:"location,omitempty" yaml:"location,omitempty"`
	NetworkId   *string                   `json:"networkId,omitempty" yaml:"networkId,omitempty"`
	Ips         []string                  `json:"ips,omitempty" yaml:"ips,omitempty"`
	CreatedOn   *time.Time                `json:"createdOn,omitempty" yaml:"createdOn,omitempty"`
	Volumes     []Volume                  `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}

type Volume struct {
	Id           *string                   `json:"id,omitempty" yaml:"id,omitempty"`
	Name         *string                   `json:"name,omitempty" yaml:"name,omitempty"`
	Description  *string                   `json:"description,omitempty" yaml:"description,omitempty"`
	Path         *string                   `json:"path,omitempty" yaml:"path,omitempty"`
	PathSuffix   *string                   `json:"pathSuffix,omitempty" yaml:"pathSuffix,omitempty"`
	CapacityInGb *int32                    `json:"capacityInGb,omitempty" yaml:"capacityInGb,omitempty"`
	Protocol     *string                   `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Status       *networkstorageapi.Status `json:"status,omitempty" yaml:"status,omitempty"`
	CreatedOn    *time.Time                `json:"createdOn,omitempty" yaml:"createdOn,omitempty"`
	Permissions  *Permissions              `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}

type Permissions struct {
	Nfs *NfsPermissions `json:"nfs,omitempty"`
}

type NfsPermissions struct {
	ReadWrite  []string `json:"readWrite,omitempty"`
	ReadOnly   []string `json:"readOnly,omitempty"`
	RootSquash []string `json:"rootSquash,omitempty"`
	NoSquash   []string `json:"noSquash,omitempty"`
	AllSquash  []string `json:"allSquash,omitempty"`
}

// From SDK
func StorageNetworkFromSdk(sdk networkstorageapi.StorageNetwork) StorageNetwork {
	return StorageNetwork{
		Id:          sdk.Id,
		Name:        sdk.Name,
		Description: sdk.Description,
		Status:      sdk.Status,
		Location:    sdk.Location,
		NetworkId:   sdk.NetworkId,
		Ips:         sdk.Ips,
		CreatedOn:   sdk.CreatedOn,
		Volumes:     iterutils.Map(sdk.Volumes, VolumeFromSdk),
	}
}

func VolumeFromSdk(sdk networkstorageapi.Volume) Volume {
	return Volume{
		Id:           sdk.Id,
		Name:         sdk.Name,
		Description:  sdk.Description,
		Path:         sdk.Path,
		PathSuffix:   sdk.PathSuffix,
		CapacityInGb: sdk.CapacityInGb,
		Protocol:     sdk.Protocol,
		Status:       sdk.Status,
		CreatedOn:    sdk.CreatedOn,
		Permissions:  iterutils.OptionalMapper(PermissionsFromSdk)(sdk.Permissions),
	}
}

func PermissionsFromSdk(sdk networkstorageapi.Permissions) Permissions {
	return Permissions{
		Nfs: iterutils.OptionalMapper(NfsPermissionsFromSdk)(sdk.Nfs),
	}
}

func NfsPermissionsFromSdk(sdk networkstorageapi.NfsPermissions) NfsPermissions {
	return NfsPermissions{
		ReadWrite:  sdk.ReadWrite,
		ReadOnly:   sdk.ReadOnly,
		RootSquash: sdk.RootSquash,
		NoSquash:   sdk.NoSquash,
		AllSquash:  sdk.AllSquash,
	}
}

// To Table String
func VolumeToTableString(sdk networkstorageapi.Volume) string {
	return fmt.Sprintf("(%s) %s [%dGb]", sdk.GetId(), sdk.GetName(), sdk.GetCapacityInGb())
}

func PermissionsToTableString(sdk *networkstorageapi.Permissions) string {
	if sdk == nil {
		return ""
	}

	permissions := ""
	if sdk.Nfs != nil {
		permissions = fmt.Sprintf("%sNFS: %s", permissions, NfsPermissionsToTableString(sdk.Nfs))
	}
	return permissions
}

func NfsPermissionsToTableString(sdk *networkstorageapi.NfsPermissions) string {
	if sdk == nil {
		return ""
	}

	rw := "ReadWrite: " + fmt.Sprintf("%v", sdk.ReadWrite)
	ro := "ReadOnly: " + fmt.Sprintf("%v", sdk.ReadOnly)
	rs := "RootSquash: " + fmt.Sprintf("%v", sdk.RootSquash)
	ns := "NoSquash: " + fmt.Sprintf("%v", sdk.NoSquash)
	as := "AllSquash: " + fmt.Sprintf("%v", sdk.AllSquash)

	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v", rw, ro, rs, ns, as)
}
