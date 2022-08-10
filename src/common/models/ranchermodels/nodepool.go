package ranchermodels

import (
	"fmt"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

type NodePool struct {
	Name       *string    `json:"name" yaml:"name"`
	NodeCount  *int32     `json:"nodeCount" yaml:"nodeCount"`
	ServerType *string    `json:"serverType" yaml:"serverType"`
	SshConfig  *SshConfig `json:"sshConfig" yaml:"sshConfig"`
	Nodes      []Node     `json:"nodes" yaml:"nodes"`
}

func (n NodePool) ToSdk() *ranchersdk.NodePool {
	var nodes []ranchersdk.Node

	if n.Nodes != nil {
		nodes = []ranchersdk.Node{}
		for _, node := range n.Nodes {
			nodes = append(nodes, *node.ToSdk())
		}
	}

	return &ranchersdk.NodePool{
		Name:       n.Name,
		NodeCount:  n.NodeCount,
		ServerType: n.ServerType,
		SshConfig:  n.SshConfig.ToSdk(),
		Nodes:      nodes,
	}
}

func NodePoolFromSdk(nodepool ranchersdk.NodePool) NodePool {
	var nodes []Node

	if nodepool.Nodes != nil {
		nodes = []Node{}
		for _, node := range nodepool.Nodes {
			nodes = append(nodes, NodeFromSdk(node))
		}
	}

	return NodePool{
		Name:       nodepool.Name,
		NodeCount:  nodepool.NodeCount,
		ServerType: nodepool.ServerType,
		SshConfig:  SshConfigFromSdk(nodepool.SshConfig),
		Nodes:      nodes,
	}
}

func NodePoolsToTableStrings(pools []ranchersdk.NodePool) []string {
	if pools == nil {
		return []string{}
	}

	var strings = []string{}

	for _, pool := range pools {
		strings = append(strings, fmt.Sprintf("%s - %d nodes", *pool.Name, *pool.NodeCount))
	}

	return strings
}
