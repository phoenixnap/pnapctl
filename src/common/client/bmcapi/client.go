package bmcapi

import (
	"context"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client BmcApiSdkClient

type BmcApiSdkClient interface {
	//Servers
	ServersPost(serverCreate bmcapisdk.ServerCreate, force bool) (*bmcapisdk.Server, error)
	ServersGet([]string) ([]bmcapisdk.Server, error)
	ServerGetById(serverId string) (*bmcapisdk.Server, error)
	ServerDelete(serverId string) (*bmcapisdk.DeleteResult, error)
	ServerPowerOff(serverId string) (*bmcapisdk.ActionResult, error)
	ServerPowerOn(serverId string) (*bmcapisdk.ActionResult, error)
	ServerReboot(serverId string) (*bmcapisdk.ActionResult, error)
	ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (*bmcapisdk.ResetResult, error)
	ServerReserve(serverId string, serverReserve bmcapisdk.ServerReserve) (*bmcapisdk.Server, error)
	ServerShutdown(serverId string) (*bmcapisdk.ActionResult, error)
	ServerPatch(serverId string, serverPatch bmcapisdk.ServerPatch) (*bmcapisdk.Server, error)
	ServerTag(serverId string, tagAssignmentRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, error)
	ServerDeprovision(serverId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, error)
	ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapisdk.ServerPrivateNetwork, force bool) (*bmcapisdk.ServerPrivateNetwork, error)
	ServerPrivateNetworkDelete(serverId string, networkId string) (string, error)
	ServerPrivateNetworkPatch(serverId string, networkId string, serverNetworkUpdate bmcapisdk.ServerNetworkUpdate, force bool) (*bmcapisdk.ServerPrivateNetwork, error)
	ServerPublicNetworkPost(serverId string, serverPublicNetwork bmcapisdk.ServerPublicNetwork, force bool) (*bmcapisdk.ServerPublicNetwork, error)
	ServerPublicNetworkDelete(serverId string, networkId string) (string, error)
	ServerPublicNetworkPatch(serverId string, networkId string, serverNetworkUpdate bmcapisdk.ServerNetworkUpdate, force bool) (*bmcapisdk.ServerPublicNetwork, error)
	ServerIpBlockPost(serverId string, serverIpBlock bmcapisdk.ServerIpBlock) (*bmcapisdk.ServerIpBlock, error)
	ServerIpBlockDelete(serverId string, ipBlockId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, error)

	//Ssh Keys
	SshKeyPost(sshkeyCreate bmcapisdk.SshKeyCreate) (*bmcapisdk.SshKey, error)
	SshKeysGet() ([]bmcapisdk.SshKey, error)
	SshKeyGetById(sshKeyId string) (*bmcapisdk.SshKey, error)
	SshKeyPut(sshKeyId string, sshKeyUpdate bmcapisdk.SshKeyUpdate) (*bmcapisdk.SshKey, error)
	SshKeyDelete(sshKeyId string) (*bmcapisdk.DeleteSshKeyResult, error)

	//Quotas
	QuotasGet() ([]bmcapisdk.Quota, error)
	QuotaGetById(quotaId string) (*bmcapisdk.Quota, error)
	QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) error
}

type MainClient struct {
	ServersApiClient bmcapisdk.ServersAPI
	SshKeysApiClient bmcapisdk.SSHKeysAPI
	QuotaApiClient   bmcapisdk.QuotasAPI
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
	bmcAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := bmcapisdk.NewAPIClient(bmcAPIconfiguration)

	return MainClient{
		ServersApiClient: api_client.ServersAPI,
		SshKeysApiClient: api_client.SSHKeysAPI,
		QuotaApiClient:   api_client.QuotasAPI,
	}
}

// Servers APIs
func (m MainClient) ServersPost(serverCreate bmcapisdk.ServerCreate, force bool) (*bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersPost(context.Background()).ServerCreate(serverCreate).Force(force).Execute())
}

func (m MainClient) ServersGet(tags []string) ([]bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersGet(context.Background()).Tag(tags).Execute())
}

func (m MainClient) ServerGetById(serverId string) (*bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdGet(context.Background(), serverId).Execute())
}

func (m MainClient) ServerDelete(serverId string) (*bmcapisdk.DeleteResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdDelete(context.Background(), serverId).Execute())
}

func (m MainClient) ServerPowerOff(serverId string) (*bmcapisdk.ActionResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute())
}

func (m MainClient) ServerPowerOn(serverId string) (*bmcapisdk.ActionResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute())
}

func (m MainClient) ServerReboot(serverId string) (*bmcapisdk.ActionResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute())
}

func (m MainClient) ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (*bmcapisdk.ResetResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute())
}

func (m MainClient) ServerReserve(serverId string, serverReserve bmcapisdk.ServerReserve) (*bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsReservePost(context.Background(), serverId).ServerReserve(serverReserve).Execute())
}

func (m MainClient) ServerShutdown(serverId string) (*bmcapisdk.ActionResult, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute())
}

