package App

import (
	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	request  *events.APIGatewayProxyRequest
	response *events.APIGatewayProxyResponse
	routes   []Route
}

type Route struct {
	method  string
	path    string
	handler Handler
}

func (router *Router) NewRouter(response *events.APIGatewayProxyResponse, request *events.APIGatewayProxyRequest) Router {
	r := Router{
		request:  request,
		response: response,
	}
	return r
}
