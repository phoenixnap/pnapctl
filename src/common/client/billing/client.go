package billing

import (
	"context"
	"net/http"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client BillingSdkClient

type BillingSdkClient interface {
	// Events
	RatedUsageGet(queryParams billingmodels.RatedUsageGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error)
}

type MainClient struct {
	RatedUsageApiClient billingapisdk.RatedUsageApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) BillingSdkClient {
	billingAPIconfiguration := billingapisdk.NewConfiguration()

	if customUrl != "" {
		billingAPIconfiguration.Servers = billingapisdk.ServerConfigurations{
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

	billingAPIconfiguration.HTTPClient = config.Client(context.Background())
	billingAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := billingapisdk.NewAPIClient(billingAPIconfiguration)

	return MainClient{
		RatedUsageApiClient: api_client.RatedUsageApi,
	}
}

func (m MainClient) RatedUsageGet(queryParams billingmodels.RatedUsageGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error) {
	request := m.RatedUsageApiClient.RatedUsageGet(context.Background())
	request = *queryParams.AttachToRequest(request)

	return request.Execute()
}
