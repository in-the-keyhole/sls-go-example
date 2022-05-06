package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var message = fmt.Sprintf("Hello %s, from AWS Lambda using Golang!", request.QueryStringParameters["name"])

	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
