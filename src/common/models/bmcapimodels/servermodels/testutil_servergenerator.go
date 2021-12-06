package servermodels

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	generators "phoenixnap.com/pnapctl/testsupport/generators"
)

func GenerateServerListSdk(n int) []bmcapisdk.Server {
	var serverlist []bmcapisdk.Server
	for i := 0; i < n; i++ {
		serverlist = append(serverlist, GenerateServerSdk())
	}
	return serverlist
}

func GenerateServerSdk() bmcapisdk.Server {
	provisionedOn := time.Now()
	return bmcapisdk.Server{
		Id:                 generators.RandSeq(10),
		Status:             generators.RandSeq(10),
		Hostname:           generators.RandSeq(10),
		Description:        generators.RandSeqPointer(10),
		Os:                 generators.RandSeq(10),
		Type:               generators.RandSeq(10),
		Location:           generators.RandSeq(10),
		Cpu:                generators.RandSeq(10),
		CpuCount:           int32(rand.Int()),
		CoresPerCpu:        int32(rand.Int()),
		CpuFrequency:       rand.Float32(),
		Ram:                generators.RandSeq(10),
		Storage:            generators.RandSeq(10),
		PrivateIpAddresses: []string{},
		PublicIpAddresses:  []string{},
		ReservationId:      generators.RandSeqPointer(10),
		PricingModel:       generators.RandSeq(10),
		Password:           generators.RandSeqPointer(10),
		NetworkType:        generators.RandSeqPointer(10),
		ClusterId:          generators.RandSeqPointer(10),
		Tags:               nil,
		ProvisionedOn:      &provisionedOn,
		OsConfiguration:    nil,
	}
}

func GenerateServerCreateCli() ServerCreate {
	return ServerCreate{
		Hostname:              generators.RandSeq(10),
		Description:           generators.RandSeqPointer(10),
		Os:                    generators.RandSeq(10),
		Type:                  generators.RandSeq(10),
		Location:              generators.RandSeq(10),
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		ReservationId:         generators.RandSeqPointer(10),
		PricingModel:          generators.RandSeqPointer(10),
		NetworkType:           generators.RandSeqPointer(10),
		OsConfiguration:       nil,
		Tags:                  nil,
		NetworkConfiguration:  nil,
	}
}

func GenerateClusterListSdk(n int) []ranchersdk.Cluster {
	var clusterlist []ranchersdk.Cluster
	for i := 0; i < n; i++ {
		clusterlist = append(clusterlist, GenerateClusterSdk())
	}
	return clusterlist
}

func GenerateClusterSdk() ranchersdk.Cluster {
	return ranchersdk.Cluster{
		Id:                    generators.RandSeqPointer(10),
		Name:                  generators.RandSeqPointer(10),
		Description:           generators.RandSeqPointer(10),
		Location:              generators.RandSeq(10),
		InitialClusterVersion: generators.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     generators.RandSeqPointer(10),
	}
}

func GenerateRancherDeleteResultSdk() ranchersdk.DeleteResult {
	return ranchersdk.DeleteResult{
		Result:    generators.RandSeq(10),
		ClusterId: generators.RandSeqPointer(10),
	}
}

func GenerateBmcApiDeleteResultSdk() bmcapisdk.DeleteResult {
	return bmcapisdk.DeleteResult{
		Result:   generators.RandSeq(10),
		ServerId: generators.RandSeq(10),
	}
}

func GenerateActionResultSdk() bmcapisdk.ActionResult {
	return bmcapisdk.ActionResult{
		Result: generators.RandSeq(10),
	}
}

func GenerateServerResetSdk() bmcapisdk.ServerReset {
	return bmcapisdk.ServerReset{
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		OsConfiguration:       nil,
	}
}

func GenerateResetResultSdk() bmcapisdk.ResetResult {
	return bmcapisdk.ResetResult{
		Result:          generators.RandSeq(10),
		Password:        nil,
		OsConfiguration: nil,
	}
}

func GenerateServerPatchSdk() bmcapisdk.ServerPatch {
	return bmcapisdk.ServerPatch{
		Hostname:    generators.RandSeqPointer(10),
		Description: generators.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestSdk() bmcapisdk.TagAssignmentRequest {
	return bmcapisdk.TagAssignmentRequest{
		Name:  generators.RandSeq(10),
		Value: generators.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestCli() TagAssignmentRequest {
	return TagAssignmentRequest{
		Name:  generators.RandSeq(10),
		Value: generators.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestListSdk(n int) []bmcapisdk.TagAssignmentRequest {
	var list []bmcapisdk.TagAssignmentRequest
	for i := 0; i < n; i++ {
		list = append(list, GenerateTagAssignmentRequestSdk())
	}
	return list
}

func GenerateServerReserveSdk() bmcapisdk.ServerReserve {
	return bmcapisdk.ServerReserve{
		PricingModel: "ONE_MONTH_RESERVATION",
	}
}

func GenerateServerPrivateNetworkSdk() bmcapisdk.ServerPrivateNetwork {
	dhcp := false
	return bmcapisdk.ServerPrivateNetwork{
		Id:                generators.RandSeq(10),
		Ips:               generators.RandListStringPointer(10),
		Dhcp:              &dhcp,
		StatusDescription: generators.RandSeqPointer(10),
	}
}
