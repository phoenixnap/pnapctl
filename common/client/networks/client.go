package networks

import (
	"context"
	"net/http"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"golang.org/x/oauth2/clientcredentials"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client NetworkSdkClient

type NetworkSdkClient interface {
	// Private Networks
	PrivateNetworksGet() ([]networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkGetById(networkId string) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworksPost(privateNetworkCreate networkapisdk.PrivateNetworkCreate) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkPut(networkId string, privateNetworkUpdate networkapisdk.PrivateNetworkModify) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkDelete(networkId string) (*http.Response, error)
}

type MainClient struct {
	PrivateNetworksClient networkapisdk.PrivateNetworksApi
}

func NewMainClient(clientId string, clientSecret string) NetworkSdkClient {
	networksAPIconfiguration := networkapisdk.NewConfiguration()

	if configuration.Hostname != "" {
		networksAPIconfiguration.Servers[0].URL = configuration.Hostname
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	networksAPIconfiguration.HTTPClient = config.Client(context.Background())

	api_client := networkapisdk.NewAPIClient(networksAPIconfiguration)

	return MainClient{
		PrivateNetworksClient: api_client.PrivateNetworksApi,
	}
}

func (m MainClient) PrivateNetworksGet() ([]networkapisdk.PrivateNetwork, *http.Response, error) {
	return m.PrivateNetworksClient.PrivateNetworksGet(context.Background()).Execute()
}

func (m MainClient) PrivateNetworkGetById(networkId string) (networkapisdk.PrivateNetwork, *http.Response, error) {
	return m.PrivateNetworksClient.PrivateNetworksNetworkIdGet(context.Background(), networkId).Execute()
}

func (m MainClient) PrivateNetworksPost(privateNetworkCreate networkapisdk.PrivateNetworkCreate) (networkapisdk.PrivateNetwork, *http.Response, error) {
	return m.PrivateNetworksClient.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(privateNetworkCreate).Execute()
}

func (m MainClient) PrivateNetworkPut(networkId string, privateNetworkUpdate networkapisdk.PrivateNetworkModify) (networkapisdk.PrivateNetwork, *http.Response, error) {
	return m.PrivateNetworksClient.PrivateNetworksNetworkIdPut(context.Background(), networkId).PrivateNetworkModify(privateNetworkUpdate).Execute()
}

func (m MainClient) PrivateNetworkDelete(networkId string) (*http.Response, error) {
	return m.PrivateNetworksClient.PrivateNetworksNetworkIdDelete(context.Background(), networkId).Execute()
}
