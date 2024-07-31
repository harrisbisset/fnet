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
	default:
		return "Render Fail"
	case ErrorFail:
		return "Error Fail"
	}
}

func (c Component) View(v Response) bool {
	switch r := Ok(v); r.Outer {
	default:
		return true
	case None:
		log.Print("view cannot be set to a nil value")
		return false
	}
}

func (c Component) Error(e Response) bool {
	switch r := Ok(e); r.Outer {
	default:
		return true
	case None:
		log.Print("view cannot be set to a nil value")
		return false
	}
}

func (c Component) Render(w http.ResponseWriter, req *http.Request) {

	//check if view assigned
	switch r := Ok(c.view); r.Outer {
	default:

		// try to render page
		err := c.view.Render(req.Context(), w)
		if err == nil {
			log.Printf("Rendered %s", c.name)
			return
		}
		log.Printf("%s occured on %s: %s", RenderFail.String(), c.name, err)

	case None:
		err := fmt.Sprintf("%s view not assigned", c.name)
		log.Print(err)
		http.Error(w, err, http.StatusInternalServerError)
	}

	c.RenderError(w, req)
}

func (c Component) RenderError(w http.ResponseWriter, req *http.Request) {
	//check if error assigned
	switch r := Ok(c.err); r.Outer {
	default:

		// try to render 404
		err := c.err.Render(req.Context(), w)
		if err == nil {
			log.Printf("404 rendered for %s", c.name)
			return
		}
		log.Printf("%s occured on %s: %s", ErrorFail.String(), c.name, err)

	case None:
		err := fmt.Sprintf("%s error not assigned", c.name)
		log.Print(err)
		http.Error(w, err, http.StatusInternalServerError)
	}
}

func NewComponent(name string) ComponentBuilder {
	return ComponentBuilder{c: Component{name: name}}
}

func (cb *ComponentBuilder) View(view Response) ComponentBuilder {
	cb.c.view = view
	return *cb
}

func (cb *ComponentBuilder) Error(err Response) ComponentBuilder {
	cb.c.err = err
	return *cb
}

func (cb ComponentBuilder) Build(view Response) Component {
	return cb.c
}
