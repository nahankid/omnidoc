package main

import (
	"net/http"
	"omnidoc/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// CreateRequest struct
type CreateRequest struct {
	UserID int    `json:"user_id"`
	AppID  int    `json:"app_id"`
	Type   string `json:"type"`
	Attrs  string `json:"attrs"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return lib.APIResponse(http.StatusOK, "")
}

func main() {
	lambda.Start(handler)
}
