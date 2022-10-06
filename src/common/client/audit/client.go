package audit

import (
	"context"
	"net/http"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/models/queryparams/audit"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client AuditSdkClient

type AuditSdkClient interface {
	// Events
	EventsGet(queryParams audit.EventsGetQueryParams) ([]auditapisdk.Event, *http.Response, error)
}

type MainClient struct {
	EventsApiClient auditapisdk.EventsApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) AuditSdkClient {
	auditAPIconfiguration := auditapisdk.NewConfiguration()

	if customUrl != "" {
		auditAPIconfiguration.Servers = auditapisdk.ServerConfigurations{
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

	auditAPIconfiguration.HTTPClient = config.Client(context.Background())
	auditAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := auditapisdk.NewAPIClient(auditAPIconfiguration)

	return MainClient{
		EventsApiClient: api_client.EventsApi,
	}
}

// Events APIs
func (m MainClient) EventsGet(queryParams audit.EventsGetQueryParams) ([]auditapisdk.Event, *http.Response, error) {
	request := m.EventsApiClient.EventsGet(context.Background())
	request = *queryParams.AttachToRequest(request)

	return request.Execute()
}
