package handlers

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

type createTaskReq struct {
	Name  string          `json:"name"`
	Tasks []createTaskReq `json:"tasks"`
}

// CreateHandler handles task creation.
func CreateHandler(ctx workflow.Context, req createTaskReq) error {
	return nil
}
