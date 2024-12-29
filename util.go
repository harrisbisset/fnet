package fnet

import (
	"net/http"

	fnettempl "github.com/harrisbisset/fnet/fnet-templ"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func RenderTempl(v fnettempl.Component, w http.ResponseWriter, r *http.Request, opts ...func(*fnettempl.ComponentHandler)) {
	fnettempl.Handler(v, opts...).ServeHTTP(w, r)
}
