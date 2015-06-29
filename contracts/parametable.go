package contracts

import "net/http"

type Parametable interface {
	PathParams(r *http.Request) map[string]string

	PathParam(r *http.Request, param string) string
}
