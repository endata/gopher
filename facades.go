package gopher

import (
	"net/http"

	"github.com/gopherlabs/gopher/contracts"
)

// Routers

func NewRouter() contracts.Routable {
	return GetContainer().Router.NewRouter()
}

// Parameters

func PathParams(r *http.Request) map[string]string {
	return GetContainer().Parameters.PathParams(r)
}

func PathParam(r *http.Request, param string) string {
	return GetContainer().Parameters.PathParam(r, param)
}

// Logger
func NewLogger() contracts.Loggable {
	return GetContainer().Logger.NewLogger()
}

// Renderer
func NewRenderer() contracts.Renderable {
	return GetContainer().Renderer.NewRenderer()
}

func View(w http.ResponseWriter, status int, name string, binding interface{}) {
	GetContainer().Renderer.View(w, status, name, binding)
}
