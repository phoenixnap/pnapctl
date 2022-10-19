package networkstorage

import (
	"context"

	networkstoragesdk "github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client NetworkStorageSdkClient

type NetworkStorageSdkClient interface {
	NetworkStorageGet() ([]networkstoragesdk.StorageNetwork, error)
	NetworkStoragePost(storageCreate networkstoragesdk.StorageNetworkCreate) (*networkstoragesdk.StorageNetwork, error)
	NetworkStorageGetById(storageId string) (*networkstoragesdk.StorageNetwork, error)
	NetworkStoragePatch(storageId string, storageUpdate networkstoragesdk.StorageNetworkUpdate) (*networkstoragesdk.StorageNetwork, error)
	NetworkStorageDelete(storageId string) error
	NetworkStorageGetVolumes(storageId string) ([]networkstoragesdk.Volume, error)
	NetworkStorageGetVolumeById(storageId string, volumeId string) (*networkstoragesdk.Volume, error)
}

type MainClient struct {
	StorageNetworksApiClient networkstoragesdk.StorageNetworksApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) MainClient {
	networkstorageAPIconfiguration := networkstoragesdk.NewConfiguration()

	if customUrl != "" {
		networkstorageAPIconfiguration.Servers = networkstoragesdk.ServerConfigurations{
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

	networkstorageAPIconfiguration.HTTPClient = config.Client(context.Background())
	networkstorageAPIconfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := networkstoragesdk.NewAPIClient(networkstorageAPIconfiguration)

	return MainClient{
		StorageNetworksApiClient: api_client.StorageNetworksApi,
	}
}

func (m MainClient) NetworkStorageGet() ([]networkstoragesdk.StorageNetwork, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksGet(context.Background()).Execute())
}

func (m MainClient) NetworkStoragePost(storageCreate networkstoragesdk.StorageNetworkCreate) (*networkstoragesdk.StorageNetwork, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksPost(context.Background()).StorageNetworkCreate(storageCreate).Execute())
}

func (m MainClient) NetworkStorageGetById(storageId string) (*networkstoragesdk.StorageNetwork, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksIdGet(context.Background(), storageId).Execute())
}

func (m MainClient) NetworkStoragePatch(storageId string, storageUpdate networkstoragesdk.StorageNetworkUpdate) (*networkstoragesdk.StorageNetwork, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksIdPatch(context.Background(), storageId).StorageNetworkUpdate(storageUpdate).Execute())
}

func (m MainClient) NetworkStorageDelete(storageId string) error {
	return client.HandleResponseWithoutBody(m.StorageNetworksApiClient.StorageNetworksIdDelete(context.Background(), storageId).Execute())
}

func (m MainClient) NetworkStorageGetVolumes(storageId string) ([]networkstoragesdk.Volume, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksStorageNetworkIdVolumesGet(context.Background(), storageId).Execute())
}

func (m MainClient) NetworkStorageGetVolumeById(storageId string, volumeId string) (*networkstoragesdk.Volume, error) {
	return client.HandleResponse(m.StorageNetworksApiClient.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(context.Background(), storageId, volumeId).Execute())
}
