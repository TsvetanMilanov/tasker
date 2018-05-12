package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
	"github.com/TsvetanMilanov/tasker-common/common"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/data"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/handlers"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func getWorkflow(bootstrap workflow.Bootstrap) *workflow.APIGatewayProxyWorkflow {
	return workflow.NewAPIGWProxyWorkflowBuilder().
		AddPostHandler("/tasks", handlers.CreateHandler).WithPreActions(common.ValidateRequest).
		SetBootstrap(bootstrap).
		Build()
}

func main() {
	bootstrap := common.CreateBootstrap(
		&di.Dependency{Value: &types.CreateHandler{}},
		&di.Dependency{Value: &data.DB{}},
		&di.Dependency{Value: &data.DBClient{}},
	)

	w := getWorkflow(bootstrap)
	lambda.Start(w.GetLambdaHandler())
}
