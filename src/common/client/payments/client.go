package bmcapi

import (
	"context"

	paymentsapisdk "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client PaymentApiSdkClient

type PaymentApiSdkClient interface {
	//Transactions
	TransactionsGet() (*paymentsapisdk.PaginatedTransactions, error)
	TransactionGetById(quotaId string) (*paymentsapisdk.Transaction, error)
}

type MainClient struct {
	TransactionApiClient paymentsapisdk.TransactionsAPI
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) PaymentApiSdkClient {
	paymentsAPIconfiguration := paymentsapisdk.NewConfiguration()

	if customUrl != "" {
		paymentsAPIconfiguration.Servers = paymentsapisdk.ServerConfigurations{
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
		Scopes:       []string{"transactions", "transactions.read"},
	}

	paymentsAPIconfiguration.HTTPClient = config.Client(context.Background())
	paymentsAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version
	paymentsAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := paymentsapisdk.NewAPIClient(paymentsAPIconfiguration)

	return MainClient{
		TransactionApiClient: api_client.TransactionsAPI,
	}
}

// Transaction APIs
func (m MainClient) TransactionsGet() (*paymentsapisdk.PaginatedTransactions, error) {
	return client.HandleResponse(m.TransactionApiClient.TransactionsGet(context.Background()).Execute())
}

func (m MainClient) TransactionGetById(transactionId string) (*paymentsapisdk.Transaction, error) {
	return client.HandleResponse(m.TransactionApiClient.TransactionIdGet(context.Background(), transactionId).Execute())
}
