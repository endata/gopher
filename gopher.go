package gopher

import (
	"github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-services"
)

const (
	LOGGER   = framework.LOGGER
	ROUTER   = framework.ROUTER
	RENDERER = framework.RENDERER
	PARAMS   = framework.PARAMS
	SAMPLE   = framework.SAMPLE
)

func NewApp(config ...framework.Config) *framework.Container {
	container := framework.NewContainer(config[0])
	registerProviders(container)
	container.ShowBanner()
	return container
}

func registerProviders(container *framework.Container) {
	container.RegisterProvider(new(services.LogProvider))
	container.RegisterProvider(new(services.RouteProvider))
	container.RegisterProvider(new(services.ParameterProvider))
	container.RegisterProvider(new(services.SampleProvider))
	container.RegisterProvider(new(services.RenderProvider))
}
