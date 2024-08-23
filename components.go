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

func (c *component) View(v Response) bool {
	if present(v) {
		c.view = v
		return true
	}
	log.Print("view cannot be set to a nil value")
	return false
}

func (c *component) Error(errorValue int, rerr respErr) bool {
	c.errors[errorValue] = rerr
	return true
}

func (c component) Render(w http.ResponseWriter, req *http.Request) {
	//check if view assigned
	switch match(c.view) {
	default:
		err := c.view.Render(req.Context(), w)
		if err == nil {
			return
		}
		log.Printf("Render Fail occured on %s: %s", c.name, err)
	case false:
		unasssignedView(c.name, RenderFail)
	}
	c.internalRenderError(0, w, req)
}

func (c component) RenderError(errorValue int, w http.ResponseWriter, req *http.Request) {
	c.internalRenderError(errorValue, w, req)
}

func (c component) internalRenderError(errorValue int, w http.ResponseWriter, req *http.Request) {
	//check if error assigned
	switch match(c.view) {
	default:
		err := c.errors[errorValue].response.Render(req.Context(), w)
		if err == nil {
			log.Printf("error rendered for %s", c.name)
			http.Error(w, c.errors[errorValue].err, c.errors[errorValue].code)
			return
		}
		log.Printf("Error Fail occured on %s: %s", c.name, err)
		if errorValue == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case false:
		unasssignedView(c.name, ErrorFail)
	}
	// if not default error, then display default
	if errorValue != 0 {
		c.internalRenderError(0, w, req)
	}
}

func unasssignedView(componentName string, ty responseErr) {
	err := fmt.Sprintf("%s %s not assigned", componentName, ty.View())
	log.Print(err)
}
