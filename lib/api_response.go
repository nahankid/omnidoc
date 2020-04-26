package lib

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// APIResponse returns APIGatewayProxyResponse
func APIResponse(code int, msg string) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=UTF-8"

	log.Printf("%d - %s\n", code, msg)

	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
		Headers:    headers,
	}, nil
}
