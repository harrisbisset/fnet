package fnet

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	Response interface {
		Render(ctx context.Context, w io.Writer) error
	}

	Component struct {
		name string
		view Response
		err  Response
	}

	ComponentBuilder struct {
		c Component
	}

	ResponseErr int
)

const (
	RenderFail ResponseErr = iota
	ErrorFail
)

func (r ResponseErr) String() string {
	switch r {
	case RenderFail:
		return "Render Fail"
	default:
		return "Error Fail"
	}
}

func (r ResponseErr) View() string {
	switch r {
	case RenderFail:
		return "view"
	default:
		return "404"
	}
}

func (c *Component) View(v Response) bool {
	switch r := Ok(v); r.Outer {
	default:
		c.view = v
		return true
	case None:
		log.Print("view cannot be set to a nil value")
		return false
	}
}

func (c *Component) Error(e Response) bool {
	switch r := Ok(e); r.Outer {
	default:
		c.err = e
		return true
	case None:
		log.Print("error cannot be set to a nil value")
		return false
	}
}

func (c Component) Render(w http.ResponseWriter, req *http.Request) {
	//check if view assigned
	switch r := Ok(c.view); r.Outer {
	default:
		err := c.view.Render(req.Context(), w)
		if err == nil {
			return
		}
		log.Printf("RenderFail occured on %s: %s", c.name, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case None:
		unasssignedView(c.name, RenderFail, w)
	}
	c.RenderError(w, req)
}

func (c Component) RenderError(w http.ResponseWriter, req *http.Request) {
	//check if error assigned
	switch r := Ok(c.err); r.Outer {
	default:
		err := c.err.Render(req.Context(), w)
		if err == nil {
			log.Printf("404 rendered for %s", c.name)
			return
		}
		log.Printf("ErrorFail occured on %s: %s", c.name, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case None:
		unasssignedView(c.name, ErrorFail, w)
	}
}

func unasssignedView(componentName string, ty ResponseErr, w http.ResponseWriter) {
	err := fmt.Sprintf("%s %s not assigned", componentName, ty.View())
	log.Print(err)
	http.Error(w, err, http.StatusInternalServerError)
}

func NewComponent(name string) *ComponentBuilder {
	com := ComponentBuilder{c: Component{name: name}}
	return &com
}

func (cb *ComponentBuilder) View(view Response) *ComponentBuilder {
	cb.c.view = view
	return cb
}

func (cb *ComponentBuilder) Error(err Response) *ComponentBuilder {
	cb.c.err = err
	return cb
}

func (cb *ComponentBuilder) Build() Component {
	return cb.c
}
