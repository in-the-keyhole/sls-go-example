package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var message = "We are the world, from AWS Lambda using Golang!"

	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
