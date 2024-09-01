package generic

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/harrisbisset/fnet"
)

type BuildError struct{}

func (b BuildError) Render(ctx context.Context, w io.Writer) error {
	return errors.New("<div style='margin:auto;'>404 - Page not Found</div>")
}

var GenericBuildError = fnet.NewError("build error", BuildError{}).Code(http.StatusInternalServerError).Build()
