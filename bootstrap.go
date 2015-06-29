package gopher

import (
	"fmt"
	"github.com/gopherlabs/gopher-services/providers"
)

func Start() {
	fmt.Println("Starting Gopher...")
	RegisterProviders()
}

func RegisterProviders() {
	GetContainer().Logger = providers.Logger{}
	GetContainer().Router = providers.RouteProvider{}
	GetContainer().Parameters = providers.ParameterProvider{}
}
