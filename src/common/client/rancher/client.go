package rancher

import (
	"context"
	"net/http"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnap-cli/commands/version"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client RancherSdkClient

type RancherSdkClient interface {
	ClusterPost(clusterCreate ranchersdk.Cluster) (ranchersdk.Cluster, *http.Response, error)
	ClustersGet() ([]ranchersdk.Cluster, *http.Response, error)
	ClusterGetById(clusterId string) (ranchersdk.Cluster, *http.Response, error)
	ClusterDelete(clusterId string) (ranchersdk.DeleteResult, *http.Response, error)
}

type MainClient struct {
	RancherSdkClient ranchersdk.ClustersApi
}

func NewMainClient(clientId string, clientSecret string) RancherSdkClient {
	rancherConfiguration := ranchersdk.NewConfiguration()

	if configuration.RancherHostname != "" {
		rancherConfiguration.Servers[0].URL = configuration.RancherHostname
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	rancherConfiguration.HTTPClient = config.Client(context.Background())
	rancherConfiguration.UserAgent = configuration.UserAgentPrefix + version.AppVersion.Version

	api_client := ranchersdk.NewAPIClient(rancherConfiguration)

	return MainClient{
		RancherSdkClient: api_client.ClustersApi,
	}
}

func (m MainClient) ClusterPost(cluster ranchersdk.Cluster) (ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersPost(context.Background()).Cluster(cluster).Execute()
}

func (m MainClient) ClustersGet() ([]ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersGet(context.Background()).Execute()
}

func (m MainClient) ClusterGetById(clusterId string) (ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdGet(context.Background(), clusterId).Execute()
}

func (m MainClient) ClusterDelete(clusterId string) (ranchersdk.DeleteResult, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdDelete(context.Background(), clusterId).Execute()
}
