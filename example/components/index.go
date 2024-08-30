package components

import (
	"net/http"

	"github.com/harrisbisset/fnet"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(dumbResponse).
	Error(0, buildError).
	Build()

func IndexPageShow(w http.ResponseWriter, req *http.Request) fnet.Result {
	return IndexPage.Render(w, req)
}
