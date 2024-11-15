package pages

import (
	"net/http"
	"test/fnet"
	view_index "test/sample_site/render/views"
)

var NewPage = fnet.Component{
	Name: "New",
	View: view_index.Show(),
	Err:  view_index.Show(),
}

type NP struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	Function func(w http.ResponseWriter, r *http.Request)
}

func CreateNP(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request)) NP {
	return NP{
		Writer:   w,
		Request:  r,
		Function: f,
	}
}

func (np NP) Result() func(w http.ResponseWriter, r *http.Request) {
	return np.Function
}

func (np NP) Wr() http.ResponseWriter {
	return np.Writer
}

func (np NP) Rq() *http.Request {
	return np.Request
}
