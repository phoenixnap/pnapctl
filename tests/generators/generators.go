package generators

import (
	"math/rand"
	"time"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/commands/create/server"
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

func GenerateServers(n int) []bmcapi.Server {
	var serverlist []bmcapi.Server
	for i := 0; i < n; i++ {
		serverlist = append(serverlist, GenerateServer())
	}
	return serverlist
}

func GenerateServer() bmcapi.Server {
	provisionedOn := time.Now()
	return bmcapi.Server{
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

func GenerateServerCreate() server.ServerCreate {
	return server.ServerCreate{
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

func GenerateDeleteResult() bmcapi.DeleteResult {
	return bmcapi.DeleteResult{
		Result:   randSeq(10),
		ServerId: randSeq(10),
	}
}

func GenerateActionResult() bmcapi.ActionResult {
	return bmcapi.ActionResult{
		Result: randSeq(10),
	}
}

func GenerateServerReset() bmcapi.ServerReset {
	return bmcapi.ServerReset{
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		OsConfiguration:       nil,
	}
}

func GenerateResetResult() bmcapi.ResetResult {
	return bmcapi.ResetResult{
		Result:          randSeq(10),
		Password:        nil,
		OsConfiguration: nil,
	}
}
