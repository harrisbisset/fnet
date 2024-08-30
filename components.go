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

	responseErrors map[int]respErr

	component struct {
		name   string
		view   Response
		errors responseErrors
	}
)

type Result struct{}

func (c *component) View(v Response) bool {
	if present(v) {
		c.view = v
		return true
	}
	log.Print("view cannot be set to a nil value")
	return false
}

func (c *component) Error(errorValue int, rerr respErr) bool {
	if !present(c.errors[errorValue]) {
		c.errors[errorValue] = rerr
		return true
	}
	log.Print("cannot override error value response")
	return false
}

func (c component) Render(w http.ResponseWriter, req *http.Request) Result {
	//check if view assigned
	switch present(c.view) {
	default:
		err := c.view.Render(req.Context(), w)
		if err == nil {
			return Result{}
		}
		log.Printf("Render Fail occured on %s: %s", c.name, err)

	case false:
		err := fmt.Sprintf("%s view not assigned", c.name)
		log.Print(err)

	}
	return c.internalRenderError(0, w, req)
}

func (c component) RenderError(errorValue int, w http.ResponseWriter, req *http.Request) {
	c.internalRenderError(errorValue, w, req)
}

func (c component) internalRenderError(errorValue int, w http.ResponseWriter, req *http.Request) Result {
	//check if error assigned
	switch present(c.view) {
	default:
		err := c.errors[errorValue].response.Render(req.Context(), w)
		if err == nil {
			log.Printf("error rendered for %s", c.name)
			http.Error(w, c.errors[errorValue].err, c.errors[errorValue].code)
			return Result{}
		}
		log.Printf("Error Fail occured on %s: %s", c.name, err)
		if errorValue == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case false:
		err := fmt.Sprintf("%s 404 not assigned", c.name)
		log.Print(err)
	}

	// if not default error, then display default
	if errorValue != 0 {
		c.internalRenderError(0, w, req)
	}
	return Result{}
}
