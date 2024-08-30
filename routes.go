package fnet

import (
	"fmt"
	"net/http"
)

type (
	Path string
)

const (
	GET    Path = "GET"
	POST   Path = "POST"
	UPDATE Path = "UPDATE"
	DELETE Path = "DELETE"
	PUT    Path = "PUT"
)

func HandleComponent(p Path, url string, fn func(w http.ResponseWriter, req *http.Request) Result) {
	http.HandleFunc(fmt.Sprintf("%s %s", p, url), func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	})
}

func Handle(p Path, url string, handler http.Handler) {
	http.Handle(fmt.Sprintf("%s %s", p, url), handler)
}
