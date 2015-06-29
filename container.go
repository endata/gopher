package gopher

import "github.com/gopherlabs/gopher/contracts"

var ctnr = container{}

type container struct {
	Logger     contracts.Loggable
	Router     contracts.Routable
	Parameters contracts.Parametable
	Renderer   contracts.Renderable
}

func getContainer() *container {
	return &ctnr
}
