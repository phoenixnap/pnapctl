package bmcapi

import (
	"context"
	"net/http"

	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
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
	BmcApiClient bmcapisdk.DefaultApi
}

func NewMainClient(clientId string, clientSecret string) BmcApiSdkClient {
	bmcAPIconfiguration := bmcapisdk.NewConfiguration()

	//TODO: Since the CLI is only working with bmc api for the time being
	// we only have one server in the array bmcAPIconfiguration.Servers.
	// We will need to revisit this configuration when we introduce other
	// apis (i.e. billing, tags, audit, etc...)
	if configuration.Hostname != "" {
		bmcAPIconfiguration.Servers[0].URL = configuration.Hostname
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
		BmcApiClient: api_client.DefaultApi,
	}
}

//Servers APIs
func (m MainClient) ServersPost(serverCreate bmcapisdk.ServerCreate) (bmcapisdk.Server, *http.Response, error) {
	return m.BmcApiClient.ServersPost(context.Background()).ServerCreate(serverCreate).Execute()
}

func (m MainClient) ServersGet() ([]bmcapisdk.Server, *http.Response, error) {
	return m.BmcApiClient.ServersGet(context.Background()).Execute()
}

func (m MainClient) ServerGetById(serverId string) (bmcapisdk.Server, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdGet(context.Background(), serverId).Execute()
}

func (m MainClient) ServerDelete(serverId string) (bmcapisdk.DeleteResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdDelete(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOff(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOn(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReboot(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (bmcapisdk.ResetResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
}

func (m MainClient) ServerShutdown(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
}

// Quota APIs
func (m MainClient) QuotasGet() ([]bmcapisdk.Quota, *http.Response, error) {
	return m.BmcApiClient.QuotasGet(context.Background()).Execute()
}

func (m MainClient) QuotaGetById(quotaId string) (bmcapisdk.Quota, *http.Response, error) {
	return m.BmcApiClient.QuotasQuotaIdGet(context.Background(), quotaId).Execute()
}

func (m MainClient) QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error) {
	return m.BmcApiClient.QuotasQuotaIdActionsRequestEditPost(context.Background(), quotaId).QuotaEditLimitRequest(quotaEditRequest).Execute()
}
