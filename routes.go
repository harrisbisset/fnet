package fnet

import (
	"fmt"
	"net/http"
)

type (
	Path int
)

const (
	GET Path = iota
	POST
	UPDATE
	DELETE
	PUT
)

func (p Path) String() string {
	switch p {
	// assume GET
	default:
		return "GET"
	case 1:
		return "POST"
	case 2:
		return "UPDATE"
	case 3:
		return "DELETE"
	case 4:
		return "PUT"
	}
}

func HandleComponent(p Path, url string, fn func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(fmt.Sprintf("%s %s", p.String(), url), fn)
}

func Handle(p Path, url string, handler http.Handler) {
	http.Handle(fmt.Sprintf("%s %s", p.String(), url), handler)
}
