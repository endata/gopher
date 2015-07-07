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
	appConf := framework.Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	container := framework.NewContainer(appConf)
	registerProviders(container)
	container.Use(framework.LoggerMiddleware)
	return container
}

func registerProviders(container *framework.Container) {
	container.RegisterProvider(new(services.LogProvider))
	container.RegisterProvider(new(services.RouteProvider))
	container.RegisterProvider(new(services.ParameterProvider))
	container.RegisterProvider(new(services.RenderProvider))
}
