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

func PathParams(req *http.Request) map[string]string {
	return GetContainer().Parameters.PathParams(req)
}

func PathParam(req *http.Request, param string) string {
	return GetContainer().Parameters.PathParam(req, param)
}

// Logger
func NewLogger() contracts.Loggable {
	return GetContainer().Logger.NewLogger()
}

// Renderer
func NewRenderer() contracts.Renderable {
	return GetContainer().Renderer.NewRenderer()
}

func View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	GetContainer().Renderer.View(rw, status, name, binding)
}
