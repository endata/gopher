package gopher

import (
	f "github.com/gopherlabs/gopher-framework"
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
	Log     f.Loggable
	Router  f.Routable
	Context f.Mappable
	Render  f.Renderable
)

func init() {
	Config()
}

func Config(config ...f.Config) *f.Container {
	appConf := f.Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	container := f.NewContainer(appConf)
	container.Use(f.LoggerMiddleware)
	registerProviders(container)
	return container
}

func registerProviders(c *f.Container) {

	c.RegisterProvider(new(services.LogProvider))
	Log = c.Log

	c.RegisterProvider(new(services.MapProvider))
	Context = c.Context

	c.RegisterProvider(new(services.RouteProvider))
	Router = c.Router

	c.RegisterProvider(new(services.ParameterProvider))

	c.RegisterProvider(new(services.RenderProvider))
	Render = c.Render
}

func ListenAndServe() {
	Router.Serve()
}
