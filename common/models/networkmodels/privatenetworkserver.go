package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

type PrivateNetworkServer struct {
	Id  string
	Ips []string
}

func (server *PrivateNetworkServer) toSdk() networksdk.PrivateNetworkServer {
	return networksdk.PrivateNetworkServer{
		Id:  server.Id,
		Ips: server.Ips,
	}
}

func PrivateNetworkServerFromSdk(server *networksdk.PrivateNetworkServer) *PrivateNetworkServer {
	if server == nil {
		return nil
	}

	return &PrivateNetworkServer{
		Id:  server.Id,
		Ips: server.Ips,
	}
}
