package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker-common/common/cutils"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
)

const (
	redirectURI     = "https://4qxq3qd0ek.execute-api.us-east-1.amazonaws.com/dev/callback"
	contentTypeJSON = "application/json"
)

type authorizationCode struct {
	Code string `json:"code"`
}

type codeGrantRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type codeGrantResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CallbackHandler handles the oauth2 authorization code request.
func CallbackHandler(ctx workflow.Context, req authorizationCode) error {
	h := new(types.BaseHandler)
	err := ctx.GetInjector().Resolve(h)
	if err != nil {
		return err
	}
	auth0Cfg := h.Config.GetAuth0Config()
	cgReq := codeGrantRequest{
		Code:         req.Code,
		ClientID:     auth0Cfg.ClientID,
		ClientSecret: auth0Cfg.ClientSecret,
		GrantType:    "authorization_code",
		RedirectURI:  redirectURI,
	}

	rBodyBytes := []byte{}
	err = h.HTTPClient.PostJSON(auth0Cfg.TokenURL, cgReq, nil, &rBodyBytes)
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	encRes := base64.StdEncoding.EncodeToString(rBodyBytes)
	location := fmt.Sprintf("http://localhost?login_response=%s", encRes)
	fmt.Println(location)
	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": location,
		},
	}

	ctx.SetRawResponse(res)
	return nil
}
