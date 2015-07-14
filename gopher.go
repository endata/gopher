package gopher

import (
	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-services"
)

var (
	c       *f.Container
	App     *app
	Log     f.Loggable
	Router  f.Routable
	Context f.Mappable
	Render  f.Renderable
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
	initApp()
}

func initApp() {
	app := new(app)
	app.ctnr = c
	app.LOGGER = f.LOGGER
	app.ROUTER = f.ROUTER
	app.RENDERER = f.RENDERER
	app.PARAMS = f.PARAMS
	app.SAMPLE = f.SAMPLE
	App = app
}

func registerProviders() {

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

type app struct {
	ctnr     *f.Container
	LOGGER   string
	ROUTER   string
	RENDERER string
	PARAMS   string
	SAMPLE   string
}

func (m *app) Use(mw f.MiddlewareHandler, args ...interface{}) {
	m.ctnr.Use(mw, args...)
}

func ListenAndServe() {
	Router.Serve()
}
