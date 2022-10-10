package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type Cluster struct {
	Id                    *string                       `json:"id" yaml:"id"`
	Name                  *string                       `json:"name" yaml:"name"`
	Description           *string                       `json:"description" yaml:"description"`
	Location              string                        `json:"location" yaml:"location"`
	InitialClusterVersion *string                       `json:"initialClusterVersion" yaml:"initialClusterVersion"`
	NodePools             []NodePool                    `json:"nodePools" yaml:"nodePools"`
	Configuration         *RancherClusterConfig         `json:"configuration" yaml:"configuration"`
	Metadata              *RancherServerMetadata        `json:"metadata" yaml:"metadata"`
	WorkloadConfiguration *ClusterWorkloadConfiguration `json:"workloadConfiguration" yaml:"workloadConfiguration"`
	StatusDescription     *string                       `json:"statusDescription" yaml:"statusDescription"`
}

func (cluster Cluster) ToSdk() ranchersdk.Cluster {
	var nodepools []ranchersdk.NodePool

	if cluster.NodePools != nil {
		nodepools = []ranchersdk.NodePool{}
		for _, nodepool := range cluster.NodePools {
			nodepools = append(nodepools, *nodepool.ToSdk())
		}
	}

	var configuration *ranchersdk.ClusterConfiguration
	var metadata *ranchersdk.ClusterMetadata
	var workloadConfiguration *ranchersdk.ClusterWorkloadConfiguration

	if cluster.Configuration != nil {
		configuration = cluster.Configuration.ToSdk()
	}

	if cluster.Metadata != nil {
		metadata = cluster.Metadata.ToSdk()
	}

	if cluster.WorkloadConfiguration != nil {
		workloadConfiguration = cluster.WorkloadConfiguration.ToSdk()
	}

	return ranchersdk.Cluster{
		Id:                    cluster.Id,
		Name:                  cluster.Name,
		Description:           cluster.Description,
		Location:              cluster.Location,
		InitialClusterVersion: cluster.InitialClusterVersion,
		NodePools:             nodepools,
		Configuration:         configuration,
		Metadata:              metadata,
		WorkloadConfiguration: workloadConfiguration,
		StatusDescription:     cluster.StatusDescription,
	}
}

func ClusterFromSdk(cluster ranchersdk.Cluster) Cluster {
	var nodepools []NodePool

	if cluster.NodePools != nil {
		nodepools = []NodePool{}
		for _, nodepool := range cluster.NodePools {
			nodepools = append(nodepools, NodePoolFromSdk(nodepool))
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
		WorkloadConfiguration: ClusterWorkloadConfigurationFromSdk(cluster.WorkloadConfiguration),
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
