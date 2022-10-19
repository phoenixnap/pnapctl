package ip

import (
	"context"

	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client IpSdkClient

type IpSdkClient interface {
	// Ip Blocks
	IpBlockPost(ipBlockCreate ipapisdk.IpBlockCreate) (*ipapisdk.IpBlock, error)
	IpBlocksGet([]string) ([]ipapisdk.IpBlock, error)
	IpBlocksGetById(ipBlockId string) (*ipapisdk.IpBlock, error)
	IpBlocksIpBlockIdDelete(ipBlockId string) (*ipapisdk.DeleteIpBlockResult, error)
	IpBlocksIpBlockIdPatch(ipBlockId string, ipBlockPatch ipapisdk.IpBlockPatch) (*ipapisdk.IpBlock, error)
	IpBlocksIpBlockIdTagsPut(ipBlockId string, tag []ipapisdk.TagAssignmentRequest) (*ipapisdk.IpBlock, error)
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
	ipAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := ipapisdk.NewAPIClient(ipAPIconfiguration)

	return MainClient{
		IpBlocksApiClient: api_client.IPBlocksApi,
	}
}

// IP APIs
func (m MainClient) IpBlockPost(ipBlockCreate ipapisdk.IpBlockCreate) (*ipapisdk.IpBlock, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksPost(context.Background()).IpBlockCreate(ipBlockCreate).Execute())
}

func (m MainClient) IpBlocksGet(tags []string) ([]ipapisdk.IpBlock, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksGet(context.Background()).Tag(tags).Execute())
}

func (m MainClient) IpBlocksGetById(ipBlockId string) (*ipapisdk.IpBlock, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksIpBlockIdGet(context.Background(), ipBlockId).Execute())
}

func (m MainClient) IpBlocksIpBlockIdDelete(ipBlockId string) (*ipapisdk.DeleteIpBlockResult, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksIpBlockIdDelete(context.Background(), ipBlockId).Execute())
}

func (m MainClient) IpBlocksIpBlockIdPatch(ipBlockId string, ipBlockPatch ipapisdk.IpBlockPatch) (*ipapisdk.IpBlock, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksIpBlockIdPatch(context.Background(), ipBlockId).IpBlockPatch(ipBlockPatch).Execute())
}

func (m MainClient) IpBlocksIpBlockIdTagsPut(ipBlockId string, tag []ipapisdk.TagAssignmentRequest) (*ipapisdk.IpBlock, error) {
	return client.HandleResponse(m.IpBlocksApiClient.IpBlocksIpBlockIdTagsPut(context.Background(), ipBlockId).TagAssignmentRequest(tag).Execute())
}
