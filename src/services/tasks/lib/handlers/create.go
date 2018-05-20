package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker-common/common/cutils"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types/requests"
)

// CreateHandler handles task creation.
func CreateHandler(ctx workflow.Context, req requests.CreateTask) error {
	h := new(types.CreateHandler)
	err := ctx.GetInjector().Resolve(h)
	if err != nil {
		return err
	}

	userInfo, err := cutils.GetAuthorizerUserFromContext(ctx)
	if err != nil {
		return cutils.SetInternalServerError(ctx, err)
	}

	taskID, err := h.Tasks.Create(req, userInfo.Sub)
	if err != nil {
		return cutils.SetInternalServerError(ctx, err)
	}

	ctx.SetResponse(taskID).SetResponseStatusCode(http.StatusCreated)
	return nil
}
