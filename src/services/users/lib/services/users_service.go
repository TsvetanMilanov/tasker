package services

import (
	"fmt"

	"github.com/TsvetanMilanov/tasker-common/common/cdeclarations"
	"github.com/TsvetanMilanov/tasker-common/common/ctypes"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/declarations"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
)

// UsersService implements IUsersService.
type UsersService struct {
	Auth       declarations.IAuthService `di:""`
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}

// GetUserInfoFromToken returns the full user info from the provided
// Authorization header value.
func (s *UsersService) GetUserInfoFromToken(authorizationHeader string) (*types.Auth0MgmtUserInfoResponse, error) {
	userInfo, err := s.GetBasicUserInfoFromToken(authorizationHeader)
	if err != nil {
		return nil, err
	}
	auth0MgmtCfg := s.Config.GetAuth0ManagementConfig()
	mgmtUserInfo := new(types.Auth0MgmtUserInfoResponse)
	mgmtUserInfoURL := fmt.Sprintf("%susers/%s", auth0MgmtCfg.MgmtAPIURL, userInfo.Sub)
	appToAppToken, err := s.Auth.GetAppToAppToken()
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", appToAppToken.AccessToken),
	}

	err = s.HTTPClient.GetJSON(mgmtUserInfoURL, headers, mgmtUserInfo)
	if err != nil {
		return nil, err
	}

	return mgmtUserInfo, nil
}

// GetBasicUserInfoFromToken returns the basic user info from the provided token.
func (s *UsersService) GetBasicUserInfoFromToken(authorizationHeader string) (*ctypes.Auth0UserInfoResponse, error) {
	auth0MgmtCfg := s.Config.GetAuth0ManagementConfig()
	headers := map[string]string{
		"Authorization": authorizationHeader,
	}
	userInfo := new(ctypes.Auth0UserInfoResponse)
	err := s.HTTPClient.GetJSON(auth0MgmtCfg.UserInfoURL, headers, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
