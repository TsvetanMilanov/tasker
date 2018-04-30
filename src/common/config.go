package common

import (
	"os"

	"github.com/TsvetanMilanov/tasker/src/common/types"
)

// Config handles common server configurations.
type Config struct {
}

// GetAuth0Config returns the auth0 config stored in env vars.
func (c *Config) GetAuth0Config() types.Auth0Config {
	return types.Auth0Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		TokenURL:     os.Getenv("AUTH0_TOKEN_URL"),
	}
}
