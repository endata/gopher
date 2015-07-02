package gopher

import "net/http"

// Routers
func (c appContainer) Router() Routable {
	return container.providers[ROUTER].(Routable).Router().(Routable)
}

// Parameters
func (c appContainer) PathParams(req *http.Request) map[string]string {
	return container.providers[PARAMS].(Parametable).PathParams(req)
}

// Parameters
func (c appContainer) PathParam(req *http.Request, param string) string {
	return container.providers[PARAMS].(Parametable).PathParam(req, param)
}

// Logger
func (c appContainer) Log() Loggable {
	return container.providers[LOGGER].(Loggable)
}

// Renderer
func (c appContainer) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	container.providers[RENDERER].(Renderable).View(rw, status, name, binding)
}
