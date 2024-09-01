package components

import (
	"context"
	"fmt"
	"io"

	"github.com/harrisbisset/fnet"
)

type DumbResponse struct{}

func (d DumbResponse) Render(ctx context.Context, w io.Writer) error {
	v := ctx.Value("strpointer")
	return fmt.Errorf("strpointer: %s", *v.(*string))
}

var (
	dumbResponse = DumbResponse{}
	buildError   = fnet.NewError("build error", dumbResponse).Build()
)
