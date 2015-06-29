package gopher

import (
	"github.com/gopherlabs/gopher-services/providers"
)

func App() container {
	registerProviders()
	container := getContainer()
	container.Logger.NewLogger().Info("Starting Gopher...")
	return container
}

func registerProviders() {
	container := getContainer()
	container.Logger = providers.LogProvider{}
	container.Router = providers.RouteProvider{}
	container.Parameters = providers.ParameterProvider{}
	container.Renderer = providers.RenderProvider{}
}
