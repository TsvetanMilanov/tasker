package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker-common/common/cutils"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/types"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/types/requests"
)

// CreateHandler handles task creation.
func CreateHandler(ctx workflow.Context, req requests.CreateTask) error {
	h := new(types.CreateHandler)
	err := ctx.GetInjector().Resolve(h)
	if err != nil {
		return err
	}

	taskID, err := h.Tasks.Create(req, "test")
	if err != nil {
		cutils.SetInternalServerError(ctx, err)
		return nil
	}

	ctx.SetResponse(taskID).SetResponseStatusCode(http.StatusCreated)
	return nil
}
