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
	c          *f.Container
	Middleware *middleware
	Log        f.Loggable
	Router     f.Routable
	Context    f.Mappable
	Render     f.Renderable
)

func init() {
	Config()
}

func Config(config ...f.Config) {
	appConf := f.Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	c = f.NewContainer(appConf)
	c.Use(f.LoggerMiddleware)
	registerProviders()
}

func registerProviders() {

	mware := new(middleware)
	mware.ctnr = c
	Middleware = mware

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

type middleware struct {
	ctnr *f.Container
}

func (m *middleware) Use(mw f.MiddlewareHandler, args ...interface{}) {
	m.ctnr.Use(mw, args...)
}

func ListenAndServe() {
	Router.Serve()
}
