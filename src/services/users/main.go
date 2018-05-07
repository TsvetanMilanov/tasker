package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
	"github.com/TsvetanMilanov/tasker/src/common"
	"github.com/TsvetanMilanov/tasker/src/services/user/handlers"
	"github.com/TsvetanMilanov/tasker/src/services/user/types"
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
