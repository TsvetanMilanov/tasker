package utils

import (
	"fmt"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/events"
)

// GetServiceURLFromEvent returns the full API Gateway url + the full path.
func GetServiceURLFromEvent(ctx workflow.Context) (string, error) {
	evt := new(events.APIGatewayProxyRequest)
	err := ctx.GetLambdaEvent(evt)
	if err != nil {
		return "", err
	}

	host := evt.Headers["Host"]
	return fmt.Sprintf("https://%s%s", host, evt.Path), nil
}
