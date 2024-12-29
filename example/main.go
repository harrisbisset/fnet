package main

import (
	"log"
	"net/http"

	"github.com/harrisbisset/fnet/example/components"
	"github.com/harrisbisset/fnet/example/models"
)

func main() {
	db := models.DatabaseRecord{
		Name: "Database Record Name",
	}

	http.HandleFunc("/", components.IndexPage.Render)
	http.HandleFunc("/db", components.DatabasePage.RenderWithData(db))

	log.Print("listening")
	http.ListenAndServe(":8080", nil)
}
