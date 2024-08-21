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

	responseError struct {
		name     string
		response Response
	}

	responseErrors map[int]responseError

	component struct {
		name   string
		view   Response
		errors responseErrors
	}

	ComponentBuilder struct {
		c component
	}

	responseErr int
)

const (
	RenderFail responseErr = iota
	ErrorFail
)

func (r responseErr) String() string {
	switch r {
	case RenderFail:
		return "Render Fail"
	default:
		return "Error Fail"
	}
}

func (r responseErr) View() string {
	switch r {
	case RenderFail:
		return "view"
	default:
		return "404"
	}
}

func (c *component) View(v Response) bool {
	switch r := Ok(v); r.Result {
	default:
		c.view = v
		return true
	case None:
		log.Print("view cannot be set to a nil value")
		return false
	}
}

func (c *component) Error(errorName string, errorValue int, e Response) bool {
	switch r := Ok(e); r.Result {
	default:
		c.errors[errorValue] = responseError{name: errorName, response: e}
		return true
	case None:
		log.Print("error cannot be set to a nil value")
		return false
	}
}

func (c component) Render(w http.ResponseWriter, req *http.Request) {
	//check if view assigned
	switch r := Ok(c.view); r.Result {
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
	c.internalRenderError(0, w, req)
}

func (c component) RenderError(errorValue int, w http.ResponseWriter, req *http.Request) {
	c.internalRenderError(errorValue, w, req)
}

func (c component) internalRenderError(errorValue int, w http.ResponseWriter, req *http.Request) {
	//check if error assigned
	switch r := Ok(c.errors[errorValue]); r.Result {
	default:
		err := c.errors[errorValue].response.Render(req.Context(), w)
		if err == nil {
			log.Printf("404 rendered for %s", c.name)
			return
		}
		log.Printf("ErrorFail occured on %s: %s", c.name, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case None:
		unasssignedView(c.name, ErrorFail, w)
	}
	// if not default error, then display
	if errorValue != 0 {
		c.internalRenderError(0, w, req)
	}
}

func unasssignedView(componentName string, ty responseErr, w http.ResponseWriter) {
	err := fmt.Sprintf("%s %s not assigned", componentName, ty.View())
	log.Print(err)
	http.Error(w, err, http.StatusInternalServerError)
}

func NewComponent(name string) *ComponentBuilder {
	com := ComponentBuilder{c: component{name: name, errors: make(responseErrors, 0)}}
	return &com
}

func (cb *ComponentBuilder) View(view Response) *ComponentBuilder {
	cb.c.view = view
	return cb
}

func (cb *ComponentBuilder) Error(errorValue int, rerr responseError) *ComponentBuilder {
	if !Present(cb.c.errors[errorValue]) {
		panic(fmt.Sprintf("reassigned %s error response, value: %d", cb.c.name, errorValue))
	}
	cb.c.errors[errorValue] = rerr
	return cb
}

func (cb *ComponentBuilder) Build() component {
	if !Present(cb.c.errors[0]) {
		panic(fmt.Sprintf("default error 0 not assigned to %s", cb.c.name))
	}
	return cb.c
}

func RespError(name string, err Response) responseError {
	return responseError{name: name, response: err}
}
