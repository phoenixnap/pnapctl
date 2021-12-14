package ranchermodels

import ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"

type Node struct {
	ServerId *string `json:"serverId" yaml:"serverId"`
}

func (n Node) ToSdk() *ranchersdk.Node {
	return &ranchersdk.Node{
		ServerId: n.ServerId,
	}
}

func NodeFromSdk(node ranchersdk.Node) Node {
	return Node{
		ServerId: node.ServerId,
	}
}
