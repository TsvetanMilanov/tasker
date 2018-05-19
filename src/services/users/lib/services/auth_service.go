package services

import (
	"github.com/TsvetanMilanov/tasker-common/common/cdeclarations"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
)

// AuthService implements IAuthService.
type AuthService struct {
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}

// GetAppToAppToken returns new auth0 machine to machine token.
func (s *AuthService) GetAppToAppToken() (*types.TokenResponse, error) {
	auth0MgmtCfg := s.Config.GetAuth0ManagementConfig()
	body := auth0MgmtMachineToMachineTokenRequest{
		GrantType:    "client_credentials",
		ClientID:     auth0MgmtCfg.ClientID,
		ClientSecret: auth0MgmtCfg.ClientSecret,
		Audience:     auth0MgmtCfg.MgmtAPIURL,
	}
	machineToMachineToken := new(types.TokenResponse)
	err := s.HTTPClient.PostJSON(auth0MgmtCfg.TokenURL, body, nil, machineToMachineToken)
	if err != nil {
		return nil, err
	}

	return machineToMachineToken, nil
}

type auth0MgmtMachineToMachineTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
}
