package sshkeymodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSshKeyCreateToSdk(test_framework *testing.T) {
	cliModel := GenerateSshKeyCreateCli()
	sshkeycreate := cliModel.toSdk()

	assert.Equal(test_framework, cliModel.Default, sshkeycreate.Default)
	assert.Equal(test_framework, cliModel.Name, sshkeycreate.Name)
	assert.Equal(test_framework, cliModel.Key, sshkeycreate.Key)
}
