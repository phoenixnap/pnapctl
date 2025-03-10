package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestToIpBlockTable(test_framework *testing.T) {
	ipBlock := generators.Generate[ipapi.IpBlock]()
	table := ToIpBlockTable(ipBlock)

	assert.Equal(test_framework, DerefString(ipBlock.Id), table.Id)
	assert.Equal(test_framework, DerefString(ipBlock.Location), table.Location)
	assert.Equal(test_framework, DerefString(ipBlock.CidrBlockSize), table.CidrBlockSize)
	assert.Equal(test_framework, DerefString(ipBlock.Cidr), table.Cidr)
	assert.Equal(test_framework, DerefString(ipBlock.Status), table.Status)
	assert.Equal(test_framework, DerefString(ipBlock.AssignedResourceId), table.AssignedResourceId)
	assert.Equal(test_framework, DerefString(ipBlock.AssignedResourceType), table.AssignedResourceType)
}
