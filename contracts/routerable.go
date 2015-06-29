package contracts

import "net/http"

type Routerable interface {
	http.Handler

	NewRouter() Routerable

	Get(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Head(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Post(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Put(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Patch(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Delete(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Options(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Match(verbs []string, path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	All(path string, fn func(http.ResponseWriter, *http.Request)) Routerable

	Serve()
}
