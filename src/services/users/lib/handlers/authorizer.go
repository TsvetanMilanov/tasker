package handlers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker-common/common/cconstants"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/declarations"
	"github.com/aws/aws-lambda-go/events"
)

// AuthorizerHandler is the handler for the API Gateway Custom Authorizer.
func AuthorizerHandler(ctx workflow.Context, evt events.APIGatewayCustomAuthorizerRequest) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = getUnauthorizedError(errors.New("panic"))
		}
	}()
	usersService := new(declarations.IUsersService)
	err = ctx.GetInjector().Resolve(usersService)
	if err != nil {
		return getUnauthorizedError(err)
	}

	userInfo, err := (*usersService).GetBasicUserInfoFromToken(evt.AuthorizationToken)
	if err != nil {
		return getUnauthorizedError(err)
	}

	bytesUserInfo, err := json.Marshal(userInfo)
	if err != nil {
		return getUnauthorizedError(err)
	}

	res := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID:    userInfo.Sub,
		PolicyDocument: getDummyPolicyDocument(evt),
		Context: map[string]interface{}{
			cconstants.UserInfoContextKey: string(bytesUserInfo),
		},
	}
	ctx.SetRawResponse(res)
	return nil
}

func getUnauthorizedError(err error) error {
	fmt.Println("Authorizer error:", err)
	return errors.New("Unauthorized")
}

func getDummyPolicyDocument(evt events.APIGatewayCustomAuthorizerRequest) events.APIGatewayCustomAuthorizerPolicy {
	denyStatement := events.IAMPolicyStatement{
		Action:   []string{"execute-api:Invoke"},
		Effect:   "Allow",
		Resource: []string{evt.MethodArn},
	}
	return events.APIGatewayCustomAuthorizerPolicy{
		Version: "2012-10-17",
		Statement: []events.IAMPolicyStatement{
			denyStatement,
		},
	}
}
