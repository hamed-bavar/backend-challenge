package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func main() {
	lambda.Start(
		func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			router := mux.NewRouter()
			return gorillamux.New(router).Proxy(req)
		})
}
