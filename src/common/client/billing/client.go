package billing

import (
	"context"
	"net/http"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client BillingSdkClient

type BillingSdkClient interface {
	// Rated Usages
	RatedUsageGet(fromYearMonth, toYearMonth, productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error)
	RatedUsageMonthToDateGet(productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error)
	ProductsGet(productCode, productCategory, skuCode, location string) ([]billingapisdk.ProductsGet200ResponseInner, *http.Response, error)
	ReservationsGet(productCategory string) ([]billingapisdk.Reservation, *http.Response, error)
	ReservationsPost(request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, *http.Response, error)
	ReservationGetById(id string) (*billingapisdk.Reservation, *http.Response, error)
	ReservationDisableAutoRenew(id string, request billingapisdk.ReservationAutoRenewDisableRequest) (*billingapisdk.Reservation, *http.Response, error)
	ReservationEnableAutoRenew(id string) (*billingapisdk.Reservation, *http.Response, error)
	ReservationConvert(id string, request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, *http.Response, error)
	AccountBillingConfigurationGet() (*billingapisdk.ConfigurationDetails, *http.Response, error)
	ProductAvailabilityGet(productCategory []string, productCode []string, showOnlyMinQuantityAvailable bool, location []string, solution []string, minQuantity float32) ([]billingapisdk.ProductAvailability, *http.Response, error)
}

type MainClient struct {
	RatedUsageApiClient            billingapisdk.RatedUsageApi
	ProductsApiClient              billingapisdk.ProductsApi
	ReservationApiClient           billingapisdk.ReservationsApi
	BillingConfigurationsApiClient billingapisdk.BillingConfigurationsApi
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
		RatedUsageApiClient:            api_client.RatedUsageApi,
		ProductsApiClient:              api_client.ProductsApi,
		ReservationApiClient:           api_client.ReservationsApi,
		BillingConfigurationsApiClient: api_client.BillingConfigurationsApi,
	}
}

func (m MainClient) RatedUsageGet(fromYearMonth, toYearMonth, productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error) {
	request := m.RatedUsageApiClient.RatedUsageGet(context.Background())

	if !client.IsZero(fromYearMonth) {
		request.FromYearMonth(fromYearMonth)
	}
	if !client.IsZero(toYearMonth) {
		request.ToYearMonth(toYearMonth)
	}
	if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
		request.ProductCategory(*enum)
	}

	return request.Execute()
}

func (m MainClient) RatedUsageMonthToDateGet(productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, *http.Response, error) {
	request := m.RatedUsageApiClient.RatedUsageMonthToDateGet(context.Background())

	if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
		request.ProductCategory(*enum)
	}

	return request.Execute()
}

func (m MainClient) ProductsGet(productCode, productCategory, skuCode, location string) ([]billingapisdk.ProductsGet200ResponseInner, *http.Response, error) {
	request := m.ProductsApiClient.ProductsGet(context.Background())

	if !client.IsZero(productCode) {
		request.ProductCode(productCode)
	}
	if !client.IsZero(productCategory) {
		request.ProductCategory(productCategory)
	}
	if !client.IsZero(skuCode) {
		request.SkuCode(skuCode)
	}
	if !client.IsZero(location) {
		request.Location(location)
	}

	return request.Execute()
}

func (m MainClient) ReservationsGet(productCategory string) ([]billingapisdk.Reservation, *http.Response, error) {
	request := m.ReservationApiClient.ReservationsGet(context.Background())

	if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
		request.ProductCategory(*enum)
	}

	return request.Execute()
}

func (m MainClient) ReservationsPost(request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, *http.Response, error) {
	return m.ReservationApiClient.ReservationsPost(context.Background()).ReservationRequest(request).Execute()
}

func (m MainClient) ReservationGetById(id string) (*billingapisdk.Reservation, *http.Response, error) {
	return m.ReservationApiClient.ReservationsReservationIdGet(context.Background(), id).Execute()
}

func (m MainClient) ReservationDisableAutoRenew(id string, request billingapisdk.ReservationAutoRenewDisableRequest) (*billingapisdk.Reservation, *http.Response, error) {
	return m.ReservationApiClient.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), id).ReservationAutoRenewDisableRequest(request).Execute()
}

func (m MainClient) ReservationEnableAutoRenew(id string) (*billingapisdk.Reservation, *http.Response, error) {
	return m.ReservationApiClient.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), id).Execute()
}

func (m MainClient) ReservationConvert(id string, request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, *http.Response, error) {
	return m.ReservationApiClient.ReservationsReservationIdActionsConvertPost(context.Background(), id).ReservationRequest(request).Execute()
}

func (m MainClient) AccountBillingConfigurationGet() (*billingapisdk.ConfigurationDetails, *http.Response, error) {
	return m.BillingConfigurationsApiClient.AccountBillingConfigurationMeGet(context.Background()).Execute()
}

func (m MainClient) ProductAvailabilityGet(productCategory []string, productCode []string, showOnlyMinQuantityAvailable bool, location []string, solution []string, minQuantity float32) ([]billingapisdk.ProductAvailability, *http.Response, error) {
	request := m.ProductsApiClient.ProductAvailabilityGet(context.Background())

	if len(productCategory) != 0 {
		request.ProductCategory(productCategory)
	}
	if len(productCode) != 0 {
		request.ProductCategory(productCategory)
	}
	if len(solution) != 0 {
		request.Solution(solution)
	}
	if !client.IsZero(minQuantity) {
		request.MinQuantity(minQuantity)
	}

	locations := iterutils.Deref(iterutils.Map(location, func(str string) *billingapisdk.LocationEnum {
		enum, _ := billingapisdk.NewLocationEnumFromValue(str)
		return enum
	}))

	if len(locations) != 0 {
		request.Location(locations)
	}
	request.ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable)

	return request.Execute()
}
