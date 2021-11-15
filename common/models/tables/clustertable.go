package tables

import (
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	"phoenixnap.com/pnap-cli/common/models/ranchermodels"
)

type ClusterTable struct {
	Id                    string   `header:"ID"`
	Name                  string   `header:"Name"`
	Description           string   `header:"Description"`
	Location              string   `header:"Location"`
	InitialClusterVersion string   `header:"Initial cluster version"`
	NodePools             []string `header:"Node pools"`
	Configuration         string   `header:"Configuration"`
	Metadata              string   `header:"Metadata"`
	StatusDescription     string   `header:"Status description"`
}

func ClusterFromSdk(cluster ranchersdk.Cluster) ClusterTable {
	return ClusterTable{
		Id:                    DerefString(cluster.Id),
		Name:                  DerefString(cluster.Name),
		Description:           DerefString(cluster.Description),
		Location:              cluster.Location,
		InitialClusterVersion: DerefString(cluster.InitialClusterVersion),
		NodePools:             ranchermodels.NodePoolsToTableStrings(cluster.NodePools),
		Configuration:         ranchermodels.RancherClusterConfigToTableString(cluster.Configuration),
		Metadata:              ranchermodels.RancherServerMetadataToTableString(cluster.Metadata),
		StatusDescription:     DerefString(cluster.StatusDescription),
	}
}
