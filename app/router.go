package app

import "serverless-go-template/controller"

type Route struct {
	method     string
	path       string
	controller controller.Controller
}

func (app *App) NewRouter() []Route {
	r := []Route{}
	return r
}
