package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/tasker/src/services/auth/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	w := workflow.NewAPIGWProxyWorkflowBuilder().
		AddGetHandler("/callback", handlers.CallbackHandler).
		Build()
	lambda.Start(w.GetLambdaHandler())
}
