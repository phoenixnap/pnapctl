package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestToSshKeyTable(test_framework *testing.T) {
	sshkey := generators.Generate[bmcapisdk.SshKey]()
	table := ToSshKeyTable(sshkey)

	assertSshKeysEqual(test_framework, sshkey, table)
}

func TestToSshKeyTableFull(test_framework *testing.T) {
	sshkey := generators.Generate[bmcapisdk.SshKey]()
	table := ToSshKeyTableFull(sshkey)

	assertSshKeysFullEqual(test_framework, sshkey, table)
}

func assertSshKeysEqual(test_framework *testing.T, sshKey bmcapisdk.SshKey, table SshKeyTable) {
	assert.Equal(test_framework, sshKey.Id, table.Id)
	assert.Equal(test_framework, sshKey.Default, table.Default)
	assert.Equal(test_framework, sshKey.Name, table.Name)
	assert.Equal(test_framework, sshKey.Fingerprint, table.Fingerprint)
	assert.Equal(test_framework, sshKey.CreatedOn.String(), table.CreatedOn)
	assert.Equal(test_framework, sshKey.LastUpdatedOn.String(), table.LastUpdatedOn)
}

func assertSshKeysFullEqual(test_framework *testing.T, sshKey bmcapisdk.SshKey, table SshKeyTableFull) {
	assert.Equal(test_framework, sshKey.Id, table.Id)
	assert.Equal(test_framework, sshKey.Default, table.Default)
	assert.Equal(test_framework, sshKey.Name, table.Name)
	assert.Equal(test_framework, sshKey.Key, table.Key)
	assert.Equal(test_framework, sshKey.Fingerprint, table.Fingerprint)
	assert.Equal(test_framework, sshKey.CreatedOn.String(), table.CreatedOn)
	assert.Equal(test_framework, sshKey.LastUpdatedOn.String(), table.LastUpdatedOn)
}
