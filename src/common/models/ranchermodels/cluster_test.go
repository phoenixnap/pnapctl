package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestClusterToSdk(test_framework *testing.T) {
	cluster := GenerateClusterCli()
	sdkCluster := cluster.ToSdk()

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFullToSdk(test_framework *testing.T) {
	cluster := GenerateClusterCli()

	config := GenerateRancherClusterConfigCli()
	metadata := GenerateRancherServerMetadataCli()
	nodepools := []NodePool{GenerateNodePoolCli()}

	cluster.Configuration = &config
	cluster.Metadata = &metadata
	cluster.NodePools = nodepools

	sdkCluster := cluster.ToSdk()

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFromSdk(test_framework *testing.T) {
	sdkCluster := GenerateClusterSdk()
	cluster := ClusterFromSdk(sdkCluster)

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFullFromSdk(test_framework *testing.T) {
	sdkCluster := GenerateClusterSdk()

	config := GenerateRancherClusterConfigSdk()
	metadata := GenerateRancherServerMetadataSdk()
	nodepools := []ranchersdk.NodePool{GenerateNodePoolSdk()}

	sdkCluster.Configuration = &config
	sdkCluster.Metadata = &metadata
	sdkCluster.NodePools = nodepools

	cluster := ClusterFromSdk(sdkCluster)

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func assertEqualCluster(test_framework *testing.T, cliCluster Cluster, sdkCluster ranchersdk.Cluster) {
	assert.Equal(test_framework, cliCluster.Id, sdkCluster.Id)
	assert.Equal(test_framework, cliCluster.Name, sdkCluster.Name)
	assert.Equal(test_framework, cliCluster.Description, sdkCluster.Description)
	assert.Equal(test_framework, cliCluster.Location, sdkCluster.Location)
	assert.Equal(test_framework, cliCluster.InitialClusterVersion, sdkCluster.InitialClusterVersion)
	assert.Equal(test_framework, cliCluster.StatusDescription, sdkCluster.StatusDescription)

	if testutil.AssertNilEquality(test_framework, "Node Pools", cliCluster.NodePools, sdkCluster.NodePools) {
		assert.Equal(test_framework, len(cliCluster.NodePools), len(sdkCluster.NodePools))
		for i := range cliCluster.NodePools {
			assertEqualNodePool(test_framework, (cliCluster.NodePools)[i], (sdkCluster.NodePools)[i])
		}
	}

	if testutil.AssertNilEquality(test_framework, "Configuration", cliCluster.Configuration, sdkCluster.Configuration) {
		assertEqualRancherClusterConfig(test_framework, *cliCluster.Configuration, *sdkCluster.Configuration)
	}

	if testutil.AssertNilEquality(test_framework, "Metadata", cliCluster.Metadata, sdkCluster.Metadata) {
		assertEqualRancherServerMetadata(test_framework, *cliCluster.Metadata, *sdkCluster.Metadata)
	}
}
