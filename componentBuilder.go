package fnet

import (
	"fmt"
	"net/http"
)

type (
	ComponentBuilder component

	respErr struct {
		err      string
		response Response
		code     int
	}

	respErrBuilder respErr
)

func NewComponent(name string) *ComponentBuilder {
	com := ComponentBuilder{name: name, errors: make(responseErrors, 0)}
	return &com
}

func (cb *ComponentBuilder) View(view Response) *ComponentBuilder {
	cb.view = view
	return cb
}

func (cb *ComponentBuilder) Error(errorValue int, rerr respErr) *ComponentBuilder {
	if present(cb.errors[errorValue]) {
		panic(fmt.Sprintf("reassigned %s error response, value: %d", cb.name, errorValue))
	}
	cb.errors[errorValue] = rerr
	return cb
}

func (cb *ComponentBuilder) Build() component {
	if !present(cb.errors[0]) {
		panic(fmt.Sprintf("default error 0 not assigned to %s", cb.name))
	}
	return component(*cb)
}

func NewError(err string, resp Response) *respErrBuilder {
	return &respErrBuilder{err: err, response: resp}
}

func (r *respErrBuilder) Code(code int) *respErrBuilder {
	r.code = code
	return r
}

func (r *respErrBuilder) Error(err string) *respErrBuilder {
	r.err = err
	return r
}

func (r *respErrBuilder) Build() respErr {
	panicField(r.err)
	panicField(r.response)

	if !present(r.code) {
		r.code = http.StatusInternalServerError
	}
	return respErr(*r)
}
