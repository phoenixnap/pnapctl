package payments

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
	TransactionsGet(limit int, offset int, sortDirection string, sortField string, from string, to string) (*paymentsapisdk.PaginatedTransactions, error)
	TransactionGetById(transactionId string) (*paymentsapisdk.Transaction, error)
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
func (m MainClient) TransactionsGet(limit int, offset int, sortDirection string, sortField string, from string, to string) (*paymentsapisdk.PaginatedTransactions, error) {
	request := m.TransactionApiClient.TransactionsGet(context.Background())

	if !client.IsZeroValue(limit) {
		request = request.Limit(int32(limit))
	}
	if !client.IsZeroValue(offset) {
		request = request.Offset(int32(offset))
	}
	if !client.IsZeroValue(sortDirection) {
		request = request.SortDirection(sortDirection)
	}
	if !client.IsZeroValue(sortField) {
		request = request.SortField(sortField)
	}
	if date := client.ParseDate(from); date != nil {
		request = request.From(*date)
	}
	if date := client.ParseDate(to); date != nil {
		request = request.To(*date)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) TransactionGetById(transactionId string) (*paymentsapisdk.Transaction, error) {
	return client.HandleResponse(m.TransactionApiClient.TransactionIdGet(context.Background(), transactionId).Execute())
}
