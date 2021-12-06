package sshkeymodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSshKeyUpdateToSdk(test_framework *testing.T) {
	cliModel := GenerateSshKeyUpdateCli()
	sshkeyupdate := cliModel.toSdk()

	assert.Equal(test_framework, cliModel.Default, sshkeyupdate.Default)
	assert.Equal(test_framework, cliModel.Name, sshkeyupdate.Name)
}
