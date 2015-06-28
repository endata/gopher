package gopher

import "github.com/gopherlabs/gopher/contracts"

var context = Context{}

type Context struct {
	Logger contracts.Log
	Router contracts.Routerable
}

func GetContext() *Context {
	return &context
}
