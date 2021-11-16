package ranchermodels

import (
	"fmt"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

type RancherServerMetadata struct {
	Url      *string `json:"url" yaml:"url"`
	Username *string `json:"username" yaml:"username"`
	Password *string `json:"password" yaml:"password"`
}

func (r RancherServerMetadata) ToSdk() *ranchersdk.RancherServerMetadata {
	return &ranchersdk.RancherServerMetadata{
		Url:      r.Url,
		Username: r.Username,
		Password: r.Password,
	}
}

func RancherServerMetadataFromSdk(metadata *ranchersdk.RancherServerMetadata) *RancherServerMetadata {
	if metadata == nil {
		return nil
	}

	return &RancherServerMetadata{
		Url:      metadata.Url,
		Username: metadata.Username,
		Password: metadata.Password,
	}
}

func RancherServerMetadataToTableString(metadata *ranchersdk.RancherServerMetadata) string {
	if metadata == nil {
		return ""
	}

	return fmt.Sprintf("(User: %s, Pass: %s, Url: %s)", *metadata.Username, *metadata.Password, *metadata.Url)
}
