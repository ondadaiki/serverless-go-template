package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"serverless-go-template/repository"

	"github.com/aws/aws-lambda-go/events"
)

func Index(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()
	user, _ := repository.FindUser(db, ctx, request.PathParameters["id"])
	body, _ := json.Marshal(user)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil

}
