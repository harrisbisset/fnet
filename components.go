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
		Name string
		View Response
		Err  Response
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
		return "RenderFail"
	case ErrorFail:
		return "ErrorFail"
	}
}

func (c Component) Render(w http.ResponseWriter, req *http.Request) {

	// try to render page
	err := c.View.Render(req.Context(), w)
	if err == nil {
		log.Printf("Rendered %s", c.Name)
		return
	}
	log.Printf("%s occured on %s: %s", RenderFail.String(), c.Name, err)

	// try to render 404
	err = c.Err.Render(req.Context(), w)
	if err == nil {
		log.Printf("404 Rendered %s", c.Name)
		return
	}

	log.Printf("%s occured on %s: %s", ErrorFail.String(), c.Name, err)
}

func HandleComponent(p Path, url string, com func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(fmt.Sprintf("%s %s", p.String(), url), com)
}

func Handle(p Path, url string, handler http.Handler) {
	http.Handle(fmt.Sprintf("%s %s", p.String(), url), handler)
}
