package gopher

import (
	"github.com/gopherlabs/gopher/contracts"
	"net/http"
)

func NewRouter() contracts.Routerable {
	return GetContext().Router.NewRouter()
}

func Get(path string, fn func(http.ResponseWriter, *http.Request)) contracts.Routerable {
	return GetContext().Router.Get(path, fn)
}

func Serve() {
	GetContext().Router.Serve()
}
