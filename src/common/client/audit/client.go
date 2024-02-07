package audit

import (
	"context"
	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client AuditSdkClient

type AuditSdkClient interface {
	// Events
	EventsGet(from string, to string, limit int, order string, username string, verb string, uri string) ([]auditapisdk.Event, error)
}

type MainClient struct {
	EventsApiClient auditapisdk.EventsAPI
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
	auditAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := auditapisdk.NewAPIClient(auditAPIconfiguration)

	return MainClient{
		EventsApiClient: api_client.EventsAPI,
	}
}

// Events APIs
func (m MainClient) EventsGet(from string, to string, limit int, order string, username string, verb string, uri string) ([]auditapisdk.Event, error) {
	request := m.EventsApiClient.EventsGet(context.Background())

	if date := client.ParseDate(from); date != nil {
		request = request.From(*date)
	}
	if date := client.ParseDate(to); date != nil {
		request = request.To(*date)
	}
	if !client.IsZeroValue(limit) {
		request = request.Limit(int32(limit))
	}
	if !client.IsZeroValue(order) {
		request = request.Order(order)
	}
	if !client.IsZeroValue(username) {
		request = request.Username(username)
	}
	if !client.IsZeroValue(verb) {
		request = request.Verb(verb)
	}
	if !client.IsZeroValue(uri) {
		request = request.Uri(uri)
	}

	return client.HandleResponse(request.Execute())
}
