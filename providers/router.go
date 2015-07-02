package providers

import (
	"net/http"

	"github.com/gopherlabs/gopher/vendor/_nuts/github.com/gorilla/mux"
)

type RouteProvider struct {
	http.Handler
	mux *mux.Router
}

func (r RouteProvider) Register(config map[string]interface{}) interface{} {
	return r
}

func (l RouteProvider) GetKey() string {
	return "ROUTER"
}

func (r RouteProvider) Router() interface{} {
	r.mux = mux.NewRouter()
	return r
}

func (r RouteProvider) Get(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("GET")
	return r
}

func (r RouteProvider) Head(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("HEAD")
	return r
}

func (r RouteProvider) Post(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("POST")
	return r
}

func (r RouteProvider) Put(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("PUT")
	return r
}

func (r RouteProvider) Patch(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("PATCH")
	return r
}

func (r RouteProvider) Delete(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("DELETE")
	return r
}

func (r RouteProvider) Options(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods("OPTIONS")
	return r
}

func (r RouteProvider) Match(verbs []string, path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn).Methods(verbs...)
	return r
}

func (r RouteProvider) All(path string, fn func(http.ResponseWriter, *http.Request)) interface{} {
	r.mux.HandleFunc(path, fn)
	return r
}

func (r RouteProvider) Serve() {
	http.ListenAndServe("0.0.0.0:3000", r.mux)
}

func (r RouteProvider) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}
