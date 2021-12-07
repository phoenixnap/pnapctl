package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
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

func assertEqualRancherClusterCertificates(test_framework *testing.T, r1 RancherClusterCertificates, r2 ranchersdk.RancherClusterCertificates) {
	assert.Equal(test_framework, r1.CaCertificate, r2.CaCertificate)
	assert.Equal(test_framework, r1.Certificate, r2.Certificate)
	assert.Equal(test_framework, r1.CertificateKey, r2.CertificateKey)
}
