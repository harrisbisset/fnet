package fnet

import (
	"net/http"

	fnettempl "github.com/harrisbisset/fnet/fnet-templ"
)

type (
	Component struct {
		ContentType string
		Response    func(w http.ResponseWriter, r *http.Request)

		Handler *fnettempl.ComponentHandler
	}

	ComponentOption func(*Component)
)

func (c Component) Render(w http.ResponseWriter, r *http.Request) {
	if c.Handler == nil {
		c.Handler.ServeHTTP(w, r)
		return
	}

	if c.

	c.Response(w, r)
}

func NewComponent(opts ...ComponentOption) Component {
	c := &Component{
		ContentType: "text/html; charset=utf-8",
	}

	for _, o := range opts {
		o(c)
	}

	return *c
}

func WithContentType(t string) ComponentOption {
	return func(c *Component) {
		c.ContentType = t
	}
}

func WithTempl(v fnettempl.Component, opts ...func(*fnettempl.ComponentHandler)) ComponentOption {
	return func(c *Component) {
		c.Handler = fnettempl.Handler(v)
	}
}

func WithRender(fn func(w http.ResponseWriter, r *http.Request)) ComponentOption {
	return func(c *Component) {

	}
}
