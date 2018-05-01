package ctypes

// Auth0Config The configuration of the auth0 client.
type Auth0Config struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	UserInfoURL  string
}

// Auth0MgmtConfig The configuration of the auth0 management client.
type Auth0MgmtConfig struct {
	Auth0Config
	MgmtAPIAudience string
}
