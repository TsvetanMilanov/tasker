package common

import (
	"os"

	"github.com/TsvetanMilanov/tasker/src/common/ctypes"
)

// Config handles common server configurations.
type Config struct {
}

// GetAuth0Config returns the auth0 config stored in env vars.
func (c *Config) GetAuth0Config() ctypes.Auth0Config {
	return ctypes.Auth0Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		TokenURL:     os.Getenv("AUTH0_TOKEN_URL"),
	}
}

// GetAuth0ManagementConfig returns the auth0 management config stored
// in env vars.
func (c *Config) GetAuth0ManagementConfig() ctypes.Auth0Config {
	return ctypes.Auth0Config{
		ClientID:     os.Getenv("AUTH0_MGMT_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_MGMT_CLIENT_SECRET"),
		TokenURL:     os.Getenv("AUTH0_TOKEN_URL"),
	}
}
