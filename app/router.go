package App

type Route struct {
	method  string
	path    string
	handler Handler
}

func (app *App) NewRouter() []Route {
	r := []Route{}
	return r
}
