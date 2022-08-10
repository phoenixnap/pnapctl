package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

type RancherClusterCertificates struct {
	CaCertificate  *string `json:"caCertificate" yaml:"caCertificate"`
	Certificate    *string `json:"certificate" yaml:"certificate"`
	CertificateKey *string `json:"certificateKey" yaml:"certificateKey"`
}

func (r RancherClusterCertificates) toSdk() *ranchersdk.RancherClusterConfigCertificates {
	return &ranchersdk.RancherClusterConfigCertificates{
		CaCertificate:  r.CaCertificate,
		Certificate:    r.Certificate,
		CertificateKey: r.CertificateKey,
	}
}

func RancherClusterCertificatesFromSdk(certificates *ranchersdk.RancherClusterConfigCertificates) *RancherClusterCertificates {
	if certificates == nil {
		return nil
	}

	return &RancherClusterCertificates{
		CaCertificate:  certificates.CaCertificate,
		Certificate:    certificates.Certificate,
		CertificateKey: certificates.CertificateKey,
	}
}
