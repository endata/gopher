package gopher

import (
	"github.com/gopherlabs/gopher-services"
)

func RegisterProviders() {

	GetContext().Logger = services.Logger2{}
}
