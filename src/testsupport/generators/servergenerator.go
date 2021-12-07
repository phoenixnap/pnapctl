package generators

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

func GenerateServers(n int) []bmcapisdk.Server {
	var serverlist []bmcapisdk.Server
	for i := 0; i < n; i++ {
		serverlist = append(serverlist, GenerateServer())
	}
	return serverlist
}

func GenerateServer() bmcapisdk.Server {
	provisionedOn := time.Now()
	return bmcapisdk.Server{
		Id:                 testutil.RandSeq(10),
		Status:             testutil.RandSeq(10),
		Hostname:           testutil.RandSeq(10),
		Description:        testutil.RandSeqPointer(10),
		Os:                 testutil.RandSeq(10),
		Type:               testutil.RandSeq(10),
		Location:           testutil.RandSeq(10),
		Cpu:                testutil.RandSeq(10),
		CpuCount:           int32(rand.Int()),
		CoresPerCpu:        int32(rand.Int()),
		CpuFrequency:       rand.Float32(),
		Ram:                testutil.RandSeq(10),
		Storage:            testutil.RandSeq(10),
		PrivateIpAddresses: []string{},
		PublicIpAddresses:  []string{},
		ReservationId:      testutil.RandSeqPointer(10),
		PricingModel:       testutil.RandSeq(10),
		Password:           testutil.RandSeqPointer(10),
		NetworkType:        testutil.RandSeqPointer(10),
		ClusterId:          testutil.RandSeqPointer(10),
		Tags:               nil,
		ProvisionedOn:      &provisionedOn,
		OsConfiguration:    nil,
	}
}

func GenerateServerCreate() servermodels.ServerCreate {
	return servermodels.ServerCreate{
		Hostname:              testutil.RandSeq(10),
		Description:           testutil.RandSeqPointer(10),
		Os:                    testutil.RandSeq(10),
		Type:                  testutil.RandSeq(10),
		Location:              testutil.RandSeq(10),
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		ReservationId:         testutil.RandSeqPointer(10),
		PricingModel:          testutil.RandSeqPointer(10),
		NetworkType:           testutil.RandSeqPointer(10),
		OsConfiguration:       nil,
		Tags:                  nil,
		NetworkConfiguration:  nil,
	}
}

func GenerateClusters(n int) []ranchersdk.Cluster {
	var clusterlist []ranchersdk.Cluster
	for i := 0; i < n; i++ {
		clusterlist = append(clusterlist, GenerateCluster())
	}
	return clusterlist
}

func GenerateCluster() ranchersdk.Cluster {
	return ranchersdk.Cluster{
		Id:                    testutil.RandSeqPointer(10),
		Name:                  testutil.RandSeqPointer(10),
		Description:           testutil.RandSeqPointer(10),
		Location:              testutil.RandSeq(10),
		InitialClusterVersion: testutil.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     testutil.RandSeqPointer(10),
	}
}

func GenerateBmcApiDeleteResult() bmcapisdk.DeleteResult {
	return bmcapisdk.DeleteResult{
		Result:   testutil.RandSeq(10),
		ServerId: testutil.RandSeq(10),
	}
}

func GenerateActionResult() bmcapisdk.ActionResult {
	return bmcapisdk.ActionResult{
		Result: testutil.RandSeq(10),
	}
}

func GenerateServerReset() bmcapisdk.ServerReset {
	return bmcapisdk.ServerReset{
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		OsConfiguration:       nil,
	}
}

func GenerateResetResult() bmcapisdk.ResetResult {
	return bmcapisdk.ResetResult{
		Result:          testutil.RandSeq(10),
		Password:        nil,
		OsConfiguration: nil,
	}
}

func GenerateServerPatch() bmcapisdk.ServerPatch {
	return bmcapisdk.ServerPatch{
		Hostname:    testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequest() bmcapisdk.TagAssignmentRequest {
	return bmcapisdk.TagAssignmentRequest{
		Name:  testutil.RandSeq(10),
		Value: testutil.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequests(n int) []bmcapisdk.TagAssignmentRequest {
	var list []bmcapisdk.TagAssignmentRequest
	for i := 0; i < n; i++ {
		list = append(list, GenerateTagAssignmentRequest())
	}
	return list
}

func GenerateServerReserve() bmcapisdk.ServerReserve {
	return bmcapisdk.ServerReserve{
		PricingModel: "ONE_MONTH_RESERVATION",
	}
}

func GenerateServerPrivateNetwork() bmcapisdk.ServerPrivateNetwork {
	dhcp := false
	return bmcapisdk.ServerPrivateNetwork{
		Id:                testutil.RandSeq(10),
		Ips:               testutil.RandListStringPointer(10),
		Dhcp:              &dhcp,
		StatusDescription: testutil.RandSeqPointer(10),
	}
}
