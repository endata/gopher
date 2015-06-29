package gopher

import (
	"fmt"
	"github.com/gopherlabs/gopher-services/providers"
)

func Start() {
	RegisterProviders()
	fmt.Println("Starting Gopher...")
}

func RegisterProviders() {
	container := GetContainer()
	container.Logger = providers.Logger{}
	container.Router = providers.RouteProvider{}
	container.Parameters = providers.ParameterProvider{}
}
