package gopher

import "net/http"

type handlerFn func(rw http.ResponseWriter, req *http.Request)

func (c appContainer) Get(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		c.providers[LOGGER].(Loggable).Info("[%s] %s", req.Method, req.URL.Path)
		fn(rw, req)
	}
	c.providers[ROUTER].(Routable).Get(path, nfn)
}

//
func (c appContainer) Serve() {
	c.providers[ROUTER].(Routable).Serve()
}

// Parameters
func (c appContainer) PathParams(req *http.Request) map[string]string {
	return c.providers[PARAMS].(Parametable).PathParams(req)
}

// Parameters
func (c appContainer) PathParam(req *http.Request, param string) string {
	return c.providers[PARAMS].(Parametable).PathParam(req, param)
}

// Logger
func (c appContainer) Log() Loggable {
	return c.providers[LOGGER].(Loggable)
}

// Renderer
func (c appContainer) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	c.providers[RENDERER].(Renderable).View(rw, status, name, binding)
}
