package fnet_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/harrisbisset/fnet"
)

type DumbResponse struct{}

func (d DumbResponse) Render(ctx context.Context, w io.Writer) error {
	return errors.New("dumb response")
}

var (
	dumbResponse = DumbResponse{}
	buildError   = fnet.NewError("build error", dumbResponse).Error("error at build time").Build()

	IndexPage = fnet.NewComponent("Index Page").
			View(dumbResponse).
			Error(0, buildError).
			Build()
)

func ShowIndexPage(w http.ResponseWriter, req *http.Request) {
	IndexPage.Render(w, req)
}

func TestApplication(t *testing.T) {
	fnet.HandleComponent(fnet.GET, "/", ShowIndexPage)
	go fnet.Start("80", "0.0.0.0", nil)

	_, err := http.Get("http://localhost/")
	if err != nil {
		panic(err)
	}
}
