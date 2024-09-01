package components

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/harrisbisset/fnet"
)

type (
	DumbResponse struct{}
	NilResponse  struct{}
	BadResponse  struct{}
	DBResponse   struct{}
)

func (b BadResponse) Render(ctx context.Context, w io.Writer) error {
	v := ctx.Value("strpointer")
	return fmt.Errorf("strpointer: %s", *v.(*string))
}

func (d DBResponse) Render(ctx context.Context, w io.Writer) error {
	v := ctx.Value("data-base-result")
	return fmt.Errorf("data-base-result: %s", v)
}

func (d DumbResponse) Render(ctx context.Context, w io.Writer) error {
	return errors.New("rendered")
}

func (n NilResponse) Render(ctx context.Context, w io.Writer) error {
	return nil
}

var (
	// dumbResponse = DumbResponse{}
	dbResponse  = DBResponse{}
	badResponse = BadResponse{}
	buildError  = fnet.NewError("build error", NilResponse{}).Build()
)
