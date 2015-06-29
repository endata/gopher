package contracts

import "net/http"

type Routable interface {
	http.Handler

	NewRouter() Routable

	Get(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Head(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Post(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Put(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Patch(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Delete(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Options(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Match(verbs []string, path string, fn func(http.ResponseWriter, *http.Request)) Routable

	All(path string, fn func(http.ResponseWriter, *http.Request)) Routable

	Serve()
}
