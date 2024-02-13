package invoicing

import (
	"context"
	"os"

	invoicingapisdk "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client InvoicingSdkClient

type InvoicingSdkClient interface {
	InvoicesGet(number string, status string, sentOnFrom string, sentOnTo string, limit int, offset int, sortField string, sortDirection string) (*invoicingapisdk.PaginatedInvoices, error)
	InvoicesInvoiceIdGet(invoiceId string) (*invoicingapisdk.Invoice, error)
	InvoicesInvoiceIdGeneratePdfPost(invoiceId string) (*os.File, error)
	InvoicesInvoiceIdPayPost(invoiceId string) (map[string]interface{}, error)
}

type MainClient struct {
	InvoicingApiClient invoicingapisdk.InvoicesAPI
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) InvoicingSdkClient {
	invoicingAPIconfiguration := invoicingapisdk.NewConfiguration()

	if customUrl != "" {
		invoicingAPIconfiguration.Servers = invoicingapisdk.ServerConfigurations{
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
		Scopes:       []string{"invoices", "invoices.read"},
	}

	invoicingAPIconfiguration.HTTPClient = config.Client(context.Background())
	invoicingAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version
	invoicingAPIconfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := invoicingapisdk.NewAPIClient(invoicingAPIconfiguration)

	return MainClient{
		InvoicingApiClient: api_client.InvoicesAPI,
	}
}

func (m MainClient) InvoicesGet(number string, status string, sentOnFrom string, sentOnTo string, limit int, offset int, sortField string, sortDirection string) (*invoicingapisdk.PaginatedInvoices, error) {
	request := m.InvoicingApiClient.InvoicesGet(context.Background())

	if !client.IsZeroValue(number) {
		request = request.Number(number)
	}
	if !client.IsZeroValue(status) {
		request = request.Status(status)
	}
	if date := client.ParseDate(sentOnFrom); date != nil {
		request = request.SentOnFrom(*date)
	}
	if date := client.ParseDate(sentOnTo); date != nil {
		request = request.SentOnTo(*date)
	}
	if !client.IsZeroValue(limit) {
		request = request.Limit(int32(limit))
	}
	if !client.IsZeroValue(offset) {
		request = request.Offset(int32(offset))
	}
	if !client.IsZeroValue(sortField) {
		request = request.SortField(sortField)
	}
	if !client.IsZeroValue(sortDirection) {
		request = request.SortDirection(sortDirection)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) InvoicesInvoiceIdGet(invoiceId string) (*invoicingapisdk.Invoice, error) {
	return client.HandleResponse(m.InvoicingApiClient.InvoicesInvoiceIdGet(context.Background(), invoiceId).Execute())
}

func (m MainClient) InvoicesInvoiceIdGeneratePdfPost(invoiceId string) (*os.File, error) {
	return client.HandleResponse(m.InvoicingApiClient.InvoicesInvoiceIdGeneratePdfPost(context.Background(), invoiceId).Execute())
}

func (m MainClient) InvoicesInvoiceIdPayPost(invoiceId string) (map[string]interface{}, error) {
	return client.HandleResponse(m.InvoicingApiClient.InvoicesInvoiceIdPayPost(context.Background(), invoiceId).Execute())
}