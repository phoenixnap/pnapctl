package networkmodels

import (
	"fmt"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
)

type PrivateNetworkServer struct {
	Id  string   `json:"id" yaml:"id"`
	Ips []string `json:"ips" yaml:"ips"`
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

	return fmt.Sprintf("ID: %s\nIps: %s\n", server.Id, server.Ips)
}
