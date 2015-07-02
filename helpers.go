package gopher

import "net/http"

type handlerFn func(rw http.ResponseWriter, req *http.Request)

//
func logRoute(c appContainer, rw http.ResponseWriter, req *http.Request) {
	c.providers[LOGGER].(Loggable).Info("[%s] %s", req.Method, req.URL.Path)
}
