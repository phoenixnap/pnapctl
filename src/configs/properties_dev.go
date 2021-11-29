//go:build dev
// +build dev

package configuration

const (
	// Hostname represents the URL entrypoint of our application
	BmcApiHostname  = "https://api-dev.phoenixnap.com/bmc/v1/"
	RancherHostname = "https://api-dev.phoenixnap.com/solutions/rancher/v1beta"
	TagsHostname    = "https://api-dev.phoenixnap.com/tag-manager/v1/"
	AuditHostname   = "https://api-dev.phoenixnap.com/audit/v1"
	// TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
	TokenURL = "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"
	// KnowledgeBaseURL represents the URL of the public knowledge base for pnapCTL
	KnowledgeBaseURL = "https://developers-dev.phoenixnap.com/cli"
	UserAgent        = "pnapctl/"
)
