package providers

import (
	"net/http"

	"github.com/gopherlabs/gopher/vendor/_nuts/github.com/unrolled/render"
)

type RenderProvider struct {
}

func (r RenderProvider) NewRenderer() interface{} {
	return r
}

func (r RenderProvider) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	render := render.New()
	render.HTML(rw, status, name, binding)
}
