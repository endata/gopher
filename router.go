package gopher

import (
	"github.com/gopherlabs/gopher/contracts"
)

func NewRouter() contracts.Routable {
	return GetContext().Router.NewRouter()
}
