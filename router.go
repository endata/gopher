package gopher

import (
	"github.com/gopherlabs/gopher/contracts"
)

func NewRouter(ctx interface{}) contracts.Routerable {
	return GetContext().Router.NewRouter(ctx)
}

func Get(path string, fn interface{}) contracts.Routerable {
	return GetContext().Router.Get(path, fn)
}
