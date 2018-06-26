package main

import (
	"context"
	"fmt"

	"serverless-go-template/App"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is Lambda Handler function
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app, err := App.NewApp(request)
	if err != nil {
		fmt.Println("Fetal error")
	}
	res, err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
	return res, nil

}

func main() {
	lambda.Start(Handler)
}
