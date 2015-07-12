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
	log     Loggable
	router  Routable
	context Mappable
)

func Log() Loggable {
	return log
}

func Router() Routable {
	return router
}

func Context() Mappable {
	return context
}

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
	log = c.Log

	c.RegisterProvider(new(services.MapProvider))
	context = c.Context

	c.RegisterProvider(new(services.RouteProvider))
	router = c.Router

	c.RegisterProvider(new(services.ParameterProvider))

	c.RegisterProvider(new(services.RenderProvider))
}

func ListenAndServe() {
	router.Serve()
}
