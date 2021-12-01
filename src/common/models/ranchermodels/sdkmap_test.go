package ranchermodels

import (
	"reflect"
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
)

/*
Tests to cover:
CLUSTER							:
NODE POOL						:
NODE 							: FROM	TO
RANCHER CLUSTER CERTIFICATES	: NIL-FROM
RANCHER CLUSTER CONFIG			:
RANCHER SERVER METADATA			:
SSH CONFIG						:
*/

// tests
func TestClusterToSdk(test_framework *testing.T) {
	cluster := GenerateCLICluster()
	sdkCluster := cluster.ToSdk()

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFullToSdk(test_framework *testing.T) {
	cluster := GenerateCLICluster()

	config := GenerateCLIRancherClusterConfig()
	metadata := GenerateCLIRancherServerMetadata()
	nodepools := []NodePool{GenerateCLINodePool()}

	cluster.Configuration = &config
	cluster.Metadata = &metadata
	cluster.NodePools = &nodepools

	sdkCluster := cluster.ToSdk()

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFromSdk(test_framework *testing.T) {
	sdkCluster := GenerateCluster()
	cluster := ClusterFromSdk(sdkCluster)

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestClusterFullFromSdk(test_framework *testing.T) {
	sdkCluster := GenerateCluster()

	config := GenerateRancherClusterConfig()
	metadata := GenerateRancherServerMetadata()
	nodepools := []ranchersdk.NodePool{GenerateNodePool()}

	sdkCluster.Configuration = &config
	sdkCluster.Metadata = &metadata
	sdkCluster.NodePools = &nodepools

	cluster := ClusterFromSdk(sdkCluster)

	assertEqualCluster(test_framework, cluster, sdkCluster)
}

func TestNodePoolToSdk(test_framework *testing.T) {
	nodePool := GenerateCLINodePool()
	sdkNodePool := *nodePool.ToSdk()

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestFullNodePoolToSdk(test_framework *testing.T) {
	nodePool := GenerateCLINodePool()
	nodePool.Nodes = &[]Node{GenerateCLINode()}
	sdkNodePool := *nodePool.ToSdk()

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestNodePoolFromSdk(test_framework *testing.T) {
	sdkNodePool := GenerateNodePool()
	nodePool := NodePoolFromSdk(sdkNodePool)

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestFullNodePoolFromSdk(test_framework *testing.T) {
	sdkNodePool := GenerateNodePool()
	sdkNodePool.Nodes = &[]ranchersdk.Node{GenerateNode()}
	nodePool := NodePoolFromSdk(sdkNodePool)

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestNodeToSdk(test_framework *testing.T) {
	node := GenerateCLINode()
	sdkNode := *node.ToSdk()

	assertEqualNode(test_framework, node, sdkNode)
}

func TestNodeFromSdk(test_framework *testing.T) {
	sdkNode := GenerateNode()
	node := NodeFromSdk(sdkNode)

	assertEqualNode(test_framework, node, sdkNode)
}

func TestRancherClusterCertificateToSdk(test_framework *testing.T) {
	rancherRancherClusterCertificates := GenerateCLIRancherClusterCertificates()
	sdkRancherClusterCertificates := *rancherRancherClusterCertificates.toSdk()

	assertEqualRancherClusterCertificates(test_framework, rancherRancherClusterCertificates, sdkRancherClusterCertificates)
}

func TestRancherClusterCertificateFromSdk(test_framework *testing.T) {
	sdkRancherClusterCertificates := GenerateRancherClusterCertificates()
	rancherRancherClusterCertificates := *RancherClusterCertificatesFromSdk(&sdkRancherClusterCertificates)

	assertEqualRancherClusterCertificates(test_framework, rancherRancherClusterCertificates, sdkRancherClusterCertificates)
}

func TestRancherClusterConfigToSdk(test_framework *testing.T) {
	rancherClusterConfig := GenerateCLIRancherClusterConfig()
	sdkRancherClusterConfig := *rancherClusterConfig.ToSdk()

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherClusterConfigFromSdk(test_framework *testing.T) {
	sdkRancherClusterConfig := GenerateRancherClusterConfig()
	rancherClusterConfig := *RancherClusterConfigFromSdk(&sdkRancherClusterConfig)

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherServerMetadataToSdk(test_framework *testing.T) {
	rancherServerMetadata := GenerateCLIRancherServerMetadata()
	sdkRancherServerMetadata := *rancherServerMetadata.ToSdk()

	assertEqualRancherServerMetadata(test_framework, rancherServerMetadata, sdkRancherServerMetadata)
}

func TestRancherServerMetadataFromSdk(test_framework *testing.T) {
	sdkRancherServerMetadata := GenerateRancherServerMetadata()
	rancherServerMetadata := *RancherServerMetadataFromSdk(&sdkRancherServerMetadata)

	assertEqualRancherServerMetadata(test_framework, rancherServerMetadata, sdkRancherServerMetadata)
}

func TestSshConfigToSdk(test_framework *testing.T) {
	sshConfig := GenerateCLISshConfig()
	sdkSshConfig := *sshConfig.ToSdk()

	assertEqualSshConfig(test_framework, sshConfig, sdkSshConfig)
}

func TestSshConfigFromSdk(test_framework *testing.T) {
	sdkSshConfig := GenerateSshConfig()
	sshConfig := *SshConfigFromSdk(&sdkSshConfig)

	assertEqualSshConfig(test_framework, sshConfig, sdkSshConfig)
}

// assertion functions
func assertEqualCluster(test_framework *testing.T, c1 Cluster, c2 ranchersdk.Cluster) {
	assert.Equal(test_framework, c1.Id, c2.Id)
	assert.Equal(test_framework, c1.Name, c2.Name)
	assert.Equal(test_framework, c1.Description, c2.Description)
	assert.Equal(test_framework, c1.Location, c2.Location)
	assert.Equal(test_framework, c1.InitialClusterVersion, c2.InitialClusterVersion)
	assert.Equal(test_framework, c1.StatusDescription, c2.StatusDescription)

	if !assertNilEquality(test_framework, "Node Pools", c1.NodePools, c2.NodePools) {
		assert.Equal(test_framework, len(*c1.NodePools), len(*c2.NodePools))
		for i := range *c1.NodePools {
			assertEqualNodePool(test_framework, (*c1.NodePools)[i], (*c2.NodePools)[i])
		}
	}

	if !assertNilEquality(test_framework, "Configuration", c1.Configuration, c2.Configuration) {
		assertEqualRancherClusterConfig(test_framework, *c1.Configuration, *c2.Configuration)
	}

	if !assertNilEquality(test_framework, "Metadata", c1.Metadata, c2.Metadata) {
		assertEqualRancherServerMetadata(test_framework, *c1.Metadata, *c2.Metadata)
	}
}

func assertEqualNodePool(test_framework *testing.T, n1 NodePool, n2 ranchersdk.NodePool) {
	assert.Equal(test_framework, n1.Name, n2.Name)
	assert.Equal(test_framework, n1.NodeCount, n2.NodeCount)
	assert.Equal(test_framework, n1.ServerType, n2.ServerType)

	if !assertNilEquality(test_framework, "SshConfig", n1.SshConfig, n2.SshConfig) {
		assertEqualSshConfig(test_framework, *n1.SshConfig, *n2.SshConfig)
	}

	if !assertNilEquality(test_framework, "Nodes", n1.Nodes, n2.Nodes) {
		assert.Equal(test_framework, len(*n1.Nodes), len(*n2.Nodes))
		for i := range *n1.Nodes {
			assertEqualNode(test_framework, (*n1.Nodes)[i], (*n2.Nodes)[i])
		}
	}

}

func assertEqualNode(test_framework *testing.T, n1 Node, n2 ranchersdk.Node) {
	assert.Equal(test_framework, n1.ServerId, n2.ServerId)
}

func assertEqualRancherClusterCertificates(test_framework *testing.T, r1 RancherClusterCertificates, r2 ranchersdk.RancherClusterCertificates) {
	assert.Equal(test_framework, r1.CaCertificate, r2.CaCertificate)
	assert.Equal(test_framework, r1.Certificate, r2.Certificate)
	assert.Equal(test_framework, r1.CertificateKey, r2.CertificateKey)
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

func assertEqualRancherServerMetadata(test_framework *testing.T, r1 RancherServerMetadata, r2 ranchersdk.RancherServerMetadata) {
	assert.Equal(test_framework, r1.Url, r2.Url)
	assert.Equal(test_framework, r1.Username, r2.Username)
	assert.Equal(test_framework, r1.Password, r2.Password)
}

func assertEqualSshConfig(test_framework *testing.T, s1 SshConfig, s2 ranchersdk.SshConfig) {
	assert.Equal(test_framework, s1.InstallDefaultKeys, s2.InstallDefaultKeys)

	if !assertNilEquality(test_framework, "Keys", s1.Keys, s2.Keys) {
		assert.Equal(test_framework, len(*s1.Keys), len(*s2.Keys))
		for i := range *s1.Keys {
			assert.Equal(test_framework, (*s1.Keys)[i], (*s2.Keys)[i])
		}
	}
	if !assertNilEquality(test_framework, "Key Ids", s1.KeyIds, s2.KeyIds) {
		assert.Equal(test_framework, len(*s1.KeyIds), len(*s2.KeyIds))
		for i := range *s1.KeyIds {
			assert.Equal(test_framework, (*s1.KeyIds)[i], (*s2.KeyIds)[i])
		}
	}
}

func assertNilEquality(test_framework *testing.T, varName string, cliVar interface{}, sdkVar interface{}) bool {
	if cliVar == nil || reflect.ValueOf(cliVar).IsNil() {
		assert.Nil(test_framework, sdkVar, "(value: "+varName+") CLI's value is nil, but not SDK's value.")
		return true
	} else if sdkVar == nil || reflect.ValueOf(sdkVar).IsNil() {
		assert.Nil(test_framework, cliVar, "(value: "+varName+") SDK's value is nil, but not CLI's value.")
		return true
	}

	return false
}
