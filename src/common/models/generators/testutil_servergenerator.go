package generators

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func GenerateServerSdk() bmcapisdk.Server {
	return Generate[bmcapisdk.Server]()
}

func GenerateBmcApiDeleteResultSdk() bmcapisdk.DeleteResult {
	return Generate[bmcapisdk.DeleteResult]()
}

func GenerateActionResultSdk() bmcapisdk.ActionResult {
	return Generate[bmcapisdk.ActionResult]()
}

func GenerateServerCreateSdk() bmcapisdk.ServerCreate {
	return Generate[bmcapisdk.ServerCreate]()
}

func GenerateServerResetSdk() bmcapisdk.ServerReset {
	return Generate[bmcapisdk.ServerReset]()
}

func GenerateResetResultSdk() bmcapisdk.ResetResult {
	return Generate[bmcapisdk.ResetResult]()
}

func GenerateServerPatchSdk() bmcapisdk.ServerPatch {
	return Generate[bmcapisdk.ServerPatch]()
}

func GenerateRelinquishIpBlockSdk() bmcapisdk.RelinquishIpBlock {
	return Generate[bmcapisdk.RelinquishIpBlock]()
}

func GenerateTagAssignmentRequestSdk() bmcapisdk.TagAssignmentRequest {
	return Generate[bmcapisdk.TagAssignmentRequest]()
}

func GenerateServerReserveSdk() bmcapisdk.ServerReserve {
	return Generate[bmcapisdk.ServerReserve]()
}

func GenerateServerPrivateNetworkSdk() bmcapisdk.ServerPrivateNetwork {
	return Generate[bmcapisdk.ServerPrivateNetwork]()
}

func GenerateServerPublicNetworkSdk() bmcapisdk.ServerPublicNetwork {
	return Generate[bmcapisdk.ServerPublicNetwork]()
}

func GenerateServerIpBlockSdk() bmcapisdk.ServerIpBlock {
	return Generate[bmcapisdk.ServerIpBlock]()
}
