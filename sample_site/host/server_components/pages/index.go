package pages

import (
	"net/http"
	"test/fnet"
	view_index "test/sample_site/render/views"
)

var IndexPage = fnet.Component{
	Name: "Index",
	View: view_index.Show(),
	Err:  view_index.Show(),
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	IndexPage.Render(w, req)
}
