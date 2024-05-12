package fnet

import (
	"context"
	"io"
	"log"
	"net/http"
)

type (
	ComponentInterface interface {
		Show(w http.ResponseWriter, req *http.Request)
	}

	Component struct {
		Name string
		View Response
		Err  Response
	}

	Response interface {
		Render(ctx context.Context, w io.Writer) error
	}
)

func (c Component) Render(w http.ResponseWriter, req *http.Request) {

	// try to render page
	err := c.View.Render(req.Context(), w)
	if err == nil {
		return
	}

	log.Printf("%s Component Failed: %s\n", c.Name, err)

	// try to render 404
	err = c.Err.Render(req.Context(), w)
	if err != nil {
		log.Printf("Error Component Failed: %s\nSomething went seriously wrong!\n", err)
	}
}
