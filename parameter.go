package gopher

import "net/http"

func PathParams(r *http.Request) map[string]string {
	return GetContext().Parameters.PathParams(r)
}

func PathParam(r *http.Request, param string) string {
	return GetContext().Parameters.PathParam(r, param)
}
