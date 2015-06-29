package contracts

import "net/http"

type Renderable interface {
	NewRenderer() Renderable
	View(w http.ResponseWriter, status int, name string, binding interface{})
}
