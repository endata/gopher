package gopher

import (
	. "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-services"
)

//const (
//	LOGGER   = framework.LOGGER
//	ROUTER   = framework.ROUTER
//	RENDERER = framework.RENDERER
//	PARAMS   = framework.PARAMS
//	SAMPLE   = framework.SAMPLE
//)

var (
	Log     Loggable
	Router  Routable
	Context Mappable
)

func init() {
	initialize()
}

func initialize(config ...Config) *Container {
	appConf := Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	container := NewContainer(appConf)
	container.Use(LoggerMiddleware)
	registerProviders(container)
	return container
}

func registerProviders(c *Container) {

	c.RegisterProvider(new(services.LogProvider))
	Log = c.Log

	c.RegisterProvider(new(services.MapProvider))
	Context = c.Context

	c.RegisterProvider(new(services.RouteProvider))
	Router = c.Router

	c.RegisterProvider(new(services.ParameterProvider))

	c.RegisterProvider(new(services.RenderProvider))
}

func ListenAndServe() {
	Router.Serve()
}
