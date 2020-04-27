package main

import (
	"net/http"
	"omnidoc/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return lib.APIResponse(http.StatusOK, "")
}

func main() {
	lambda.Start(handler)
}
