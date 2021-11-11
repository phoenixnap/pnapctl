package ranchermodels

import (
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type Cluster struct {
	Id                    *string                `json:"id" yaml:"id"`
	Name                  *string                `json:"name" yaml:"name"`
	Description           *string                `json:"description" yaml:"description"`
	Location              string                 `json:"location" yaml:"location"`
	InitialClusterVersion *string                `json:"initialClusterVersion" yaml:"initialClusterVersion"`
	NodePools             *[]NodePool            `json:"nodePools" yaml:"nodePools"`
	Configuration         *RancherClusterConfig  `json:"configuration" yaml:"configuration"`
	Metadata              *RancherServerMetadata `json:"metadata" yaml:"metadata"`
	StatusDescription     *string                `json:"statusDescription" yaml:"statusDescription"`
}

func (c Cluster) ToSdk() ranchersdk.Cluster {
	var nodepools *[]ranchersdk.NodePool

	if nodepools != nil {
		for _, nodepool := range *c.NodePools {
			*nodepools = append(*nodepools, *nodepool.ToSdk())
		}
	}

	var configuration *ranchersdk.RancherClusterConfig
	var metadata *ranchersdk.RancherServerMetadata

	if c.Configuration != nil {
		configuration = c.Configuration.ToSdk()
	}

	if c.Metadata != nil {
		metadata = c.Metadata.ToSdk()
	}

	return ranchersdk.Cluster{
		Id:                    c.Id,
		Name:                  c.Name,
		Description:           c.Description,
		Location:              c.Location,
		InitialClusterVersion: c.InitialClusterVersion,
		NodePools:             nodepools,
		Configuration:         configuration,
		Metadata:              metadata,
		StatusDescription:     c.StatusDescription,
	}
}

func ClusterFromSdk(cluster ranchersdk.Cluster) Cluster {
	var nodepools *[]NodePool

	if nodepools != nil {
		for _, nodepool := range *cluster.NodePools {
			*nodepools = append(*nodepools, NodePoolFromSdk(nodepool))
		}
	}

	return Cluster{
		Id:                    cluster.Id,
		Name:                  cluster.Name,
		Description:           cluster.Description,
		Location:              cluster.Location,
		InitialClusterVersion: cluster.InitialClusterVersion,
		NodePools:             nodepools,
		Configuration:         RancherClusterConfigFromSdk(cluster.Configuration),
		Metadata:              RancherServerMetadataFromSdk(cluster.Metadata),
		StatusDescription:     cluster.StatusDescription,
	}
}

func CreateClusterFromFile(filename string, commandname string) (*ranchersdk.Cluster, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var cluster Cluster

	err = files.Unmarshal(data, &cluster, commandname)

	if err != nil {
		return nil, err
	}

	sdkCluster := cluster.ToSdk()

	return &sdkCluster, nil
}
