package networkmodels

import (
	"math/rand"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GeneratePrivateNetworkSdk() networkapisdk.PrivateNetwork {
	return networkapisdk.PrivateNetwork{
		Id:              testutil.RandSeq(10),
		Name:            testutil.RandSeq(10),
		Description:     testutil.RandSeqPointer(10),
		VlanId:          int32(rand.Int()),
		Type:            testutil.RandSeq(10),
		Location:        testutil.RandSeq(10),
		LocationDefault: false,
		Cidr:            testutil.RandSeq(10),
		Servers:         []networkapisdk.PrivateNetworkServer{},
	}
}

func GeneratePrivateNetworkListSdk(n int) []networkapisdk.PrivateNetwork {
	var privateNetworkList []networkapisdk.PrivateNetwork
	for i := 0; i < n; i++ {
		privateNetworkList = append(privateNetworkList, GeneratePrivateNetworkSdk())
	}
	return privateNetworkList
}

func GeneratePrivateNetworkServerSdk() networkapisdk.PrivateNetworkServer {
	return networkapisdk.PrivateNetworkServer{
		Id:  testutil.RandSeq(10),
		Ips: []string{testutil.RandSeq(10)},
	}
}

func GeneratePrivateNetworkCreateCli() PrivateNetworkCreate {
	return PrivateNetworkCreate{
		Name:            testutil.RandSeq(10),
		Description:     testutil.RandSeqPointer(10),
		Location:        testutil.RandSeq(10),
		LocationDefault: nil,
		Cidr:            testutil.RandSeq(10),
	}
}

func GeneratePrivateNetworkModifyCli() PrivateNetworkModify {
	return PrivateNetworkModify{
		Name:            testutil.RandSeq(10),
		Description:     testutil.RandSeqPointer(10),
		LocationDefault: false,
	}
}
