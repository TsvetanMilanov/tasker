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
		UserInfoURL:  os.Getenv("AUTH0_USER_INFO_URL"),
	}
}

// GetAuth0ManagementConfig returns the auth0 management config stored
// in env vars.
func (c *Config) GetAuth0ManagementConfig() ctypes.Auth0MgmtConfig {
	return ctypes.Auth0MgmtConfig{
		Auth0Config:     c.GetAuth0Config(),
		MgmtAPIAudience: os.Getenv("AUTH0_MGMT_API_AUDIENCE"),
	}
}
