package fnet

import (
	"fmt"
	"net/http"
)

type (
	ComponentBuilder struct {
		c component
	}

	respErr struct {
		name     string
		response Response
		code     int
		err      string
	}

	RErrBuilder respErr
)

func NewComponent(name string) *ComponentBuilder {
	com := ComponentBuilder{c: component{name: name, errors: make(responseErrors, 0)}}
	return &com
}

func (cb *ComponentBuilder) View(view Response) *ComponentBuilder {
	cb.c.view = view
	return cb
}

func (cb *ComponentBuilder) Error(errorValue int, rerr RErrBuilder) *ComponentBuilder {
	if present(cb.c.errors[errorValue]) {
		panic(fmt.Sprintf("reassigned %s error response, value: %d", cb.c.name, errorValue))
	}
	cb.c.errors[errorValue] = checkRErr(rerr)
	return cb
}

func (cb *ComponentBuilder) Build() component {
	if !present(cb.c.errors[0]) {
		panic(fmt.Sprintf("default error 0 not assigned to %s", cb.c.name))
	}
	return cb.c
}

func checkRErr(builder RErrBuilder) respErr {
	panicField(builder.name)
	panicField(builder.err)
	panicField(builder.response)

	if !checkField(builder.code) {
		builder.code = http.StatusInternalServerError
	}
	return respErr(builder)
}
