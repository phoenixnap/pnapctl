package rancher

import (
	"context"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client RancherSdkClient

type RancherSdkClient interface {
	ClusterPost(clusterCreate ranchersdk.Cluster) (*ranchersdk.Cluster, error)
	ClustersGet() ([]ranchersdk.Cluster, error)
	ClusterGetById(clusterId string) (*ranchersdk.Cluster, error)
	ClusterDelete(clusterId string) (*ranchersdk.DeleteResult, error)
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

func (m MainClient) ClusterPost(cluster ranchersdk.Cluster) (*ranchersdk.Cluster, error) {
	return client.HandleResponse(m.RancherSdkClient.ClustersPost(context.Background()).Cluster(cluster).Execute())
}

func (m MainClient) ClustersGet() ([]ranchersdk.Cluster, error) {
	return client.HandleResponse(m.RancherSdkClient.ClustersGet(context.Background()).Execute())
}

func (m MainClient) ClusterGetById(clusterId string) (*ranchersdk.Cluster, error) {
	return client.HandleResponse(m.RancherSdkClient.ClustersIdGet(context.Background(), clusterId).Execute())
}

func (m MainClient) ClusterDelete(clusterId string) (*ranchersdk.DeleteResult, error) {
	return client.HandleResponse(m.RancherSdkClient.ClustersIdDelete(context.Background(), clusterId).Execute())
}
