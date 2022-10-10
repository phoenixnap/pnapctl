package ranchermodels

import (
	"fmt"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
)

type RancherServerMetadata struct {
	Url      *string `json:"url" yaml:"url"`
	Username *string `json:"username" yaml:"username"`
	Password *string `json:"password" yaml:"password"`
}

func (r RancherServerMetadata) ToSdk() *ranchersdk.ClusterMetadata {
	return &ranchersdk.ClusterMetadata{
		Url:      r.Url,
		Username: r.Username,
		Password: r.Password,
	}
}

func RancherServerMetadataFromSdk(metadata *ranchersdk.ClusterMetadata) *RancherServerMetadata {
	if metadata == nil {
		return nil
	}

	return &RancherServerMetadata{
		Url:      metadata.Url,
		Username: metadata.Username,
		Password: metadata.Password,
	}
}

func RancherServerMetadataToTableString(metadata *ranchersdk.ClusterMetadata) string {
	if metadata == nil {
		return ""
	}

	username := ""
	password := ""
	url := ""

	if metadata.Username != nil {
		username = "User: " + *metadata.Username + "\n"
	}
	if metadata.Password != nil {
		password = "Pass: " + *metadata.Password + "\n"
	}
	if metadata.Url != nil {
		url = "Url: " + *metadata.Url + "\n"
	}

	return fmt.Sprintf("%s%s%s", username, password, url)
}
