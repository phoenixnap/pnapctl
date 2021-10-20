package client

import (
	"context"
	"net/http"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"golang.org/x/oauth2/clientcredentials"
)

var Client SdkClient

type SdkClient interface {
	ServersPost(serverCreate bmcapi.ServerCreate) (bmcapi.Server, *http.Response, error)
	ServersGet() ([]bmcapi.Server, *http.Response, error)
	ServerGetById(serverId string) (bmcapi.Server, *http.Response, error)
	ServerDelete(serverId string) (bmcapi.DeleteResult, *http.Response, error)
	ServerPowerOff(serverId string) (bmcapi.ActionResult, *http.Response, error)
	ServerPowerOn(serverId string) (bmcapi.ActionResult, *http.Response, error)
	ServerReboot(serverId string) (bmcapi.ActionResult, *http.Response, error)
	ServerReset(serverId string, serverReset bmcapi.ServerReset) (bmcapi.ResetResult, *http.Response, error)
	ServerShutdown(serverId string) (bmcapi.ActionResult, *http.Response, error)
}

type MainClient struct {
	BmcApiClient bmcapi.DefaultApi
}

func NewMainClient(clientId string, clientSecret string) SdkClient {
	configuration := bmcapi.NewConfiguration()

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token",
		Scopes:       []string{"bmc", "bmc.read"},
	}

	configuration.HTTPClient = config.Client(context.Background())

	api_client := bmcapi.NewAPIClient(configuration)

	return MainClient{
		BmcApiClient: api_client.DefaultApi,
	}
}

func (m MainClient) ServersPost(serverCreate bmcapi.ServerCreate) (bmcapi.Server, *http.Response, error) {
	return m.BmcApiClient.ServersPost(context.Background()).ServerCreate(serverCreate).Execute()
}

func (m MainClient) ServersGet() ([]bmcapi.Server, *http.Response, error) {
	return m.BmcApiClient.ServersGet(context.Background()).Execute()
}

func (m MainClient) ServerGetById(serverId string) (bmcapi.Server, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdGet(context.Background(), serverId).Execute()
}

func (m MainClient) ServerDelete(serverId string) (bmcapi.DeleteResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdDelete(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOff(serverId string) (bmcapi.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerPowerOn(serverId string) (bmcapi.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReboot(serverId string) (bmcapi.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
}

func (m MainClient) ServerReset(serverId string, serverReset bmcapi.ServerReset) (bmcapi.ResetResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
}

func (m MainClient) ServerShutdown(serverId string) (bmcapi.ActionResult, *http.Response, error) {
	return m.BmcApiClient.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
}
