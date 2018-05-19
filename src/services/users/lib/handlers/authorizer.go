package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/events"
)

// AuthorizerHandler is the handler for the API Gateway Custom Authorizer.
func AuthorizerHandler(ctx workflow.Context, evt events.APIGatewayCustomAuthorizerRequest) error {
	r, _ := json.Marshal(ctx)
	fmt.Println(string(r))

	r, _ = json.Marshal(evt)
	fmt.Println(string(r))
	ctx.SetRawResponse(events.APIGatewayCustomAuthorizerResponse{})
	return nil
}
