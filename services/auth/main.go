package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, evt events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       `{"message": "Hello World!!!"}`,
	}
	return res, nil
}

func main() {
	lambda.Start(handler)
}
