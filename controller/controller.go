package controller

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
)

type Controller func(events.APIGatewayProxyRequest, *sql.DB) (events.APIGatewayProxyResponse, error)
