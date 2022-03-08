package tables

import (
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"testing"
)

func TestToIpBlockTable(test_framework *testing.T) {
	ipBlock := ipmodels.GenerateIpBlockSdk()
	table := ToIpBlockTable(ipBlock)

	assert.Equal(test_framework, ipBlock.Id, table.Id)
	assert.Equal(test_framework, ipBlock.Location, table.Location)
	assert.Equal(test_framework, ipBlock.CidrBlockSize, table.CidrBlockSize)
	assert.Equal(test_framework, ipBlock.Cidr, table.Cidr)
	assert.Equal(test_framework, ipBlock.Status, table.Status)
	assert.Equal(test_framework, ipBlock.AssignedResourceId, table.AssignedResourceId)
	assert.Equal(test_framework, ipBlock.AssignedResourceType, table.AssignedResourceType)
}
