package gopher

import "net/http"

type Providerable interface {
	Register(config map[string]string) interface{}
	GetKey() string
}

type Loggable interface {
	Providerable
	//Log() interface{}
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Panic(msg string, args ...interface{})
}

type Parametable interface {
	Providerable
	PathParams(r *http.Request) map[string]string
	PathParam(r *http.Request, param string) string
}

type Renderable interface {
	Providerable
	NewRenderer() interface{}
	View(rw http.ResponseWriter, status int, name string, binding interface{})
}

type Routable interface {
	Providerable
	http.Handler
	Router() interface{}
	Get(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Head(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Post(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Put(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Patch(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Delete(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Options(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Match(verbs []string, path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	All(path string, fn func(http.ResponseWriter, *http.Request)) interface{}
	Serve()
}
