package networkmodels

import (
	"fmt"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

type PrivateNetworkServer struct {
	Id  string   `json:"id" yaml:"id"`
	Ips []string `json:"ips" yaml:"ips"`
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

func PrivateNetworkServerToTableString(server *networksdk.PrivateNetworkServer) string {
	if server == nil {
		return ""
	}

	return fmt.Sprintf("%s : %s", server.Id, server.Ips)
}
