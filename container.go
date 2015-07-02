package gopher

import "github.com/gopherlabs/gopher/providers"

const (
	LOGGER   = "LOGGER"
	ROUTER   = "ROUTER"
	RENDERER = "RENDERER"
	PARAMS   = "PARAMS"
)

var (
	container = appContainer{}
)

type appContainer struct {
	config    map[string]map[string]interface{}
	providers map[string]Providerable
}

func App(config ...map[string]map[string]interface{}) *appContainer {
	if len(config) > 0 {
		container.config = config[0]
	}
	registerProviders()
	showBanner()
	return &container
}

func registerProviders() {
	container.providers = map[string]Providerable{}
	register(providers.LogProvider{})
	register(providers.RouteProvider{})
	register(providers.RenderProvider{})
	register(providers.ParameterProvider{})
}

func register(provider interface{}) {
	key := provider.(Providerable).GetKey()
	config := defaultConfig[key]
	if len(container.config) > 0 {
		config = container.config[key]
	}
	container.providers[key] = provider.(Providerable).Register(config).(Providerable)
	if key == LOGGER {
		container.providers[LOGGER].(Loggable).Info(`|----------------------------------------|`)
		container.providers[LOGGER].(Loggable).Info(`| LOADING SERVICE PROVIDERS ...`)
		container.providers[LOGGER].(Loggable).Info(`|----------------------------------------|`)
	}
	container.providers[LOGGER].(Loggable).Info("| * " + key + " ✓")
}

func showBanner() {
	log := container.Log()
	log.Info(`|----------------------------------------|`)
	log.Info(`|    _____								`)
	log.Info(`|   / ____|           | |					`)
	log.Info(`|  | |  __  ___  _ __ | |__   ___ _ __	`)
	log.Info(`|  | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|	`)
	log.Info(`|  | |__| | (_) | |_) | | | |  __/ |		`)
	log.Info(`|   \_____|\___/| .__/|_| |_|\___|_|		`)
	log.Info(`|               | |						`)
	log.Info(`|               |_|						`)
	log.Info(`|----------------------------------------|`)
	log.Info(`| GOPHER READY FOR ACTION ON PORT 3000							`)
	log.Info(`|----------------------------------------|`)
}
