package gopher

import (
	"github.com/gopherlabs/gopher-services/providers"
)

func RegisterProviders() {
	GetContext().Logger = providers.Logger{}
	GetContext().Router = providers.RouteProvider{}
}
