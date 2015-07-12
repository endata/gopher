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

var (
	Log    framework.Loggable
	Router framework.Routable
)

func Initialize(config ...framework.Config) *framework.Container {
	appConf := framework.Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	container := framework.NewContainer(appConf)
	container.Use(framework.LoggerMiddleware)
	registerProviders(container)
	return container
}

func registerProviders(c *framework.Container) {
	c.RegisterProvider(new(services.LogProvider))
	Log = c.Log

	c.RegisterProvider(new(services.RouteProvider))
	Router = c.Router

	c.RegisterProvider(new(services.ParameterProvider))
	c.RegisterProvider(new(services.RenderProvider))
}

func ListenAndServe() {
	Router.Serve()
}
