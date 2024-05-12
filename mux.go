package fnet

import (
	"fmt"
	"net/http"
)

type (
	ServerMux struct {
		HTTPMux *http.ServeMux
	}
)

func CreateMux() ServerMux {
	return ServerMux{
		http.NewServeMux(),
	}
}

func (s ServerMux) HandleComponent(p Path, url string, ci ComponentInterface) {
	s.HTTPMux.HandleFunc(fmt.Sprintf("%s %s", p.String(), url), ci.Show)
}

func (s ServerMux) HandleStatic(p Path, url string, handler http.Handler) {
	s.HTTPMux.Handle(fmt.Sprintf("%s %s", p.String(), url), handler)
}

func (s ServerMux) HandleUpdate(p Path, url string, handler http.Handler) {
	s.HTTPMux.Handle(fmt.Sprintf("%s %s", p.String(), url), handler)
}
