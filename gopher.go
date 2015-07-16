package gopher

import (
	f "github.com/gopherlabs/gopher-framework"
	"github.com/gopherlabs/gopher-services"
)

var (
	c          *f.Container
	App        *app
	Log        f.Loggable
	Route      f.Routable
	RouteGroup f.Routegroupable
	Context    f.Mappable
	Render     f.Renderable
)

type app struct {
	ctnr *f.Container
}

func init() {
	initApp()
	App.Config()
}

func initApp() {
	app := new(app)
	App = app
}

type Config map[string]interface{}
type ConfigLogger f.ConfigLogger
type ConfigRouter f.ConfigRouter
type ConfigRenderer f.ConfigRenderer

func (m *app) Config(config ...map[string]interface{}) {
	if len(config) > 0 {
		c = f.NewContainer(convertConfig(config[0]))
	} else {
		c = f.NewContainer()
	}
	App.ctnr = c
	c.Use(f.LoggerMiddleware)
	registerProviders()
}

func (m *app) Use(mw f.MiddlewareHandler, args ...interface{}) {
	m.ctnr.Use(mw, args...)
}

func convertConfig(in Config) f.Config {
	conf := f.Config{}
	if in[KEY_LOGGER] != nil {
		conf[KEY_LOGGER] = f.ConfigLogger(in[KEY_LOGGER].(ConfigLogger))
	}
	if in[KEY_ROUTER] != nil {
		conf[KEY_ROUTER] = f.ConfigRouter(in[KEY_ROUTER].(ConfigRouter))
	}
	if in[KEY_RENDERER] != nil {
		conf[KEY_RENDERER] = f.ConfigRenderer(in[KEY_RENDERER].(ConfigRenderer))
	}
	return conf
}

func registerProviders() {
	c.RegisterProvider(new(services.LogProvider))
	Log = c.Log
	c.RegisterProvider(new(services.MapProvider))
	Context = c.Context
	c.RegisterProvider(new(services.RouteProvider))
	Route = c.Route
	RouteGroup = c.RouteGroup
	c.RegisterProvider(new(services.RenderProvider))
	Render = c.Render
	f.Initialized = true
}

func ListenAndServe() {
	Route.(f.Servable).Serve()
}

const (
	KEY_LOGGER   = f.LOGGER
	KEY_ROUTER   = f.ROUTER
	KEY_RENDERER = f.RENDERER
	KEY_MAPPER   = f.MAPPER
)

// Log Levels
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	LEVEL_PANIC uint8 = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	LEVEL_FATAL
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LEVEL_ERROR
	// WarnLevel level. Non-critical entries that deserve eyes.
	LEVEL_WARN
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	LEVEL_INFO
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	LEVEL_DEBUG
)
