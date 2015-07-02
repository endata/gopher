package providers

import (
	"net/http"

	"github.com/gopherlabs/gopher/vendor/_nuts/github.com/unrolled/render"
)

type RenderProvider struct {
	render *render.Render
}

func (r RenderProvider) Register(config map[string]interface{}) interface{} {
	r.render = render.New()
	return r
}

func (l RenderProvider) GetKey() string {
	return "RENDERER"
}

func (r RenderProvider) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	r.render.HTML(rw, status, name, binding)
}
