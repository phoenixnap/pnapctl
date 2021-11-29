//go:build !dev
// +build !dev

package configuration

const (
	// Hostname represents the URL entrypoint of our application. By default SDK points to prod
	BmcApiHostname  = ""
	RancherHostname = ""
	TagsHostname    = ""
	AuditHostname   = ""
	// TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
	TokenURL = "https://auth.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"
	// KnowledgeBaseURL represents the URL of the public knowledge base for pnapCTL
	KnowledgeBaseURL = "https://developers.phoenixnap.com/cli"
	UserAgent        = "pnapctl/"
)
