package servermodels

import (
	"math/rand"
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"

	"phoenixnap.com/pnapctl/testsupport/testutil"
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
	tagAssignments := GenerateTagAssignmentListSdk(2)
	osConfiguration := GenerateOsConfigurationSdk()
	networkConfiguration := GenerateNetworkConfigurationSdk()
	return bmcapisdk.Server{
		Id:                   testutil.RandSeq(10),
		Status:               testutil.RandSeq(10),
		Hostname:             testutil.RandSeq(10),
		Description:          testutil.RandSeqPointer(10),
		Os:                   testutil.RandSeq(10),
		Type:                 testutil.RandSeq(10),
		Location:             testutil.RandSeq(10),
		Cpu:                  testutil.RandSeq(10),
		CpuCount:             int32(rand.Int()),
		CoresPerCpu:          int32(rand.Int()),
		CpuFrequency:         rand.Float32(),
		Ram:                  testutil.RandSeq(10),
		Storage:              testutil.RandSeq(10),
		PrivateIpAddresses:   *testutil.RandListStringPointer(10),
		PublicIpAddresses:    *testutil.RandListStringPointer(10),
		ReservationId:        testutil.RandSeqPointer(10),
		PricingModel:         testutil.RandSeq(10),
		Password:             testutil.RandSeqPointer(10),
		NetworkType:          testutil.RandSeqPointer(10),
		ClusterId:            testutil.RandSeqPointer(10),
		Tags:                 &tagAssignments,
		ProvisionedOn:        &provisionedOn,
		OsConfiguration:      &osConfiguration,
		NetworkConfiguration: networkConfiguration,
	}
}

func GenerateServerCreateCli() ServerCreate {
	flag := false
	osConfiguration := GenerateOsConfigurationCli()
	tagAssignmentRequests := []TagAssignmentRequest{
		GenerateTagAssignmentRequestCli(), GenerateTagAssignmentRequestCli(),
	}
	networkConfiguration := GenerateNetworkConfigurationCli()
	return ServerCreate{
		Hostname:              testutil.RandSeq(10),
		Description:           testutil.RandSeqPointer(10),
		Os:                    testutil.RandSeq(10),
		Type:                  testutil.RandSeq(10),
		Location:              testutil.RandSeq(10),
		InstallDefaultSshKeys: &flag,
		SshKeys:               testutil.RandListStringPointer(10),
		SshKeyIds:             testutil.RandListStringPointer(10),
		ReservationId:         testutil.RandSeqPointer(10),
		PricingModel:          testutil.RandSeqPointer(10),
		NetworkType:           testutil.RandSeqPointer(10),
		OsConfiguration:       &osConfiguration,
		Tags:                  &tagAssignmentRequests,
		NetworkConfiguration:  &networkConfiguration,
	}
}

func GenerateBmcApiDeleteResultSdk() bmcapisdk.DeleteResult {
	return bmcapisdk.DeleteResult{
		Result:   testutil.RandSeq(10),
		ServerId: testutil.RandSeq(10),
	}
}

