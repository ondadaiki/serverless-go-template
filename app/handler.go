package App

import "github.com/aws/aws-lambda-go/events"

type Handler func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
