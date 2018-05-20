package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
	"github.com/TsvetanMilanov/tasker-common/common"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/handlers"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/services"
	"github.com/TsvetanMilanov/tasker/src/services/users/lib/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func getWorkflow(bootstrap workflow.Bootstrap) *workflow.APIGatewayAuthorizerWorkflow {
	return workflow.NewAPIGWAuthorizerWorkflowBuilder().
		SetHandler(handlers.AuthorizerHandler).
		SetBootstrap(bootstrap).
		Build()
}

func main() {
	bootstrap := common.CreateBootstrap(
		&di.Dependency{Value: new(types.BaseHandler)},
		&di.Dependency{Value: new(services.AuthService)},
		&di.Dependency{Value: new(services.UsersService)},
	)
	w := getWorkflow(bootstrap)
	lambda.Start(w.GetLambdaHandler())
}
