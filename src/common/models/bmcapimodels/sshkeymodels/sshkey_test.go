package sshkeymodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSshKeySdkToDto(test_framework *testing.T) {
	sdkModel := GenerateSshKeySdk()
	sshkey := SshKeySdkToDto(sdkModel)

	assert.Equal(test_framework, sdkModel.Id, sshkey.Id)
	assert.Equal(test_framework, sdkModel.Default, sshkey.Default)
	assert.Equal(test_framework, sdkModel.Name, sshkey.Name)
	assert.Equal(test_framework, sdkModel.Key, sshkey.Key)
	assert.Equal(test_framework, sdkModel.Fingerprint, sshkey.Fingerprint)
	assert.Equal(test_framework, sdkModel.CreatedOn, sshkey.CreatedOn)
	assert.Equal(test_framework, sdkModel.LastUpdatedOn, sshkey.LastUpdatedOn)
}
