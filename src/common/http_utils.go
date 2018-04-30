package common

import (
	"fmt"
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

func setInternalServerError(ctx workflow.Context, err error) {
	fmt.Println(err)
	e := struct {
		Message string `json:"message"`
	}{Message: "Internal server error"}

	ctx.
		SetResponse(e).
		SetResponseStatusCode(http.StatusInternalServerError)
}
