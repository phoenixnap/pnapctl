package generators

import (
	"math/rand"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func GeneratePrivateNetwork() networkapisdk.PrivateNetwork {
	return networkapisdk.PrivateNetwork{
		Id:              randSeq(10),
		Name:            randSeq(10),
		Description:     randSeqPointer(10),
		VlanId:          int32(rand.Int()),
		Type:            randSeq(10),
		Location:        randSeq(10),
		LocationDefault: false,
		Cidr:            randSeq(10),
		Servers:         []networkapisdk.PrivateNetworkServer{},
	}
}

func GeneratePrivateNetworkCreate() networkapisdk.PrivateNetworkCreate {
	return networkapisdk.PrivateNetworkCreate{
		Name:            randSeq(10),
		Description:     randSeqPointer(10),
		Location:        randSeq(10),
		LocationDefault: nil,
		Cidr:            randSeq(10),
	}
}

func GeneratePrivateNetworkModify() networkapisdk.PrivateNetworkModify {
	return networkapisdk.PrivateNetworkModify{
		Name:            randSeq(10),
		Description:     randSeqPointer(10),
		LocationDefault: false,
	}
}

func GeneratePrivateNetworks(n int) []networkapisdk.PrivateNetwork {
	var privateNetworkList []networkapisdk.PrivateNetwork
	for i := 0; i < n; i++ {
		privateNetworkList = append(privateNetworkList, GeneratePrivateNetwork())
	}
	return privateNetworkList
}
