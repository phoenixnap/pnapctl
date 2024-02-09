package locations

import (
	"context"

	locationapisdk "github.com/phoenixnap/go-sdk-bmc/locationapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client LocationSdkClient

type LocationSdkClient interface {
	// Locations
	LocationsGet(location, productCategory string) ([]locationapisdk.Location, error)
}

type MainClient struct {
	LocationsApiClient locationapisdk.LocationsAPI
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) LocationSdkClient {
	billingAPIconfiguration := locationapisdk.NewConfiguration()

	if customUrl != "" {
		billingAPIconfiguration.Servers = locationapisdk.ServerConfigurations{
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
	billingAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := locationapisdk.NewAPIClient(billingAPIconfiguration)

	return MainClient{
		LocationsApiClient: api_client.LocationsAPI,
	}
}

func (m MainClient) LocationsGet(location, productCategory string) ([]locationapisdk.Location, error) {
	request := m.LocationsApiClient.GetLocations(context.Background())

	if !client.IsZeroValue(location) {
		request = request.Location(locationapisdk.LocationEnum(location))
	}
	if !client.IsZeroValue(productCategory) {
		request = request.ProductCategory(locationapisdk.ProductCategoryEnum(productCategory))
	}

	return client.HandleResponse(request.Execute())
}
