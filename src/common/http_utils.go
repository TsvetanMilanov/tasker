package common

import (
	"fmt"
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

// SetInternalServerError sets the response to be struct with message
// "Internal server error" and the status code to 500. Also it prints the
// original error.
func SetInternalServerError(ctx workflow.Context, err error) {
	fmt.Println(err)
	e := struct {
		Message string `json:"message"`
	}{Message: "Internal server error"}

	ctx.
		SetResponse(e).
		SetResponseStatusCode(http.StatusInternalServerError)
}
