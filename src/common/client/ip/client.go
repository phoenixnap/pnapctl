package ip

import (
	"context"
	"net/http"

	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client IpSdkClient

type IpSdkClient interface {
	// Ip Blocks
	IpBlockPost(ipBlockCreate ipapisdk.IpBlockCreate) (*ipapisdk.IpBlock, *http.Response, error)
	IpBlocksGet() ([]ipapisdk.IpBlock, *http.Response, error)
	IpBlocksGetById(ipBlockId string) (*ipapisdk.IpBlock, *http.Response, error)
	IpBlocksIpBlockIdDelete(ipBlockId string) (*ipapisdk.DeleteIpBlockResult, *http.Response, error)
	IpBlocksIpBlockIdPatch(ipBlockId string, ipBlockPatch ipapisdk.IpBlockPatch) (*ipapisdk.IpBlock, *http.Response, error)
	IpBlocksIpBlockIdTagsPut(ipBlockId string) (*ipapisdk.IpBlock, *http.Response, error)
}

type MainClient struct {
	IpBlocksApiClient ipapisdk.IPBlocksApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) IpSdkClient {
	ipAPIconfiguration := ipapisdk.NewConfiguration()

	if customUrl != "" {
		ipAPIconfiguration.Servers = ipapisdk.ServerConfigurations{
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

	ipAPIconfiguration.HTTPClient = config.Client(context.Background())
	ipAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := ipapisdk.NewAPIClient(ipAPIconfiguration)

	return MainClient{
		IpBlocksApiClient: api_client.IPBlocksApi,
	}
}

// IP APIs
func (m MainClient) IpBlockPost(ipBlockCreate ipapisdk.IpBlockCreate) (*ipapisdk.IpBlock, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksPost(context.Background()).IpBlockCreate(ipBlockCreate).Execute()
}

func (m MainClient) IpBlocksGet() ([]ipapisdk.IpBlock, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksGet(context.Background()).Execute()
}

func (m MainClient) IpBlocksGetById(ipBlockId string) (*ipapisdk.IpBlock, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksIpBlockIdGet(context.Background(), ipBlockId).Execute()
}

func (m MainClient) IpBlocksIpBlockIdDelete(ipBlockId string) (*ipapisdk.DeleteIpBlockResult, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksIpBlockIdDelete(context.Background(), ipBlockId).Execute()
}

func (m MainClient) IpBlocksIpBlockIdPatch(ipBlockId string, ipBlockPatch ipapisdk.IpBlockPatch) (*ipapisdk.IpBlock, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksIpBlockIdPatch(context.Background(), ipBlockId).IpBlockPatch(ipBlockPatch).Execute()
}

func (m MainClient) IpBlocksIpBlockIdTagsPut(ipBlockId string) (*ipapisdk.IpBlock, *http.Response, error) {
	return m.IpBlocksApiClient.IpBlocksIpBlockIdTagsPut(context.Background(), ipBlockId).Execute()
}
