package gopher

import (
	"net/http"

	"github.com/gopherlabs/gopher/contracts"
)

// Routers

func Router() contracts.Routable {
	return getContainer().Router.Router()
}

// Parameters

func PathParams(req *http.Request) map[string]string {
	return getContainer().Parameters.PathParams(req)
}

func PathParam(req *http.Request, param string) string {
	return getContainer().Parameters.PathParam(req, param)
}

// Logger
func NewLogger() contracts.Loggable {
	return getContainer().Logger.NewLogger()
}

// Renderer
func NewRenderer() contracts.Renderable {
	return getContainer().Renderer.NewRenderer()
}

func View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	getContainer().Renderer.View(rw, status, name, binding)
}
