package gopher

import (
	"github.com/gopherlabs/gopher/contracts"
	"net/http"
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
