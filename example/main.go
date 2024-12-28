package main

import (
	"net/http"

	"github.com/harrisbisset/fnet/example/components"
)

func main() {
	http.HandleFunc("/", components.IndexPage.Render)
	http.HandleFunc("/db", components.DatabasePage.Render)

	http.ListenAndServe(":8080", nil)
}
