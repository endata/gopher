package gopher

import (
	"net/http"

	"github.com/gopherlabs/gopher/contracts"
)

// Routers

func (c container) Router() contracts.Routable {
	return getContainer().router.Router()
}

// Parameters

func PathParams(req *http.Request) map[string]string {
	return getContainer().parameters.PathParams(req)
}

func PathParam(req *http.Request, param string) string {
	return getContainer().parameters.PathParam(req, param)
}

// Logger
func (c container) Log() contracts.Loggable {
	return getContainer().logger.Log()
}

// Renderer
func NewRenderer() contracts.Renderable {
	return getContainer().renderer.NewRenderer()
}

func View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	getContainer().renderer.View(rw, status, name, binding)
}
