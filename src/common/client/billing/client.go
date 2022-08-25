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
	// Rated Usages
	RatedUsageGet(queryParams billingmodels.RatedUsageGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error)
	RatedUsageMonthToDateGet(queryParams billingmodels.RatedUsageMonthToDateGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error)
	ProductsGet(queryParams billingmodels.ProductsGetQueryParams) ([]billingapisdk.ProductsGet200ResponseInner, *http.Response, error)
}

type MainClient struct {
	RatedUsageApiClient billingapisdk.RatedUsageApi
	ProductsApiClient   billingapisdk.ProductsApi
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
		ProductsApiClient:   api_client.ProductsApi,
	}
}

func (m MainClient) RatedUsageGet(queryParams billingmodels.RatedUsageGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error) {
	request := m.RatedUsageApiClient.RatedUsageGet(context.Background())
	queryParams.AttachToRequest(&request)

	return request.Execute()
}

func (m MainClient) RatedUsageMonthToDateGet(queryParams billingmodels.RatedUsageMonthToDateGetQueryParams) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error) {
	request := m.RatedUsageApiClient.RatedUsageMonthToDateGet(context.Background())
	queryParams.AttachToRequest(&request)

	return request.Execute()
}

func (m MainClient) ProductsGet(queryParams billingmodels.ProductsGetQueryParams) ([]billingapisdk.ProductsGet200ResponseInner, *http.Response, error) {
	request := m.ProductsApiClient.ProductsGet(context.Background())
	queryParams.AttachToRequest(&request)

	return request.Execute()
}
