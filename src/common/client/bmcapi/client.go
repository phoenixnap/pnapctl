package bmcapi

import (
	"context"
	"net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client BmcApiSdkClient

type BmcApiSdkClient interface {
	//Servers
	ServersPost(serverCreate bmcapisdk.ServerCreate) (*bmcapisdk.Server, *http.Response, error)
	ServersGet([]string) ([]bmcapisdk.Server, *http.Response, error)
	ServerGetById(serverId string) (*bmcapisdk.Server, *http.Response, error)
	ServerDelete(serverId string) (*bmcapisdk.DeleteResult, *http.Response, error)
	ServerPowerOff(serverId string) (*bmcapisdk.ActionResult, *http.Response, error)
	ServerPowerOn(serverId string) (*bmcapisdk.ActionResult, *http.Response, error)
	ServerReboot(serverId string) (*bmcapisdk.ActionResult, *http.Response, error)
	ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (*bmcapisdk.ResetResult, *http.Response, error)
	ServerReserve(serverId string, serverReserve bmcapisdk.ServerReserve) (*bmcapisdk.Server, *http.Response, error)
	ServerShutdown(serverId string) (*bmcapisdk.ActionResult, *http.Response, error)
	ServerPatch(serverId string, serverPatch bmcapisdk.ServerPatch) (*bmcapisdk.Server, *http.Response, error)
	ServerTag(serverId string, tagAssignmentRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, *http.Response, error)
	ServerDeprovision(serverId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, *http.Response, error)
	ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapisdk.ServerPrivateNetwork) (*bmcapisdk.ServerPrivateNetwork, *http.Response, error)
	ServerPrivateNetworkDelete(serverId string, networkId string) (string, *http.Response, error)
	ServerPublicNetworkPost(serverId string, serverPublicNetwork bmcapisdk.ServerPublicNetwork) (*bmcapisdk.ServerPublicNetwork, *http.Response, error)
	ServerPublicNetworkDelete(serverId string, networkId string) (string, *http.Response, error)
	ServerIpBlockPost(serverId string, serverIpBlock bmcapisdk.ServerIpBlock) (*bmcapisdk.ServerIpBlock, *http.Response, error)
	ServerIpBlockDelete(serverId string, ipBlockId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, *http.Response, error)

	//Ssh Keys
	SshKeyPost(sshkeyCreate bmcapisdk.SshKeyCreate) (*bmcapisdk.SshKey, *http.Response, error)
	SshKeysGet() ([]bmcapisdk.SshKey, *http.Response, error)
	SshKeyGetById(sshKeyId string) (*bmcapisdk.SshKey, *http.Response, error)
	SshKeyPut(sshKeyId string, sshKeyUpdate bmcapisdk.SshKeyUpdate) (*bmcapisdk.SshKey, *http.Response, error)
	SshKeyDelete(sshKeyId string) (*bmcapisdk.DeleteSshKeyResult, *http.Response, error)

	//Quotas
	QuotasGet() ([]bmcapisdk.Quota, *http.Response, error)
	QuotaGetById(quotaId string) (*bmcapisdk.Quota, *http.Response, error)
	QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error)
}

