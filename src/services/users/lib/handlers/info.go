package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker-common/common/cutils"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
)

type userInfoRequest struct {
	Authorization string
}

type auth0MgmtMachineToMachineTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
}

type auth0MgmtMachineToMachineTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type auth0UserInfoResponse struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type auth0MgmtUserInfoResponse struct {
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

// InfoHandler handles user info request.
func InfoHandler(ctx workflow.Context, req userInfoRequest) error {
	infoHandler := &types.InfoHandler{}
	err := ctx.GetInjector().Resolve(infoHandler)
	if err != nil {
		return err
	}

	auth0MgmtCfg := infoHandler.Config.GetAuth0ManagementConfig()

	headers := map[string]string{
		"Authorization": req.Authorization,
	}
	userInfo := auth0UserInfoResponse{}
	err = infoHandler.HTTPClient.GetJSON(auth0MgmtCfg.UserInfoURL, headers, &userInfo)
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	// Get machine to machine token.
	body := auth0MgmtMachineToMachineTokenRequest{
		GrantType:    "client_credentials",
		ClientID:     auth0MgmtCfg.ClientID,
		ClientSecret: auth0MgmtCfg.ClientSecret,
		Audience:     auth0MgmtCfg.MgmtAPIURL,
	}
	machineToMachineToken := auth0MgmtMachineToMachineTokenResponse{}
	err = infoHandler.HTTPClient.PostJSON(auth0MgmtCfg.TokenURL, body, nil, &machineToMachineToken)
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	mgmtUserInfo := auth0MgmtUserInfoResponse{}
	mgmtUserInfoURL := fmt.Sprintf("%susers/%s", auth0MgmtCfg.MgmtAPIURL, userInfo.Sub)
	headers = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", machineToMachineToken.AccessToken),
	}

	err = infoHandler.HTTPClient.GetJSON(mgmtUserInfoURL, headers, &mgmtUserInfo)
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	fmt.Println(mgmtUserInfo)
	ctx.SetResponse(mgmtUserInfo).SetResponseStatusCode(http.StatusOK)
	return nil
}
