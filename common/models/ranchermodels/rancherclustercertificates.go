package ranchermodels

import (
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
)

type RancherClusterCertificates struct {
	CaCertificate  *string `json:"caCertificate" yaml:"caCertificate"`
	Certificate    *string `json:"certificate" yaml:"certificate"`
	CertificateKey *string `json:"certificateKey" yaml:"certificateKey"`
}

func (r RancherClusterCertificates) toSdk() *ranchersdk.RancherClusterCertificates {
	return &ranchersdk.RancherClusterCertificates{
		CaCertificate:  r.CaCertificate,
		Certificate:    r.Certificate,
		CertificateKey: r.CertificateKey,
	}
}

func RancherClusterCertificatesFromSdk(certificates *ranchersdk.RancherClusterCertificates) *RancherClusterCertificates {
	if certificates == nil {
		return nil
	}

	return &RancherClusterCertificates{
		CaCertificate:  certificates.CaCertificate,
		Certificate:    certificates.Certificate,
		CertificateKey: certificates.CertificateKey,
	}
}
