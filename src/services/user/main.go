package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker/src/services/user/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	w := workflow.NewAPIGWProxyWorkflowBuilder().
		AddGetHandler("/info", handlers.InfoHandler).
		Build()
	lambda.Start(w.GetLambdaHandler())
}
