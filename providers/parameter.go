package providers

import (
	"net/http"

	"github.com/gopherlabs/gopher/vendor/_nuts/github.com/gorilla/mux"
)

type ParameterProvider struct{}

func (p ParameterProvider) Register(config map[string]interface{}) interface{} {
	return p
}

func (l ParameterProvider) GetKey() string {
	return "PARAMS"
}

func (p ParameterProvider) PathParams(req *http.Request) map[string]string {
	return mux.Vars(req)
}

func (p ParameterProvider) PathParam(req *http.Request, param string) string {
	return mux.Vars(req)[param]
}
