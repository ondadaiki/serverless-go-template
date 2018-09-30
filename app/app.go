package app

import (
	"database/sql"
	"serverless-go-template/config"
	"serverless-go-template/controller"
	"serverless-go-template/controller/user"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type App struct {
	routes  []Route
	request events.APIGatewayProxyRequest
	config  config.Config
	db      *sql.DB
}

type Request struct {
	events.APIGatewayProxyRequest
}

func NewApp(request events.APIGatewayProxyRequest) (*App, error) {
	app := &App{
		request: request,
	}

	app.initRouter()
	err := app.initDb()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) initRouter() {
	app.addRoute("GET", "/users/{id}", user.Index)
}

func (app *App) initDb() error {
	db, err := sql.Open("postgres", app.config.Postgres.ConnectionString())
	if err != nil {
		return err
	}
	app.db = db

	return nil
}

func (app *App) addRoute(method, path string, controller controller.Controller) {
	route := Route{
		method:     strings.ToUpper(method),
		path:       path,
		controller: controller,
	}
	app.routes = append(app.routes, route)
}

func (app *App) Run() (events.APIGatewayProxyResponse, error) {
	for _, value := range app.routes {
		if value.path == app.request.Resource {
			if value.method == app.request.HTTPMethod {
				return value.controller(app.request, app.db)
			}
		}
	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 404}, nil
}
