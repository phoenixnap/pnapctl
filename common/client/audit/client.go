package audit

import (
	"context"
	"net/http"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnap-cli/common/models/auditmodels"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client AuditSdkClient

type AuditSdkClient interface {
	// Events
	EventsGet(queryParams auditmodels.EventsGetQueryParams) ([]auditapisdk.Event, *http.Response, error)
}

type MainClient struct {
	EventsApiClient auditapisdk.EventsApi
}

func NewMainClient(clientId string, clientSecret string) AuditSdkClient {
	auditAPIconfiguration := auditapisdk.NewConfiguration()

	if configuration.AuditHostname != "" {
		auditAPIconfiguration.Servers[0].URL = configuration.AuditHostname
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
func (m MainClient) EventsGet(queryParams auditmodels.EventsGetQueryParams) ([]auditapisdk.Event, *http.Response, error) {
	request := m.EventsApiClient.EventsGet(context.Background())

	if queryParams.From != nil {
		request = request.From(*queryParams.From)
	}
	if queryParams.To != nil {
		request = request.To(*queryParams.To)
	}
	if queryParams.Limit != 0 {
		request = request.Limit(int32(queryParams.Limit))
	}
	if queryParams.Order != "" {
		request = request.Order(queryParams.Order)
	}
	if queryParams.Username != "" {
		request = request.Username(queryParams.Username)
	}
	if queryParams.Verb != "" {
		request = request.Verb(queryParams.Verb)
	}
	if queryParams.Uri != "" {
		request = request.Uri(queryParams.Uri)
	}

	return request.Execute()
}
