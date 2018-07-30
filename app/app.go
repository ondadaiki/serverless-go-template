package App

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type App struct {
	routes   []Route
	request  events.APIGatewayProxyRequest
	response events.APIGatewayProxyResponse
}

type Request struct {
	events.APIGatewayProxyRequest
}

func NewApp(request events.APIGatewayProxyRequest) (*App, error) {
	app := &App{
		request: request,
	}

	app.initRouter()

	return app, nil
}

func (app *App) initRouter() {
	app.addRoute("GET", "/", app.HandleTest)
}

func (app *App) addRoute(method, path string, handler Handler) {
	route := Route{
		method:  strings.ToUpper(method),
		path:    path,
		handler: handler,
	}
	app.routes = append(app.routes, route)
}

func (app *App) Run() (events.APIGatewayProxyResponse, error) {
	for _, value := range app.routes {
		if value.path == app.request.Resource {
			if value.method == app.request.HTTPMethod {
				return value.handler(app.request)
			}
		}
	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 404}, nil
}
