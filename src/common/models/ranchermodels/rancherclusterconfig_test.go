package ranchermodels

import (
	"fmt"
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
)

func TestRancherClusterConfigToSdk(test_framework *testing.T) {
	rancherClusterConfig := GenerateRancherClusterConfigCli()
	sdkRancherClusterConfig := *rancherClusterConfig.ToSdk()

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherClusterConfigFromSdk(test_framework *testing.T) {
	sdkRancherClusterConfig := GenerateRancherClusterConfigSdk()
	rancherClusterConfig := *RancherClusterConfigFromSdk(&sdkRancherClusterConfig)

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherClusterConfigToTableString_nilConfig(test_framework *testing.T) {
	result := RancherClusterConfigToTableString(nil)
	assert.Equal(test_framework, "", result)
}

func TestNodePoolsToTableStrings_withClusterConfig(test_framework *testing.T) {
	sdkModel := GenerateRancherClusterConfigSdk()

	result := RancherClusterConfigToTableString(&sdkModel)

	assert.Equal(test_framework, result, generateClusterConfigResultString(&sdkModel))
}

func generateClusterConfigResultString(config *ranchersdk.RancherClusterConfig) string {
	return fmt.Sprintf("Token: %s, Domain: %s", *config.Token, *config.ClusterDomain)
}

func assertEqualRancherClusterConfig(test_framework *testing.T, r1 RancherClusterConfig, r2 ranchersdk.RancherClusterConfig) {
	assert.Equal(test_framework, r1.Token, r2.Token)
	assert.Equal(test_framework, r1.TlsSan, r2.TlsSan)
	assert.Equal(test_framework, r1.EtcdSnapshotScheduleCron, r2.EtcdSnapshotScheduleCron)
	assert.Equal(test_framework, r1.EtcdSnapshotRetention, r2.EtcdSnapshotRetention)
	assert.Equal(test_framework, r1.NodeTaint, r2.NodeTaint)
	assert.Equal(test_framework, r1.ClusterDomain, r2.ClusterDomain)

	if !assertNilEquality(test_framework, "Certificates", r1.Certificates, r2.Certificates) {
		assertEqualRancherClusterCertificates(test_framework, *r1.Certificates, *r2.Certificates)
	}
}
