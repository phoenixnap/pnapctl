package rancher

import (
	"context"
	"net/http"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client RancherSdkClient

type RancherSdkClient interface {
	ClusterPost(clusterCreate ranchersdk.Cluster) (*ranchersdk.Cluster, *http.Response, error)
	ClustersGet() ([]ranchersdk.Cluster, *http.Response, error)
	ClusterGetById(clusterId string) (*ranchersdk.Cluster, *http.Response, error)
	ClusterDelete(clusterId string) (*ranchersdk.DeleteResult, *http.Response, error)
}

type MainClient struct {
	RancherSdkClient ranchersdk.ClustersApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) RancherSdkClient {
	rancherConfiguration := ranchersdk.NewConfiguration()

	if customUrl != "" {
		rancherConfiguration.Servers = ranchersdk.ServerConfigurations{
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

	rancherConfiguration.HTTPClient = config.Client(context.Background())
	rancherConfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := ranchersdk.NewAPIClient(rancherConfiguration)

	return MainClient{
		RancherSdkClient: api_client.ClustersApi,
	}
}

func (m MainClient) ClusterPost(cluster ranchersdk.Cluster) (*ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersPost(context.Background()).Cluster(cluster).Execute()
}

func (m MainClient) ClustersGet() ([]ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersGet(context.Background()).Execute()
}

func (m MainClient) ClusterGetById(clusterId string) (*ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdGet(context.Background(), clusterId).Execute()
}

func (m MainClient) ClusterDelete(clusterId string) (*ranchersdk.DeleteResult, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdDelete(context.Background(), clusterId).Execute()
}
