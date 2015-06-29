package gopher

import (
	"github.com/gopherlabs/gopher-services/providers"
)

func Start() {
	RegisterProviders()
	GetContainer().Logger.NewLogger().Info("Starting Gopher...")
}

func RegisterProviders() {
	container := GetContainer()
	container.Logger = providers.LogProvider{}
	container.Router = providers.RouteProvider{}
	container.Parameters = providers.ParameterProvider{}
}
