package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
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
	assert.Equal(test_framework, iterutils.MapRef(cluster.NodePools, models.NodePoolToTableString), table.NodePools)
	assert.Equal(test_framework, models.ClusterConfigurationToTableString(cluster.Configuration), table.Configuration)
	assert.Equal(test_framework, models.ClusterMetadataToTableString(cluster.Metadata), table.Metadata)
	assert.Equal(test_framework, DerefString(cluster.StatusDescription), table.StatusDescription)
}
