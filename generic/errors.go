package generic

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/harrisbisset/fnet"
)

type (
	DefaultError struct{}
	BuildError   struct{}
)

var (
	GenericBuildError    = fnet.NewError("build error", BuildError{}).Code(http.StatusInternalServerError).Build()
	DefaultResponseError = fnet.NewError("default error", BuildError{}).Code(http.StatusInternalServerError).Build()
)

func (d DefaultError) Render(ctx context.Context, w io.Writer) error {
	buf, isBuf := w.(*bytes.Buffer)
	if !isBuf {
		buf = fnet.GetBuffer()
		defer fnet.ReleaseBuffer(buf)
	}
	_, err := buf.WriteString("<div style='margin:auto;'>404 - Not Found, default error</div>")
	return err
}

func (b BuildError) Render(ctx context.Context, w io.Writer) error {
	buf, isBuf := w.(*bytes.Buffer)
	if !isBuf {
		buf = fnet.GetBuffer()
		defer fnet.ReleaseBuffer(buf)
	}
	_, err := buf.WriteString("<div style='margin:auto;'>404 - Not Found, error occured at build time</div>")
	return err
}
