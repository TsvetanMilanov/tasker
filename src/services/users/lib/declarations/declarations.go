package declarations

import "github.com/TsvetanMilanov/tasker/src/services/users/lib/types"

// IUsersService describes methods for working with auth0 users.
type IUsersService interface {
	GetUserInfoFromToken(authorizationHeader string) (*types.Auth0MgmtUserInfoResponse, error)
}

// IAuthService describes methods for working with the auth0 authentication service.
type IAuthService interface {
	GetAppToAppToken() (*types.TokenResponse, error)
}
