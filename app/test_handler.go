package App

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func (app *App) HandleTest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonByteArr, _ := json.Marshal(request)

	return events.APIGatewayProxyResponse{Body: string(jsonByteArr), StatusCode: 200}, nil

}
