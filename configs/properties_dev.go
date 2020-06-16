// +build dev

package configuration

const (
	// Hostname represents the URL entrypoint of our application
	Hostname = "https://api-dev.phoenixnap.com/bmc/v1beta/"
	// TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
	TokenURL = "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"
	// KnowledgeBaseURL represents the URL of the public knowledge base for pnapCTL
	KnowledgeBaseURL = "https://developers-dev.phoenixnap.com/cli"
)