type MainClient struct {
	ServersApiClient bmcapisdk.ServersApi
	SshKeysApiClient bmcapisdk.SSHKeysApi
	QuotaApiClient   bmcapisdk.QuotasApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) BmcApiSdkClient {
	bmcAPIconfiguration := bmcapisdk.NewConfiguration()

	if customUrl != "" {
		bmcAPIconfiguration.Servers = bmcapisdk.ServerConfigurations{
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

	bmcAPIconfiguration.HTTPClient = config.Client(context.Background())
	bmcAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := bmcapisdk.NewAPIClient(bmcAPIconfiguration)

	return MainClient{
		ServersApiClient: api_client.ServersApi,
		SshKeysApiClient: api_client.SSHKeysApi,
		QuotaApiClient:   api_client.QuotasApi,
	}
}

// Servers APIs
func (m MainClient) ServersPost(serverCreate bmcapisdk.ServerCreate) (*bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersPost(context.Background()).ServerCreate(serverCreate).Execute()
}

func (m MainClient) ServersGet(tags []string) ([]bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersGet(context.Background()).Tag(tags).Execute()
}

func (m MainClient) ServerGetById(serverId string) (*bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdGet(context.Background(), serverId).Execute()
}

func (m MainClient) ServerDelete(serverId string) (*bmcapisdk.DeleteResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdDelete(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOff(serverId string) (*bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOn(serverId string) (*bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReboot(serverId string) (*bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (*bmcapisdk.ResetResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
}

func (m MainClient) ServerReserve(serverId string, serverReserve bmcapisdk.ServerReserve) (*bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsReservePost(context.Background(), serverId).ServerReserve(serverReserve).Execute()
}

func (m MainClient) ServerShutdown(serverId string) (*bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPatch(serverId string, serverPatch bmcapisdk.ServerPatch) (*bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdPatch(context.Background(), serverId).ServerPatch(serverPatch).Execute()
}

func (m MainClient) ServerTag(serverId string, tagAssignmentRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdTagsPut(context.Background(), serverId).TagAssignmentRequest(tagAssignmentRequests).Execute()
}

func (m MainClient) ServerDeprovision(serverId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsDeprovisionPost(context.Background(), serverId).RelinquishIpBlock(relinquishIpBlock).Execute()
}

func (m MainClient) ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapisdk.ServerPrivateNetwork) (*bmcapisdk.ServerPrivateNetwork, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdPrivateNetworksPost(context.Background(), serverId).ServerPrivateNetwork(serverPrivateNetwork).Execute()
}

func (m MainClient) ServerPrivateNetworkDelete(serverId string, networkId string) (string, *http.Response, error) {
	return m.ServersApiClient.DeletePrivateNetwork(context.Background(), serverId, networkId).Execute()
}

func (m MainClient) ServerPublicNetworkPost(serverId string, serverPublicNetwork bmcapisdk.ServerPublicNetwork) (*bmcapisdk.ServerPublicNetwork, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdPublicNetworksPost(context.Background(), serverId).ServerPublicNetwork(serverPublicNetwork).Execute()
}

func (m MainClient) ServerPublicNetworkDelete(serverId string, networkId string) (string, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdPublicNetworksDelete(context.Background(), serverId, networkId).Execute()
}

func (m MainClient) ServerIpBlockPost(serverId string, serverIpBlock bmcapisdk.ServerIpBlock) (*bmcapisdk.ServerIpBlock, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdIpBlocksPost(context.Background(), serverId).ServerIpBlock(serverIpBlock).Execute()
}

func (m MainClient) ServerIpBlockDelete(serverId string, ipBlockId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdIpBlocksIpBlockIdDelete(context.Background(), serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()
}

// SSH Key APIs
func (m MainClient) SshKeyPost(sshKeyCreate bmcapisdk.SshKeyCreate) (*bmcapisdk.SshKey, *http.Response, error) {
	return m.SshKeysApiClient.SshKeysPost(context.Background()).SshKeyCreate(sshKeyCreate).Execute()
}

func (m MainClient) SshKeysGet() ([]bmcapisdk.SshKey, *http.Response, error) {
	return m.SshKeysApiClient.SshKeysGet(context.Background()).Execute()
}

func (m MainClient) SshKeyGetById(sshKeyId string) (*bmcapisdk.SshKey, *http.Response, error) {
	return m.SshKeysApiClient.SshKeysSshKeyIdGet(context.Background(), sshKeyId).Execute()
}

func (m MainClient) SshKeyPut(sshKeyId string, sshKeyUpdate bmcapisdk.SshKeyUpdate) (*bmcapisdk.SshKey, *http.Response, error) {
	return m.SshKeysApiClient.SshKeysSshKeyIdPut(context.Background(), sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()
}

func (m MainClient) SshKeyDelete(sshKeyId string) (*bmcapisdk.DeleteSshKeyResult, *http.Response, error) {
	return m.SshKeysApiClient.SshKeysSshKeyIdDelete(context.Background(), sshKeyId).Execute()
}

// Quota APIs
func (m MainClient) QuotasGet() ([]bmcapisdk.Quota, *http.Response, error) {
	return m.QuotaApiClient.QuotasGet(context.Background()).Execute()
}

func (m MainClient) QuotaGetById(quotaId string) (*bmcapisdk.Quota, *http.Response, error) {
	return m.QuotaApiClient.QuotasQuotaIdGet(context.Background(), quotaId).Execute()
}

func (m MainClient) QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error) {
	return m.QuotaApiClient.QuotasQuotaIdActionsRequestEditPost(context.Background(), quotaId).QuotaEditLimitRequest(quotaEditRequest).Execute()
}
