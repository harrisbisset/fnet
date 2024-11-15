package components

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/harrisbisset/fnet"
)

type (
	DumbResponse struct{}
	BadResponse  struct{}
	DBResponse   struct{}
)

func (b BadResponse) Render(ctx context.Context, w io.Writer) error {
	return errors.New("")
}

func (d DBResponse) Render(ctx context.Context, w io.Writer) error {
	buf, isBuf := w.(*bytes.Buffer)
	if !isBuf {
		buf = fnet.GetBuffer()
		defer fnet.ReleaseBuffer(buf)
	}
	_, err := buf.WriteString("dumb response")
	return err

}

func (d DumbResponse) Render(ctx context.Context, w io.Writer) error {
	buf, isBuf := w.(*bytes.Buffer)
	if !isBuf {
		buf = fnet.GetBuffer()
		defer fnet.ReleaseBuffer(buf)
	}
	_, err := buf.WriteString("dumb response")
	return err
}

var (
	dumbResponse = DumbResponse{}
	dbResponse   = DBResponse{}
	badResponse  = BadResponse{}
)
