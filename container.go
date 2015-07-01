package gopher

import (
	"fmt"

	"github.com/gopherlabs/gopher/providers"
)

var container = appContainer{}

type Provider int

const (
	PROVIDER_LOGGER Provider = iota
	PROVIDER_ROUTER
	PROVIDER_RENDERER
	PROVIDER_PARAMS
)

func (p Provider) String() string {
	key := ""
	switch p {
	case PROVIDER_LOGGER:
		key = "Logger"
	case PROVIDER_ROUTER:
		key = "Router"
	case PROVIDER_RENDERER:
		key = "Renderer"
	case PROVIDER_PARAMS:
		key = "Params"
	}
	return fmt.Sprintf(key)
}

type appContainer struct {
	logger     Providerable
	router     Routable
	parameters Parametable
	renderer   Renderable
}

func App() *appContainer {
	registerProviders()
	showBanner()
	return &container
}

func registerProviders() {
	RegisterProvider(PROVIDER_LOGGER, providers.LogProvider{})
	RegisterProvider(PROVIDER_ROUTER, providers.RouteProvider{})
	RegisterProvider(PROVIDER_RENDERER, providers.RenderProvider{})
	RegisterProvider(PROVIDER_PARAMS, providers.ParameterProvider{})
}

func RegisterProvider(key Provider, provider interface{}) {
	switch key {
	case PROVIDER_LOGGER:
		container.logger = provider.(Providerable).Register().(Providerable)
		container.logger.(Loggable).Info(" * " + key.String() + " ✓")
	case PROVIDER_ROUTER:
		container.router = provider.(Routable)
	case PROVIDER_RENDERER:
		container.renderer = provider.(Renderable)
	case PROVIDER_PARAMS:
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
