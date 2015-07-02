package gopher

import (
	"fmt"

	"github.com/gopherlabs/gopher/providers"
)

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
	config     map[string]map[string]interface{}
	logger     Providerable
	router     Routable
	parameters Parametable
	renderer   Renderable
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
	registerProvider(LOGGER, providers.LogProvider{})
	//	registerProvider(PROVIDER_ROUTER, providers.RouteProvider{})
	//	registerProvider(PROVIDER_RENDERER, providers.RenderProvider{})
	//	registerProvider(PROVIDER_PARAMS, providers.ParameterProvider{})
}

func registerProvider(key string, provider interface{}) {
	switch key {
	case LOGGER:
		fmt.Println(container.config)
		container.logger = provider.(Providerable).Register(defaultConfig[LOGGER]).(Providerable)
		container.logger.(Loggable).Info(" * " + key + " ✓" + " has key of " + provider.(Providerable).GetKey())
	case ROUTER:
		container.router = provider.(Routable)
	case RENDERER:
		container.renderer = provider.(Renderable)
	case PARAMS:
		container.parameters = provider.(Parametable)
	}
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
