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

func HandleComponent(p Path, url string, fn func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(fmt.Sprintf("%s %s", p.String(), url), fn)
}

func Handle(p Path, url string, handler http.Handler) {
	http.Handle(fmt.Sprintf("%s %s", p.String(), url), handler)
}
