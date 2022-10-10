package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestRancherClusterCertificateToSdk(test_framework *testing.T) {
	rancherRancherClusterCertificates := GenerateRancherClusterCertificatesCli()
	sdkRancherClusterCertificates := *rancherRancherClusterCertificates.toSdk()

	assertEqualRancherClusterCertificates(test_framework, rancherRancherClusterCertificates, sdkRancherClusterCertificates)
}

func TestRancherClusterCertificateFromSdk(test_framework *testing.T) {
	sdkRancherClusterCertificates := GenerateRancherClusterCertificatesSdk()
	rancherRancherClusterCertificates := *RancherClusterCertificatesFromSdk(&sdkRancherClusterCertificates)

	assertEqualRancherClusterCertificates(test_framework, rancherRancherClusterCertificates, sdkRancherClusterCertificates)
}

func assertEqualRancherClusterCertificates(test_framework *testing.T, cliRancherClusterCertificates RancherClusterCertificates, sdkRancherClusterCertificates ranchersdk.RancherClusterConfigCertificates) {
	assert.Equal(test_framework, cliRancherClusterCertificates.CaCertificate, sdkRancherClusterCertificates.CaCertificate)
	assert.Equal(test_framework, cliRancherClusterCertificates.Certificate, sdkRancherClusterCertificates.Certificate)
	assert.Equal(test_framework, cliRancherClusterCertificates.CertificateKey, sdkRancherClusterCertificates.CertificateKey)
}
