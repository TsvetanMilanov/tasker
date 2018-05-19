package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
	"github.com/TsvetanMilanov/tasker-common/common"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/handlers"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func getWorkflow(bootstrap workflow.Bootstrap) *workflow.APIGatewayProxyWorkflow {
	return workflow.NewAPIGWProxyWorkflowBuilder().
		AddGetHandler("/me", handlers.InfoHandler).
		SetBootstrap(bootstrap).
		Build()
}

func main() {
	bootstrap := common.CreateBootstrap(
		&di.Dependency{Value: &types.InfoHandler{}},
	)

	w := getWorkflow(bootstrap)
	lambda.Start(w.GetLambdaHandler())
}
