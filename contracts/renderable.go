package contracts

import "net/http"

type Renderable interface {
	NewRenderer() Renderable

	View(rw http.ResponseWriter, status int, name string, binding interface{})
}
