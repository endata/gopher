package gopher

import "github.com/gopherlabs/gopher/contracts"

var container = Container{}

type Container struct {
	Logger     contracts.Loggable
	Router     contracts.Routable
	Parameters contracts.Parametable
}

func GetContainer() *Container {
	return &container
}
