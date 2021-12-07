package generators

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randSeqPointer(n int) *string {
	random := randSeq(n)
	return &random
}

func randListStringPointer(n int) *[]string {
	b := []string{}
	for i := range b {
		b[i] = randSeq(10)
	}
	return &b
}

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
		Id:                 randSeq(10),
		Status:             randSeq(10),
		Hostname:           randSeq(10),
		Description:        randSeqPointer(10),
		Os:                 randSeq(10),
		Type:               randSeq(10),
		Location:           randSeq(10),
		Cpu:                randSeq(10),
		CpuCount:           int32(rand.Int()),
		CoresPerCpu:        int32(rand.Int()),
		CpuFrequency:       rand.Float32(),
		Ram:                randSeq(10),
		Storage:            randSeq(10),
		PrivateIpAddresses: []string{},
		PublicIpAddresses:  []string{},
		ReservationId:      randSeqPointer(10),
		PricingModel:       randSeq(10),
		Password:           randSeqPointer(10),
		NetworkType:        randSeqPointer(10),
		ClusterId:          randSeqPointer(10),
		Tags:               nil,
		ProvisionedOn:      &provisionedOn,
		OsConfiguration:    nil,
	}
}

func GenerateServerCreate() bmcapimodels.ServerCreate {
	return bmcapimodels.ServerCreate{
		Hostname:              randSeq(10),
		Description:           randSeqPointer(10),
		Os:                    randSeq(10),
		Type:                  randSeq(10),
		Location:              randSeq(10),
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		ReservationId:         randSeqPointer(10),
		PricingModel:          randSeqPointer(10),
		NetworkType:           randSeqPointer(10),
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
		Id:                    randSeqPointer(10),
		Name:                  randSeqPointer(10),
		Description:           randSeqPointer(10),
		Location:              randSeq(10),
		InitialClusterVersion: randSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     randSeqPointer(10),
	}
}

func GenerateRancherDeleteResult() ranchersdk.DeleteResult {
	return ranchersdk.DeleteResult{
		Result:    randSeq(10),
		ClusterId: randSeqPointer(10),
	}
}

func GenerateBmcApiDeleteResult() bmcapisdk.DeleteResult {
	return bmcapisdk.DeleteResult{
		Result:   randSeq(10),
		ServerId: randSeq(10),
	}
}

func GenerateActionResult() bmcapisdk.ActionResult {
	return bmcapisdk.ActionResult{
		Result: randSeq(10),
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
		Result:          randSeq(10),
		Password:        nil,
		OsConfiguration: nil,
	}
}

func GenerateServerPatch() bmcapisdk.ServerPatch {
	return bmcapisdk.ServerPatch{
		Hostname:    randSeqPointer(10),
		Description: randSeqPointer(10),
	}
}

func GenerateTagAssignmentRequest() bmcapisdk.TagAssignmentRequest {
	return bmcapisdk.TagAssignmentRequest{
		Name:  randSeq(10),
		Value: randSeqPointer(10),
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
		Id:                randSeq(10),
		Ips:               randListStringPointer(10),
		Dhcp:              &dhcp,
		StatusDescription: randSeqPointer(10),
	}
}