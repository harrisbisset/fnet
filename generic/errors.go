package generic

import (
	"context"
	"errors"
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
	return errors.New("<div style='margin:auto;'>404 - Page not Found</div>")
}

func (b BuildError) Render(ctx context.Context, w io.Writer) error {
	return errors.New("<div style='margin:auto;'>404 - Page not Found</div>")
}
