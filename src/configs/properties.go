package configuration

const (
	// TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
	TokenURL = "https://auth.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"
	// KnowledgeBaseURL represents the URL of the public knowledge base for pnapCTL
	KnowledgeBaseURL = "https://developers.phoenixnap.com/cli"
	UserAgent        = "pnapctl/"
	XPoweredBy       = "pnapctl/"
)
