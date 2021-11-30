package networks

import (
	"context"
	"net/http"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client NetworkSdkClient

type NetworkSdkClient interface {
	// Private Networks
	PrivateNetworksGet(location string) ([]networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkGetById(networkId string) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworksPost(privateNetworkCreate networkapisdk.PrivateNetworkCreate) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkPut(networkId string, privateNetworkUpdate networkapisdk.PrivateNetworkModify) (networkapisdk.PrivateNetwork, *http.Response, error)
	PrivateNetworkDelete(networkId string) (*http.Response, error)
}

type MainClient struct {
	PrivateNetworksClient networkapisdk.PrivateNetworksApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) NetworkSdkClient {
	networksAPIconfiguration := networkapisdk.NewConfiguration()

	if customUrl != "" {
		networksAPIconfiguration.Servers = networkapisdk.ServerConfigurations{
			{
				URL: customUrl,
			},
		}
	}

	tokenUrl := configuration.TokenURL
	if customTokenURL != "" {
		tokenUrl = customTokenURL
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	networksAPIconfiguration.HTTPClient = config.Client(context.Background())
	networksAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := networkapisdk.NewAPIClient(networksAPIconfiguration)

	return MainClient{
		PrivateNetworksClient: api_client.PrivateNetworksApi,
	}
}

func (m MainClient) PrivateNetworksGet(location string) ([]networkapisdk.PrivateNetwork, *http.Response, error) {
	request := m.PrivateNetworksClient.PrivateNetworksGet(context.Background())

	if location != "" {
		request = request.Location(location)
	}

	return request.Execute()
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
