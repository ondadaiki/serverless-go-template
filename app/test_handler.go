package App

import (
	"github.com/aws/aws-lambda-go/events"
)

func (app *App) HandleTest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{Body: "aaaaa", StatusCode: 200}, nil

}
