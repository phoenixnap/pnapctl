package generators

import (
	"math/rand"
	"time"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"phoenixnap.com/pnapctl/common/models/queryparams/network"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// Private Networks

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
		Memberships:     testutil.GenN(2, GenerateNetworkMembershipSdk),
		CreatedOn:       time.Now(),
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

// Public Networks
func GeneratePublicNetworkSdk() networkapisdk.PublicNetwork {
	return networkapisdk.PublicNetwork{
		Id:          testutil.RandSeq(10),
		VlanId:      rand.Int31(),
		Memberships: testutil.GenN(2, GenerateNetworkMembershipSdk),
		Name:        testutil.RandSeq(10),
		Location:    testutil.RandSeq(10),
		Description: testutil.AsPointer(testutil.RandSeq(10)),
		CreatedOn:   time.Now(),
		IpBlocks:    testutil.GenN(2, GeneratePublicNetworkIpBlockSdk),
	}
}

func GenerateNetworkMembershipSdk() networkapisdk.NetworkMembership {
	return networkapisdk.NetworkMembership{
		ResourceId:   testutil.RandSeq(10),
		ResourceType: testutil.RandSeq(10),
		Ips:          testutil.RandListStringPointer(2),
	}
}

func GeneratePublicNetworkIpBlockSdk() networkapisdk.PublicNetworkIpBlock {
	return networkapisdk.PublicNetworkIpBlock{
		Id: testutil.RandSeq(10),
	}
}

func GeneratePublicNetworksGetQueryParams() network.PublicNetworksGetQueryParams {
	return network.PublicNetworksGetQueryParams{
		Location: testutil.AsPointer(network.AllowedLocations[0]),
	}
}

func GeneratePrivateNetworkModifySdk() networkapisdk.PrivateNetworkModify {
	return networkapisdk.PrivateNetworkModify{
		Name:            testutil.RandSeq(10),
		Description:     testutil.RandSeqPointer(10),
		LocationDefault: false,
	}
}

func GeneratePublicNetworkModifySdk() networkapisdk.PublicNetworkModify {
	return networkapisdk.PublicNetworkModify{
		Name:        testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
	}
}

func GeneratePublicNetworkCreateSdk() networkapisdk.PublicNetworkCreate {
	return networkapisdk.PublicNetworkCreate{
		Name:        testutil.RandSeq(10),
		Description: testutil.RandSeqPointer(10),
		Location:    testutil.RandSeq(10),
		IpBlocks:    testutil.GenN(10, GeneratePublicNetworkIpBlockSdk),
	}
}

func GeneratePrivateNetworkCreateSdk() networkapisdk.PrivateNetworkCreate {
	return networkapisdk.PrivateNetworkCreate{}
}
