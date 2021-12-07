package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
)

func TestNodeToSdk(test_framework *testing.T) {
	node := GenerateNodeCli()
	sdkNode := *node.ToSdk()

	assertEqualNode(test_framework, node, sdkNode)
}

func TestNodeFromSdk(test_framework *testing.T) {
	sdkNode := GenerateNodeSdk()
	node := NodeFromSdk(sdkNode)

	assertEqualNode(test_framework, node, sdkNode)
}

func assertEqualNode(test_framework *testing.T, n1 Node, n2 ranchersdk.Node) {
	assert.Equal(test_framework, n1.ServerId, n2.ServerId)
}
