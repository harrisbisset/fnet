package pages

import (
	"database/sql"
	"net/http"
	"test/fnet"
	view_index "test/sample_site/render/views"
)

var AboutPage = fnet.Component{
	Name: "Index",
	View: view_index.Show(),
	Err:  view_index.Show(),
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	AboutPage.Render(w, req)
}

func GetAuth(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return AboutHandler
}
