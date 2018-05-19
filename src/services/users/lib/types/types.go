package types

import (
	"time"

	"github.com/TsvetanMilanov/tasker-common/common/cdeclarations"
)

// BaseHandler ...
type BaseHandler struct {
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}

// TokenResponse ...
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// Auth0MgmtUserInfoResponse ...
type Auth0MgmtUserInfoResponse struct {
	Email         string                 `json:"email"`
	Username      string                 `json:"username"`
	EmailVerified bool                   `json:"email_verified"`
	UserID        string                 `json:"user_id"`
	CreatedAt     time.Time              `json:"created_at"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	Picture       string                 `json:"picture"`
	Nickname      string                 `json:"nickname"`
	Blocket       bool                   `json:"blocked"`
}
