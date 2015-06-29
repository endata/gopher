package gopher

import "github.com/gopherlabs/gopher/contracts"

var context = Context{}

type Context struct {
	Logger     contracts.Loggable
	Router     contracts.Routable
	Parameters contracts.Parametable
}

func GetContext() *Context {
	return &context
}
