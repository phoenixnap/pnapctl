package generators

import (
	"math/rand"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func GeneratePrivateNetwork() networkapisdk.PrivateNetwork {
	return networkapisdk.PrivateNetwork{
		Id:              RandSeq(10),
		Name:            RandSeq(10),
		Description:     RandSeqPointer(10),
		VlanId:          int32(rand.Int()),
		Type:            RandSeq(10),
		Location:        RandSeq(10),
		LocationDefault: false,
		Cidr:            RandSeq(10),
		Servers:         []networkapisdk.PrivateNetworkServer{},
	}
}

func GeneratePrivateNetworkCreate() networkapisdk.PrivateNetworkCreate {
	return networkapisdk.PrivateNetworkCreate{
		Name:            RandSeq(10),
		Description:     RandSeqPointer(10),
		Location:        RandSeq(10),
		LocationDefault: nil,
		Cidr:            RandSeq(10),
	}
}

func GeneratePrivateNetworkModify() networkapisdk.PrivateNetworkModify {
	return networkapisdk.PrivateNetworkModify{
		Name:            RandSeq(10),
		Description:     RandSeqPointer(10),
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
