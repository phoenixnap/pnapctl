package client

import (
	"context"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"golang.org/x/oauth2/clientcredentials"
)

var BmcApiClient bmcapi.DefaultApi

func SetupBmcApiClient(clientId string, clientSecret string) {
	configuration := bmcapi.NewConfiguration()

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token",
		Scopes:       []string{"bmc", "bmc.read"},
	}

	configuration.HTTPClient = config.Client(context.Background())

	api_client := bmcapi.NewAPIClient(configuration)
	BmcApiClient = api_client.DefaultApi
}
