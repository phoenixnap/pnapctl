package audit

import (
	"context"
	"net/http"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"golang.org/x/oauth2/clientcredentials"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client AuditSdkClient

type AuditSdkClient interface {
	// Events
	EventsGet() ([]auditapisdk.Event, *http.Response, error)
}

type MainClient struct {
	EventsApiClient auditapisdk.EventsApi
}

func NewMainClient(clientId string, clientSecret string) AuditSdkClient {
	auditAPIconfiguration := auditapisdk.NewConfiguration()

	if configuration.Hostname != "" {
		auditAPIconfiguration.Servers[0].URL = configuration.Hostname
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	auditAPIconfiguration.HTTPClient = config.Client(context.Background())

	api_client := auditapisdk.NewAPIClient(auditAPIconfiguration)

	return MainClient{
		EventsApiClient: api_client.EventsApi,
	}
}

// Events APIs
func (m MainClient) EventsGet() ([]auditapisdk.Event, *http.Response, error) {
	return m.EventsApiClient.EventsGet(context.Background()).Execute()
}
