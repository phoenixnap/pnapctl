package ranchermodels

import (
	"fmt"

	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
)

type RancherClusterConfig struct {
	Token                    *string
	TlsSan                   *string
	EtcdSnapshotScheduleCron *string
	EtcdSnapshotRetention    *int32
	NodeTaint                *string
	ClusterDomain            *string
	Certificates             *RancherClusterCertificates
}

func (r RancherClusterConfig) ToSdk() *ranchersdk.RancherClusterConfig {
	return &ranchersdk.RancherClusterConfig{
		Token:                    r.Token,
		TlsSan:                   r.TlsSan,
		EtcdSnapshotScheduleCron: r.EtcdSnapshotScheduleCron,
		EtcdSnapshotRetention:    r.EtcdSnapshotRetention,
		NodeTaint:                r.NodeTaint,
		ClusterDomain:            r.ClusterDomain,
		Certificates:             nil,
	}
}

func RancherClusterConfigFromSdk(config *ranchersdk.RancherClusterConfig) *RancherClusterConfig {
	if config == nil {
		return nil
	}

	return &RancherClusterConfig{
		Token:                    config.Token,
		TlsSan:                   config.TlsSan,
		EtcdSnapshotScheduleCron: config.EtcdSnapshotScheduleCron,
		EtcdSnapshotRetention:    config.EtcdSnapshotRetention,
		NodeTaint:                config.NodeTaint,
		ClusterDomain:            config.ClusterDomain,
		Certificates:             nil,
	}
}

func RancherClusterConfigToTableString(config *ranchersdk.RancherClusterConfig) string {
	if config == nil {
		return ""
	}

	return fmt.Sprintf("Token: %s, Domain: %s", *config.Token, *config.ClusterDomain)
}
