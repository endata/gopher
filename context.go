package gopher

import "github.com/gopherlabs/gopher/contracts"

var context = Context{}

type Context struct {
	Logger contracts.Log
}

func GetContext() *Context {
	return &context
}
