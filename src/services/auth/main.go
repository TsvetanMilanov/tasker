package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
	"github.com/TsvetanMilanov/tasker/src/common"
	"github.com/TsvetanMilanov/tasker/src/services/auth/handlers"
	"github.com/TsvetanMilanov/tasker/src/services/auth/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func getWorkflow(bootstrap workflow.Bootstrap) *workflow.APIGatewayProxyWorkflow {
	return workflow.NewAPIGWProxyWorkflowBuilder().
		AddGetHandler("/callback", handlers.CallbackHandler).
		SetBootstrap(bootstrap).
		Build()
}

func main() {
	bootstrap := func() workflow.Injector {
		c := di.NewContainer()
		err := c.Register(
			&di.Dependency{Value: &types.CallbackHandler{}},
			&di.Dependency{Value: &common.Config{}},
			&di.Dependency{Value: &common.HTTPClient{}},
		)
		if err != nil {
			panic(err)
		}

		return c
	}
	w := getWorkflow(bootstrap)
	lambda.Start(w.GetLambdaHandler())
}
