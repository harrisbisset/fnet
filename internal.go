package fnet

import (
	"context"
	"errors"
	"io"
	"net/http"
)

//
// internal types below
//

type (
	responseErr int

	none              struct{}
	optionT           interface{ any | none }
	Option[T optionT] struct{ Result T }

	internalbuildError struct{}
)

const (
	RenderFail responseErr = iota
	ErrorFail
)

func Opt[T optionT](o Option[T]) optionT {
	return *new(T)
}

func None() none {
	return *new(none)
}

func present[T comparable](V T) bool {
	return !(V == *new(T))
}

func panicField[T comparable](field T) {
	if !present(field) {
		panic("field required")
	}
}

func (b internalbuildError) Render(ctx context.Context, w io.Writer) error {
	return errors.New("<div style='margin:auto;'>404 - Page not Found</div>")
}

var buildError = NewError("build error", internalbuildError{}).Code(http.StatusInternalServerError).Build()
