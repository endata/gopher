package contracts

import "net/http"

type Routerable interface {
	http.Handler
	NewRouter() Routerable
	Get(path string, fn func(http.ResponseWriter, *http.Request)) Routerable
	Serve()
}
