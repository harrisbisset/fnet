package components

import (
	"net/http"

	"test/fnet"
)

type DefaultWrapper fnet.Component

func (a DefaultWrapper) Show(w http.ResponseWriter, req *http.Request) {
	fnet.Component(a).Render(w, req)
}
