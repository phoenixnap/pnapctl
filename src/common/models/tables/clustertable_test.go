package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestClusterFromSdk(test_framework *testing.T) {
	cluster := generators.GenerateClusterSdk()
	table := ClusterFromSdk(cluster)

	assertClustersEqual(test_framework, cluster, table)
}

func assertClustersEqual(test_framework *testing.T, cluster ranchersdk.Cluster, table ClusterTable) {
	assert.Equal(test_framework, DerefString(cluster.Id), table.Id)
	assert.Equal(test_framework, DerefString(cluster.Name), table.Name)
	assert.Equal(test_framework, DerefString(cluster.Description), table.Description)
	assert.Equal(test_framework, cluster.Location, table.Location)
	assert.Equal(test_framework, DerefString(cluster.InitialClusterVersion), table.InitialClusterVersion)
	assert.Equal(test_framework, models.NodePoolsToTableStrings(cluster.NodePools), table.NodePools)
	assert.Equal(test_framework, models.ClusterConfigurationToTableString(cluster.Configuration), table.Configuration)
	assert.Equal(test_framework, models.ClusterMetadataToTableString(cluster.Metadata), table.Metadata)
	assert.Equal(test_framework, DerefString(cluster.StatusDescription), table.StatusDescription)
}