func GenerateActionResultSdk() bmcapisdk.ActionResult {
	return bmcapisdk.ActionResult{
		Result: testutil.RandSeq(10),
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

func GenerateServerResetCli() ServerReset {
	var flag = false
	return ServerReset{
		InstallDefaultSshKeys: &flag,
		SshKeys:               testutil.RandListStringPointer(10),
		SshKeyIds:             testutil.RandListStringPointer(10),
		OsConfiguration:       GenerateOsConfigurationMapCli(),
	}
}

func GenerateResetResultSdk() bmcapisdk.ResetResult {
	return bmcapisdk.ResetResult{
		Result:          testutil.RandSeq(10),
		Password:        nil,
		OsConfiguration: nil,
	}
}

func GenerateServerPatchSdk() bmcapisdk.ServerPatch {
	return bmcapisdk.ServerPatch{
		Hostname:    testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
	}
}

func GenerateServerPatchCli() ServerPatch {
	return ServerPatch{
		Hostname:    testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestSdk() bmcapisdk.TagAssignmentRequest {
	return bmcapisdk.TagAssignmentRequest{
		Name:  testutil.RandSeq(10),
		Value: testutil.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestCli() TagAssignmentRequest {
	return TagAssignmentRequest{
		Name:  testutil.RandSeq(10),
		Value: testutil.RandSeqPointer(10),
	}
}

func GenerateTagAssignmentRequestListSdk(n int) []bmcapisdk.TagAssignmentRequest {
	var list []bmcapisdk.TagAssignmentRequest
	for i := 0; i < n; i++ {
		list = append(list, GenerateTagAssignmentRequestSdk())
	}
	return list
}

func GenerateTagAssignmentSdk() bmcapisdk.TagAssignment {
	return bmcapisdk.TagAssignment{
		Id:           testutil.RandSeq(10),
		Name:         testutil.RandSeq(10),
		Value:        testutil.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagAssignmentListSdk(n int) []bmcapisdk.TagAssignment {
	var list []bmcapisdk.TagAssignment
	for i := 0; i < n; i++ {
		list = append(list, GenerateTagAssignmentSdk())
	}
	return list
}

func GenerateTagAssignmentCli() TagAssignment {
	return TagAssignment{
		Id:           testutil.RandSeq(10),
		Name:         testutil.RandSeq(10),
		Value:        testutil.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateServerReserveSdk() bmcapisdk.ServerReserve {
	return bmcapisdk.ServerReserve{
		PricingModel: "ONE_MONTH_RESERVATION",
	}
}

func GenerateServerReserveCli() ServerReserve {
	return ServerReserve{
		PricingModel: "ONE_MONTH_RESERVATION",
	}
}

func GenerateNetworkConfigurationCli() NetworkConfiguration {
	privateNetworkConfiguration := GeneratePrivateNetworkConfigurationCli()
	ipBlocksConfiguration := GenerateIpBlocksConfigurationCli()
	return NetworkConfiguration{
		PrivateNetworkConfiguration: &privateNetworkConfiguration,
		IpBlocksConfiguration:       ipBlocksConfiguration,
	}
}

func GenerateNetworkConfigurationSdk() bmcapisdk.NetworkConfiguration {
	privateNetworkConfiguration := GeneratePrivateNetworkConfigurationSdk()
	ipBlocksConfiguration := GenerateIpBlockConfigurationSdk()
	return bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: &privateNetworkConfiguration,
		IpBlocksConfiguration:       ipBlocksConfiguration,
	}
}

func GeneratePrivateNetworkConfigurationCli() PrivateNetworkConfiguration {
	serverPrivateNetworks := GenerateServerPrivateNetworkListCli(2)
	return PrivateNetworkConfiguration{
		GatewayAddress:    testutil.RandSeqPointer(10),
		ConfigurationType: testutil.RandSeqPointer(10),
		PrivateNetworks:   &serverPrivateNetworks,
	}
}

func GeneratePrivateNetworkConfigurationSdk() bmcapisdk.PrivateNetworkConfiguration {
	serverPrivateNetworks := GenerateServerPrivateNetworkListSdk(2)
	return bmcapisdk.PrivateNetworkConfiguration{
		GatewayAddress:    testutil.RandSeqPointer(10),
		ConfigurationType: testutil.RandSeqPointer(10),
		PrivateNetworks:   &serverPrivateNetworks,
	}
}

func GenerateServerPrivateNetworkListCli(n int) []ServerPrivateNetwork {
	var list []ServerPrivateNetwork
	for i := 0; i < n; i++ {
		list = append(list, GenerateServerPrivateNetworkCli())
	}
	return list
}

func GenerateServerPrivateNetworkListSdk(n int) []bmcapisdk.ServerPrivateNetwork {
	var list []bmcapisdk.ServerPrivateNetwork
	for i := 0; i < n; i++ {
		list = append(list, GenerateServerPrivateNetworkSdk())
	}
	return list
}

func GenerateServerPrivateNetworkCli() ServerPrivateNetwork {
	dhcp := false
	return ServerPrivateNetwork{
		Id:                testutil.RandSeq(10),
		Ips:               testutil.RandListStringPointer(10),
		Dhcp:              &dhcp,
		StatusDescription: testutil.RandSeqPointer(10),
	}
}

func GenerateServerPrivateNetworkSdk() bmcapisdk.ServerPrivateNetwork {
	dhcp := false
	return bmcapisdk.ServerPrivateNetwork{
		Id:                testutil.RandSeq(10),
		Ips:               testutil.RandListStringPointer(10),
		Dhcp:              &dhcp,
		StatusDescription: testutil.RandSeqPointer(10),
	}
}

func GenerateOsConfigurationCli() OsConfiguration {
	return OsConfiguration{
		Windows:                    GenerateOsConfigurationWindowsCli(),
		RootPassword:               testutil.RandSeqPointer(10),
		ManagementUiUrl:            testutil.RandSeqPointer(10),
		ManagementAccessAllowedIps: testutil.RandListStringPointer(10),
	}
}

func GenerateOsConfigurationSdk() bmcapisdk.OsConfiguration {
	return bmcapisdk.OsConfiguration{
		Windows:                    GenerateOsConfigurationWindowsSdk(),
		RootPassword:               testutil.RandSeqPointer(10),
		ManagementUiUrl:            testutil.RandSeqPointer(10),
		ManagementAccessAllowedIps: testutil.RandListStringPointer(10),
	}
}

func GenerateOsConfigurationWindowsCli() *OsConfigurationWindows {
	return &OsConfigurationWindows{
		RdpAllowedIps: testutil.RandListStringPointer(10),
	}
}

func GenerateOsConfigurationWindowsSdk() *bmcapisdk.OsConfigurationWindows {
	return &bmcapisdk.OsConfigurationWindows{
		RdpAllowedIps: testutil.RandListStringPointer(10),
	}
}

func GenerateOsConfigurationMapCli() *OsConfigurationMap {
	return &OsConfigurationMap{
		Windows: GenerateOsConfigurationWindowsCli(),
		Esxi:    GenerateOsConfigurationMapEsxiCli(),
	}
}

func GenerateOsConfigurationMapEsxiCli() *OsConfigurationMapEsxi {
	return &OsConfigurationMapEsxi{
		RootPassword:               testutil.RandSeqPointer(10),
		ManagementUiUrl:            testutil.RandSeqPointer(10),
		ManagementAccessAllowedIps: testutil.RandListStringPointer(10),
	}
}

func GenerateServerIpBlockCli() ServerIpBlock {
	return ServerIpBlock{
		Id:     testutil.RandSeq(10),
		VlanId: testutil.RanNumberPointer(),
	}
}

func GenerateServerIpBlockSdk() bmcapisdk.ServerIpBlock {
	return bmcapisdk.ServerIpBlock{
		Id:     testutil.RandSeq(10),
		VlanId: testutil.RanNumberPointer(),
	}
}

func GenerateServerIpBlockListCli(n int) *[]ServerIpBlock {
	var list []ServerIpBlock
	for i := 0; i < n; i++ {
		list = append(list, GenerateServerIpBlockCli())
	}
	return &list
}

func GenerateServerIpBlockListSdk(n int) *[]bmcapisdk.ServerIpBlock {
	var list []bmcapisdk.ServerIpBlock
	for i := 0; i < n; i++ {
		list = append(list, GenerateServerIpBlockSdk())
	}
	return &list
}

func GenerateIpBlocksConfigurationCli() *IpBlocksConfiguration {
	return &IpBlocksConfiguration{
		ConfigurationType: testutil.RandSeqPointer(10),
		IpBlocks:          GenerateServerIpBlockListCli(2),
	}
}

func GenerateIpBlockConfigurationSdk() *bmcapisdk.IpBlocksConfiguration {
	return &bmcapisdk.IpBlocksConfiguration{
		ConfigurationType: testutil.RandSeqPointer(10),
		IpBlocks:          GenerateServerIpBlockListSdk(2),
	}
}
