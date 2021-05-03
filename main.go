package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("method of request : ", req.HTTPMethod)
	return events.APIGatewayProxyResponse{Body: "pong", StatusCode: 200}, nil
}
