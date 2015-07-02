package gopher

import "net/http"

type Providerable interface {
	Register(config map[string]interface{}) interface{}
	GetKey() string
}

type Loggable interface {
	Providerable
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
	View(rw http.ResponseWriter, status int, name string, binding interface{})
}

type Routable interface {
	Providerable
	http.Handler
	Get(path string, fn func(http.ResponseWriter, *http.Request))
	Head(path string, fn func(http.ResponseWriter, *http.Request))
	Post(path string, fn func(http.ResponseWriter, *http.Request))
	Put(path string, fn func(http.ResponseWriter, *http.Request))
	Patch(path string, fn func(http.ResponseWriter, *http.Request))
	Delete(path string, fn func(http.ResponseWriter, *http.Request))
	Options(path string, fn func(http.ResponseWriter, *http.Request))
	Match(path string, fn func(http.ResponseWriter, *http.Request), verbs ...string)
	All(path string, fn func(http.ResponseWriter, *http.Request))
	Serve()
}