func (m MainClient) ServerPatch(serverId string, serverPatch bmcapisdk.ServerPatch) (*bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPatch(context.Background(), serverId).ServerPatch(serverPatch).Execute())
}

func (m MainClient) ServerTag(serverId string, tagAssignmentRequests []bmcapisdk.TagAssignmentRequest) (*bmcapisdk.Server, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdTagsPut(context.Background(), serverId).TagAssignmentRequest(tagAssignmentRequests).Execute())
}

func (m MainClient) ServerDeprovision(serverId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdActionsDeprovisionPost(context.Background(), serverId).RelinquishIpBlock(relinquishIpBlock).Execute())
}

func (m MainClient) ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapisdk.ServerPrivateNetwork, force bool) (*bmcapisdk.ServerPrivateNetwork, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPrivateNetworksPost(context.Background(), serverId).ServerPrivateNetwork(serverPrivateNetwork).Force(force).Execute())
}

func (m MainClient) ServerPrivateNetworkDelete(serverId string, networkId string) (string, error) {
	return client.HandleResponse(m.ServersApiClient.DeletePrivateNetwork(context.Background(), serverId, networkId).Execute())
}

func (m MainClient) ServerPrivateNetworkPatch(serverId string, networkId string, serverNetworkUpdate bmcapisdk.ServerNetworkUpdate, force bool) (*bmcapisdk.ServerPrivateNetwork, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPrivateNetworksPatch(context.Background(), serverId, networkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute())
}

func (m MainClient) ServerPublicNetworkPost(serverId string, serverPublicNetwork bmcapisdk.ServerPublicNetwork, force bool) (*bmcapisdk.ServerPublicNetwork, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPublicNetworksPost(context.Background(), serverId).ServerPublicNetwork(serverPublicNetwork).Force(force).Execute())
}

func (m MainClient) ServerPublicNetworkDelete(serverId string, networkId string) (string, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPublicNetworksDelete(context.Background(), serverId, networkId).Execute())
}

func (m MainClient) ServerPublicNetworkPatch(serverId string, networkId string, serverNetworkUpdate bmcapisdk.ServerNetworkUpdate, force bool) (*bmcapisdk.ServerPublicNetwork, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdPublicNetworksPatch(context.Background(), serverId, networkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute())
}

func (m MainClient) ServerIpBlockPost(serverId string, serverIpBlock bmcapisdk.ServerIpBlock) (*bmcapisdk.ServerIpBlock, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdIpBlocksPost(context.Background(), serverId).ServerIpBlock(serverIpBlock).Execute())
}

func (m MainClient) ServerIpBlockDelete(serverId string, ipBlockId string, relinquishIpBlock bmcapisdk.RelinquishIpBlock) (string, error) {
	return client.HandleResponse(m.ServersApiClient.ServersServerIdIpBlocksIpBlockIdDelete(context.Background(), serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute())
}

// SSH Key APIs
func (m MainClient) SshKeyPost(sshKeyCreate bmcapisdk.SshKeyCreate) (*bmcapisdk.SshKey, error) {
	return client.HandleResponse(m.SshKeysApiClient.SshKeysPost(context.Background()).SshKeyCreate(sshKeyCreate).Execute())
}

func (m MainClient) SshKeysGet() ([]bmcapisdk.SshKey, error) {
	return client.HandleResponse(m.SshKeysApiClient.SshKeysGet(context.Background()).Execute())
}

func (m MainClient) SshKeyGetById(sshKeyId string) (*bmcapisdk.SshKey, error) {
	return client.HandleResponse(m.SshKeysApiClient.SshKeysSshKeyIdGet(context.Background(), sshKeyId).Execute())
}

func (m MainClient) SshKeyPut(sshKeyId string, sshKeyUpdate bmcapisdk.SshKeyUpdate) (*bmcapisdk.SshKey, error) {
	return client.HandleResponse(m.SshKeysApiClient.SshKeysSshKeyIdPut(context.Background(), sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute())
}

func (m MainClient) SshKeyDelete(sshKeyId string) (*bmcapisdk.DeleteSshKeyResult, error) {
	return client.HandleResponse(m.SshKeysApiClient.SshKeysSshKeyIdDelete(context.Background(), sshKeyId).Execute())
}

// Quota APIs
func (m MainClient) QuotasGet() ([]bmcapisdk.Quota, error) {
	return client.HandleResponse(m.QuotaApiClient.QuotasGet(context.Background()).Execute())
}

func (m MainClient) QuotaGetById(quotaId string) (*bmcapisdk.Quota, error) {
	return client.HandleResponse(m.QuotaApiClient.QuotasQuotaIdGet(context.Background(), quotaId).Execute())
}

func (m MainClient) QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) error {
	return client.HandleResponseWithoutBody(m.QuotaApiClient.QuotasQuotaIdActionsRequestEditPost(context.Background(), quotaId).QuotaEditLimitRequest(quotaEditRequest).Execute())
}
