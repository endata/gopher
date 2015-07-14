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

type app struct {
	ctnr     *f.Container
	LOGGER   string
	ROUTER   string
	RENDERER string
	PARAMS   string
	SAMPLE   string
}

func init() {
	initApp()
	App.Config()
}

func initApp() {
	app := new(app)
	app.LOGGER = f.LOGGER
	app.ROUTER = f.ROUTER
	app.RENDERER = f.RENDERER
	app.PARAMS = f.PARAMS
	app.SAMPLE = f.SAMPLE
	App = app
}

func (m *app) Config(config ...f.Config) {
	appConf := f.Config{}
	if len(config) > 0 {
		appConf = config[0]
	}
	c = f.NewContainer(appConf)
	App.ctnr = c
	c.Use(f.LoggerMiddleware)
	registerProviders()
}

func (m *app) Use(mw f.MiddlewareHandler, args ...interface{}) {
	m.ctnr.Use(mw, args...)
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

func ListenAndServe() {
	Router.(f.Servable).Serve()
}
