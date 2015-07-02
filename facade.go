package gopher

import "net/http"

//
func (c appContainer) Get(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Get(path, nfn)
}

func (c appContainer) Head(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Head(path, nfn)
}

func (c appContainer) Post(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Post(path, nfn)
}

func (c appContainer) Put(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Put(path, nfn)
}

func (c appContainer) Patch(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Patch(path, nfn)
}

func (c appContainer) Delete(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Delete(path, nfn)
}

func (c appContainer) Options(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Options(path, nfn)
}

func (c appContainer) Match(path string, fn handlerFn, verbs ...string) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).Match(path, nfn, verbs...)
}

func (c appContainer) All(path string, fn handlerFn) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	c.providers[ROUTER].(Routable).All(path, nfn)
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
