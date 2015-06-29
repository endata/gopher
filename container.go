package gopher

import "github.com/gopherlabs/gopher/contracts"

var ctnr = container{}

type container struct {
	logger     contracts.Loggable
	router     contracts.Routable
	parameters contracts.Parametable
	renderer   contracts.Renderable
}

func getContainer() *container {
	return &ctnr
}
