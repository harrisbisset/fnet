package components

import (
	"context"
	"errors"
	"io"

	"github.com/harrisbisset/fnet"
)

type DumbResponse struct{}

func (d DumbResponse) Render(ctx context.Context, w io.Writer) error {
	return errors.New("dumb response")
}

var (
	dumbResponse = DumbResponse{}
	buildError   = fnet.NewError("build error", dumbResponse).Build()
)
