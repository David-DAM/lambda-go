package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

/**
 * Lambda function handler for API requests
 *
 * Handles requests to /api/hello and /api/key paths
 *
 * Uses Amazon Linux 2 as runtime with architecture arm64
 * Compiles with: GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
 * Sets runtime settings in AWS Lambda console handler to main.handler
 */
func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {

	path := request.RequestContext.HTTP.Path

	var response events.LambdaFunctionURLResponse

	switch path {
	case "/api/hello":
		response = handleHello(request)
	case "/api/key":
		response = handleKey(request)
	default:
		response = handleNotFound(request)
	}

	return response, nil
}

func handleHello(request events.LambdaFunctionURLRequest) events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}
}

func handleKey(request events.LambdaFunctionURLRequest) events.LambdaFunctionURLResponse {
	apiKey := request.Headers["Authorization"]
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       apiKey,
	}
}

func handleNotFound(request events.LambdaFunctionURLRequest) events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		StatusCode: 404,
		Body:       "Not Found",
	}
}
