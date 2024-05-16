package common

import "github.com/aws/aws-lambda-go/events"

type HandlerFunc func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
