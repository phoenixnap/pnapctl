package billing

import (
	"context"
	"fmt"

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
	RatedUsageGet(fromYearMonth, toYearMonth, productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, error)
	RatedUsageMonthToDateGet(productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, error)
	ProductsGet(productCode, productCategory, skuCode, location string) ([]billingapisdk.ProductsGet200ResponseInner, error)
	ReservationsGet(productCategory string) ([]billingapisdk.Reservation, error)
	ReservationsPost(request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, error)
	ReservationGetById(id string) (*billingapisdk.Reservation, error)
	ReservationDisableAutoRenew(id string, request billingapisdk.ReservationAutoRenewDisableRequest) (*billingapisdk.Reservation, error)
	ReservationEnableAutoRenew(id string) (*billingapisdk.Reservation, error)
	ReservationConvert(id string, request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, error)
	AccountBillingConfigurationGet() (*billingapisdk.ConfigurationDetails, error)
	ProductAvailabilityGet(productCategory []string, productCode []string, showOnlyMinQuantityAvailable bool, location []string, solution []string, minQuantity float32) ([]billingapisdk.ProductAvailability, error)
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

func (m MainClient) RatedUsageGet(fromYearMonth, toYearMonth, productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, error) {
	request := m.RatedUsageApiClient.RatedUsageGet(context.Background())

	if !client.IsZero(fromYearMonth) {
		request = request.FromYearMonth(fromYearMonth)
	}
	if !client.IsZero(toYearMonth) {
		request = request.ToYearMonth(toYearMonth)
	}
	if !client.IsZero(productCategory) {
		if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
			request = request.ProductCategory(*enum)
		} else {
			fmt.Printf("Product category passed (%s) isn't valid %v. Ignoring...\n", productCategory, billingapisdk.AllowedProductCategoryEnumEnumValues)
		}
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) RatedUsageMonthToDateGet(productCategory string) ([]billingapisdk.RatedUsageGet200ResponseInner, error) {
	request := m.RatedUsageApiClient.RatedUsageMonthToDateGet(context.Background())

	if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
		request = request.ProductCategory(*enum)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) ProductsGet(productCode, productCategory, skuCode, location string) ([]billingapisdk.ProductsGet200ResponseInner, error) {
	request := m.ProductsApiClient.ProductsGet(context.Background())

	if !client.IsZero(productCode) {
		request = request.ProductCode(productCode)
	}
	if !client.IsZero(productCategory) {
		request = request.ProductCategory(productCategory)
	}
	if !client.IsZero(skuCode) {
		request = request.SkuCode(skuCode)
	}
	if !client.IsZero(location) {
		request = request.Location(location)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) ReservationsGet(productCategory string) ([]billingapisdk.Reservation, error) {
	request := m.ReservationApiClient.ReservationsGet(context.Background())

	if !client.IsZero(productCategory) {
		if enum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory); err == nil && enum != nil {
			request = request.ProductCategory(*enum)
		} else {
			fmt.Printf("Product category passed (%s) isn't valid %v. Ignoring...\n", productCategory, billingapisdk.AllowedProductCategoryEnumEnumValues)
		}
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) ReservationsPost(request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, error) {
	return client.HandleResponse(m.ReservationApiClient.ReservationsPost(context.Background()).ReservationRequest(request).Execute())
}

func (m MainClient) ReservationGetById(id string) (*billingapisdk.Reservation, error) {
	return client.HandleResponse(m.ReservationApiClient.ReservationsReservationIdGet(context.Background(), id).Execute())
}

func (m MainClient) ReservationDisableAutoRenew(id string, request billingapisdk.ReservationAutoRenewDisableRequest) (*billingapisdk.Reservation, error) {
	return client.HandleResponse(m.ReservationApiClient.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), id).ReservationAutoRenewDisableRequest(request).Execute())
}

func (m MainClient) ReservationEnableAutoRenew(id string) (*billingapisdk.Reservation, error) {
	return client.HandleResponse(m.ReservationApiClient.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), id).Execute())
}

func (m MainClient) ReservationConvert(id string, request billingapisdk.ReservationRequest) (*billingapisdk.Reservation, error) {
	return client.HandleResponse(m.ReservationApiClient.ReservationsReservationIdActionsConvertPost(context.Background(), id).ReservationRequest(request).Execute())
}

func (m MainClient) AccountBillingConfigurationGet() (*billingapisdk.ConfigurationDetails, error) {
	return client.HandleResponse(m.BillingConfigurationsApiClient.AccountBillingConfigurationMeGet(context.Background()).Execute())
}

func (m MainClient) ProductAvailabilityGet(productCategory []string, productCode []string, showOnlyMinQuantityAvailable bool, location []string, solution []string, minQuantity float32) ([]billingapisdk.ProductAvailability, error) {
	request := m.ProductsApiClient.ProductAvailabilityGet(context.Background())

	if len(productCategory) != 0 {
		request = request.ProductCategory(productCategory)
	}
	if len(productCode) != 0 {
		request = request.ProductCode(productCode)
	}
	if len(solution) != 0 {
		request = request.Solution(solution)
	}
	if !client.IsZero(minQuantity) {
		request = request.MinQuantity(minQuantity)
	}

	locations := iterutils.Deref(iterutils.Map(location, func(str string) *billingapisdk.LocationEnum {
		enum, err := billingapisdk.NewLocationEnumFromValue(str)
		if err != nil {
			fmt.Printf("Location passed (%s) isn't valid %v. Ignoring...\n", str, billingapisdk.AllowedLocationEnumEnumValues)
		}
		return enum
	}))

	if len(locations) != 0 {
		request = request.Location(locations)
	}
	request = request.ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable)

	return client.HandleResponse(request.Execute())
}
