package fnet

import (
	"net/http"

	fnettempl "github.com/harrisbisset/fnet/fnet-templ"
)

type (
	Component struct {
		Response Handler
		Handler  *fnettempl.ComponentHandler
	}

	DataComponent[T any] struct {
		Response func(v T) Handler
		Data     T
	}

	ComponentOption            func(*Component)
	DataComponentOption[T any] func(*DataComponent[T])
)

func (c Component) Render(w http.ResponseWriter, r *http.Request) {
	if c.Handler != nil {
		c.Handler.ServeHTTP(w, r)
		return
	}

	if c.Response != nil {
		c.Response(w, r)
		return
	}
}

func (c *DataComponent[T]) RenderWithData(v T) func(w http.ResponseWriter, r *http.Request) {
	c.Data = v
	return c.Response(c.Data)
}

func NewComponent(opts ...ComponentOption) Component {
	c := &Component{}
	for _, o := range opts {
		o(c)
	}
	return *c
}

func NewDataComponent[T any](opts ...DataComponentOption[T]) *DataComponent[T] {
	dc := &DataComponent[T]{}
	for _, o := range opts {
		o(dc)
	}
	return dc
}

func WithTempl(v fnettempl.Component) ComponentOption {
	return func(c *Component) {
		c.Handler = fnettempl.Handler(v)
	}
}

func WithRender(fn Handler) ComponentOption {
	return func(c *Component) {
		c.Response = fn
	}
}

func SetRender[T any](fn func(v T) Handler) DataComponentOption[T] {
	return func(dc *DataComponent[T]) {
		dc.Response = fn
	}
}
