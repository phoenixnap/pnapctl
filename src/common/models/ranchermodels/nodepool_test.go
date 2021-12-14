package ranchermodels

import (
	"fmt"
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestNodePoolToSdk(test_framework *testing.T) {
	nodePool := GenerateNodePoolCli()
	sdkNodePool := *nodePool.ToSdk()

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestFullNodePoolToSdk(test_framework *testing.T) {
	nodePool := GenerateNodePoolCli()
	nodePool.Nodes = &[]Node{GenerateNodeCli()}
	sdkNodePool := *nodePool.ToSdk()

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestNodePoolFromSdk(test_framework *testing.T) {
	sdkNodePool := GenerateNodePoolSdk()
	nodePool := NodePoolFromSdk(sdkNodePool)

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestFullNodePoolFromSdk(test_framework *testing.T) {
	sdkNodePool := GenerateNodePoolSdk()
	sdkNodePool.Nodes = &[]ranchersdk.Node{GenerateNodeSdk()}
	nodePool := NodePoolFromSdk(sdkNodePool)

	assertEqualNodePool(test_framework, nodePool, sdkNodePool)
}

func TestNodePoolsToTableStrings_nilPools(test_framework *testing.T) {
	result := NodePoolsToTableStrings(nil)
	assert.Equal(test_framework, []string{}, result)
}

func TestNodePoolsToTableStrings_emptyPoolList(test_framework *testing.T) {
	result := NodePoolsToTableStrings(&[]ranchersdk.NodePool{})
	assert.Equal(test_framework, []string{}, result)
}

func TestNodePoolsToTableStrings_withPoolList(test_framework *testing.T) {
	sdkModel_1 := GenerateNodePoolSdk()
	sdkModel_2 := GenerateNodePoolSdk()

	list := []ranchersdk.NodePool{
		sdkModel_1, sdkModel_2,
	}

	result := NodePoolsToTableStrings(&list)

	assert.Equal(test_framework, len(result), 2)
	assert.Equal(test_framework, result[0], generateNodePoolResultString(sdkModel_1))
	assert.Equal(test_framework, result[1], generateNodePoolResultString(sdkModel_2))
}

func assertEqualNodePool(test_framework *testing.T, cliNodePool NodePool, sdkNodePool ranchersdk.NodePool) {
	assert.Equal(test_framework, cliNodePool.Name, sdkNodePool.Name)
	assert.Equal(test_framework, cliNodePool.NodeCount, sdkNodePool.NodeCount)
	assert.Equal(test_framework, cliNodePool.ServerType, sdkNodePool.ServerType)

	if testutil.AssertNilEquality(test_framework, "SshConfig", cliNodePool.SshConfig, sdkNodePool.SshConfig) {
		assertEqualSshConfig(test_framework, *cliNodePool.SshConfig, *sdkNodePool.SshConfig)
	}

	if testutil.AssertNilEquality(test_framework, "Nodes", cliNodePool.Nodes, sdkNodePool.Nodes) {
		assert.Equal(test_framework, len(*cliNodePool.Nodes), len(*sdkNodePool.Nodes))
		for i := range *cliNodePool.Nodes {
			assertEqualNode(test_framework, (*cliNodePool.Nodes)[i], (*sdkNodePool.Nodes)[i])
		}
	}

}

func generateNodePoolResultString(nodePool ranchersdk.NodePool) string {
	return fmt.Sprintf("%s - %d nodes", *nodePool.Name, *nodePool.NodeCount)
}
