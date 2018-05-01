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
		AddGetHandler("/info", handlers.InfoHandler).
		SetBootstrap(bootstrap).
		Build()
}

func main() {
	bootstrap := func() workflow.Injector {
		c := di.NewContainer()
		err := c.Register(
			&di.Dependency{Value: &types.InfoHandler{}},
			&di.Dependency{Value: &common.HTTPClient{}},
			&di.Dependency{Value: &common.Config{}},
		)
		if err != nil {
			panic(err)
		}

		return c
	}

	w := getWorkflow(bootstrap)
	lambda.Start(w.GetLambdaHandler())
}
