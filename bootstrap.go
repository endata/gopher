package gopher

import (
	"github.com/gopherlabs/gopher-services/providers"
)

func App() container {
	registerProviders()
	container := getContainer()
	container.logger.NewLogger().Info("Starting Gopher...")
	return *container
}

func registerProviders() {
	container := getContainer()
	container.logger = providers.LogProvider{}
	container.router = providers.RouteProvider{}
	container.parameters = providers.ParameterProvider{}
	container.renderer = providers.RenderProvider{}
}
