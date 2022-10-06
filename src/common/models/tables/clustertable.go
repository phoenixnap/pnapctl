package tables

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/common/models/ranchermodels"
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
		Configuration:         ranchermodels.ClusterConfigurationToTableString(cluster.Configuration),
		Metadata:              ranchermodels.ClusterMetadataToTableString(cluster.Metadata),
		StatusDescription:     DerefString(cluster.StatusDescription),
	}
}
