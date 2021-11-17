package bmcapi

import (
	"context"
	"net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"golang.org/x/oauth2/clientcredentials"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client BmcApiSdkClient

type BmcApiSdkClient interface {
	//Servers
	ServersPost(serverCreate bmcapisdk.ServerCreate) (bmcapisdk.Server, *http.Response, error)
	ServersGet() ([]bmcapisdk.Server, *http.Response, error)
	ServerGetById(serverId string) (bmcapisdk.Server, *http.Response, error)
	ServerDelete(serverId string) (bmcapisdk.DeleteResult, *http.Response, error)
	ServerPowerOff(serverId string) (bmcapisdk.ActionResult, *http.Response, error)
	ServerPowerOn(serverId string) (bmcapisdk.ActionResult, *http.Response, error)
	ServerReboot(serverId string) (bmcapisdk.ActionResult, *http.Response, error)
	ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (bmcapisdk.ResetResult, *http.Response, error)
	ServerShutdown(serverId string) (bmcapisdk.ActionResult, *http.Response, error)

	//Quotas
	QuotasGet() ([]bmcapisdk.Quota, *http.Response, error)
	QuotaGetById(quotaId string) (bmcapisdk.Quota, *http.Response, error)
	QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error)
}

type MainClient struct {
	ServersApiClient bmcapisdk.ServersApi
	QuotaApiClient   bmcapisdk.QuotasApi
}

func NewMainClient(clientId string, clientSecret string) BmcApiSdkClient {
	bmcAPIconfiguration := bmcapisdk.NewConfiguration()

	if configuration.BmcApiHostname != "" {
		bmcAPIconfiguration.Servers[0].URL = configuration.BmcApiHostname
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	bmcAPIconfiguration.HTTPClient = config.Client(context.Background())

	api_client := bmcapisdk.NewAPIClient(bmcAPIconfiguration)

	return MainClient{
		ServersApiClient: api_client.ServersApi,
		QuotaApiClient:   api_client.QuotasApi,
	}
}

//Servers APIs
func (m MainClient) ServersPost(serverCreate bmcapisdk.ServerCreate) (bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersPost(context.Background()).ServerCreate(serverCreate).Execute()
}

func (m MainClient) ServersGet() ([]bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersGet(context.Background()).Execute()
}

func (m MainClient) ServerGetById(serverId string) (bmcapisdk.Server, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdGet(context.Background(), serverId).Execute()
}

func (m MainClient) ServerDelete(serverId string) (bmcapisdk.DeleteResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdDelete(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOff(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOn(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReboot(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (bmcapisdk.ResetResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
}

func (m MainClient) ServerShutdown(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.ServersApiClient.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
}

// Quota APIs
func (m MainClient) QuotasGet() ([]bmcapisdk.Quota, *http.Response, error) {
	return m.QuotaApiClient.QuotasGet(context.Background()).Execute()
}

func (m MainClient) QuotaGetById(quotaId string) (bmcapisdk.Quota, *http.Response, error) {
	return m.QuotaApiClient.QuotasQuotaIdGet(context.Background(), quotaId).Execute()
}

func (m MainClient) QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error) {
	return m.QuotaApiClient.QuotasQuotaIdActionsRequestEditPost(context.Background(), quotaId).QuotaEditLimitRequest(quotaEditRequest).Execute()
}
