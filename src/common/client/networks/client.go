package networks

import (
	"context"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client NetworkSdkClient

type NetworkSdkClient interface {
	// Private Networks
	PrivateNetworksGet(location string) ([]networkapisdk.PrivateNetwork, error)
	PrivateNetworkGetById(networkId string) (*networkapisdk.PrivateNetwork, error)
	PrivateNetworksPost(privateNetworkCreate networkapisdk.PrivateNetworkCreate) (*networkapisdk.PrivateNetwork, error)
	PrivateNetworkPut(networkId string, privateNetworkUpdate networkapisdk.PrivateNetworkModify) (*networkapisdk.PrivateNetwork, error)
	PrivateNetworkDelete(networkId string) error

	PublicNetworksGet(location string) ([]networkapisdk.PublicNetwork, error)
	PublicNetworkGetById(networkId string) (*networkapisdk.PublicNetwork, error)
	PublicNetworksPost(publicNetworkCreate networkapisdk.PublicNetworkCreate) (*networkapisdk.PublicNetwork, error)
	PublicNetworkDelete(networkId string) error
	PublicNetworkPatch(networkId string, publicNetworkPatch networkapisdk.PublicNetworkModify) (*networkapisdk.PublicNetwork, error)
	PublicNetworkIpBlockPost(networkId string, idBlockCreate networkapisdk.PublicNetworkIpBlock) (*networkapisdk.PublicNetworkIpBlock, error)
	PublicNetworkIpBlockDelete(networkId string, ipBlockId string) (string, error)
}

type MainClient struct {
	PrivateNetworksClient networkapisdk.PrivateNetworksApi
	PublicNetworksClient  networkapisdk.PublicNetworksApi
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
	networksAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := networkapisdk.NewAPIClient(networksAPIconfiguration)

	return MainClient{
		PrivateNetworksClient: api_client.PrivateNetworksApi,
		PublicNetworksClient:  api_client.PublicNetworksApi,
	}
}

func (m MainClient) PrivateNetworksGet(location string) ([]networkapisdk.PrivateNetwork, error) {
	request := m.PrivateNetworksClient.PrivateNetworksGet(context.Background())

	if location != "" {
		request = request.Location(location)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) PrivateNetworkGetById(networkId string) (*networkapisdk.PrivateNetwork, error) {
	return client.HandleResponse(m.PrivateNetworksClient.PrivateNetworksNetworkIdGet(context.Background(), networkId).Execute())
}

func (m MainClient) PrivateNetworksPost(privateNetworkCreate networkapisdk.PrivateNetworkCreate) (*networkapisdk.PrivateNetwork, error) {
	return client.HandleResponse(m.PrivateNetworksClient.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(privateNetworkCreate).Execute())
}

func (m MainClient) PrivateNetworkPut(networkId string, privateNetworkUpdate networkapisdk.PrivateNetworkModify) (*networkapisdk.PrivateNetwork, error) {
	return client.HandleResponse(m.PrivateNetworksClient.PrivateNetworksNetworkIdPut(context.Background(), networkId).PrivateNetworkModify(privateNetworkUpdate).Execute())
}

func (m MainClient) PrivateNetworkDelete(networkId string) error {
	return client.HandleResponseWithoutBody(m.PrivateNetworksClient.PrivateNetworksNetworkIdDelete(context.Background(), networkId).Execute())
}

func (m MainClient) PublicNetworksGet(location string) ([]networkapisdk.PublicNetwork, error) {
	request := m.PublicNetworksClient.PublicNetworksGet(context.Background())

	if !client.IsZeroValue(location) {
		request = request.Location(location)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) PublicNetworkGetById(networkId string) (*networkapisdk.PublicNetwork, error) {
	return client.HandleResponse(m.PublicNetworksClient.PublicNetworksNetworkIdGet(context.Background(), networkId).Execute())
}

func (m MainClient) PublicNetworksPost(publicNetworkCreate networkapisdk.PublicNetworkCreate) (*networkapisdk.PublicNetwork, error) {
	return client.HandleResponse(m.PublicNetworksClient.PublicNetworksPost(context.Background()).PublicNetworkCreate(publicNetworkCreate).Execute())
}

func (m MainClient) PublicNetworkDelete(networkId string) error {
	return client.HandleResponseWithoutBody(m.PublicNetworksClient.PublicNetworksNetworkIdDelete(context.Background(), networkId).Execute())
}

func (m MainClient) PublicNetworkPatch(networkId string, publicNetworkPatch networkapisdk.PublicNetworkModify) (*networkapisdk.PublicNetwork, error) {
	return client.HandleResponse(m.PublicNetworksClient.PublicNetworksNetworkIdPatch(context.Background(), networkId).PublicNetworkModify(publicNetworkPatch).Execute())
}

func (m MainClient) PublicNetworkIpBlockPost(networkId string, idBlockCreate networkapisdk.PublicNetworkIpBlock) (*networkapisdk.PublicNetworkIpBlock, error) {
	return client.HandleResponse(m.PublicNetworksClient.PublicNetworksNetworkIdIpBlocksPost(context.Background(), networkId).PublicNetworkIpBlock(idBlockCreate).Execute())
}

func (m MainClient) PublicNetworkIpBlockDelete(networkId string, ipBlockId string) (string, error) {
	return client.HandleResponse(m.PublicNetworksClient.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(context.Background(), networkId, ipBlockId).Execute())
}
