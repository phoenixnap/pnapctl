package servermodels

import bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"

type ServerIpBlock struct {
	Id     string `yaml:"id" json:"id"`
	VlanId *int32 `yaml:"vlanId,omitempty" json:"vlanId,omitempty"`
}

func (serverIpBlock ServerIpBlock) ToSdk() bmcapisdk.ServerIpBlock {
	var serverIpBlockSdk = bmcapisdk.ServerIpBlock{
		Id:     serverIpBlock.Id,
		VlanId: serverIpBlock.VlanId,
	}

	return serverIpBlockSdk
}

func mapServerIpBlocksToSdk(serverIpBlocks []ServerIpBlock) []bmcapisdk.ServerIpBlock {
	if serverIpBlocks == nil {
		return nil
	}

	var serverIpBlocksSdk []bmcapisdk.ServerIpBlock

	for _, serverIpBlock := range serverIpBlocks {
		serverIpBlocksSdk = append(serverIpBlocksSdk, serverIpBlock.ToSdk())
	}

	return serverIpBlocksSdk
}

func mapServerIpBlocksToCLI(serverIpBlocks []bmcapisdk.ServerIpBlock) []ServerIpBlock {
	if serverIpBlocks == nil {
		return nil
	}

	var serverIpBlocksCLI []ServerIpBlock

	for _, serverIpBlock := range serverIpBlocks {
		serverIpBlocksCLI = append(serverIpBlocksCLI, ServerIpBlock{
			Id:     serverIpBlock.Id,
			VlanId: serverIpBlock.VlanId,
		})
	}

	return serverIpBlocksCLI
}
