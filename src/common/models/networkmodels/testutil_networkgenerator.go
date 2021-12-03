package networkmodels

import (
	"math/rand"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/tests/generators"
)

func GeneratePrivateNetworkSdk() networkapisdk.PrivateNetwork {
	return networkapisdk.PrivateNetwork{
		Id:              generators.RandSeq(10),
		Name:            generators.RandSeq(10),
		Description:     generators.RandSeqPointer(10),
		VlanId:          int32(rand.Int()),
		Type:            generators.RandSeq(10),
		Location:        generators.RandSeq(10),
		LocationDefault: false,
		Cidr:            generators.RandSeq(10),
		Servers:         []networkapisdk.PrivateNetworkServer{},
	}
}

func GeneratePrivateNetworkServerSdk() networkapisdk.PrivateNetworkServer {
	return networkapisdk.PrivateNetworkServer{
		Id:  generators.RandSeq(10),
		Ips: []string{generators.RandSeq(10)},
	}
}

func GeneratePrivateNetworkCreateCli() PrivateNetworkCreate {
	return PrivateNetworkCreate{
		Name:            generators.RandSeq(10),
		Description:     generators.RandSeqPointer(10),
		Location:        generators.RandSeq(10),
		LocationDefault: nil,
		Cidr:            generators.RandSeq(10),
	}
}

func GeneratePrivateNetworkModifyCli() PrivateNetworkModify {
	return PrivateNetworkModify{
		Name:            generators.RandSeq(10),
		Description:     generators.RandSeqPointer(10),
		LocationDefault: false,
	}
}

func GeneratePrivateNetworkListSdk(n int) []networkapisdk.PrivateNetwork {
	var privateNetworkList []networkapisdk.PrivateNetwork
	for i := 0; i < n; i++ {
		privateNetworkList = append(privateNetworkList, GeneratePrivateNetworkSdk())
	}
	return privateNetworkList
}
