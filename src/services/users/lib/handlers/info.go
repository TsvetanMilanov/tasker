package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/tasker-common/common/cutils"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/declarations"
)

type userInfoRequest struct {
	Authorization string
}

// InfoHandler handles user info request.
func InfoHandler(ctx workflow.Context, req userInfoRequest) error {
	usersService := new(declarations.IUsersService)
	err := ctx.GetInjector().Resolve(usersService)
	if err != nil {
		return err
	}

	mgmtUserInfo, err := (*usersService).GetUserInfoFromToken(req.Authorization)
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	ctx.SetResponse(mgmtUserInfo).SetResponseStatusCode(http.StatusOK)
	return nil
}
