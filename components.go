package fnet

import (
	"context"
	"io"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

type (
	Response interface {
		Render(ctx context.Context, w io.Writer) error
	}

	responseErrors map[int]respErr

	component struct {
		name   string
		view   Response
		errors responseErrors
	}
)

func (c *component) SetView(v Response) bool {
	if present(v) {
		c.view = v
		return true
	}
	log.Print("view cannot be set to a nil value")
	return false
}

func (c component) View() Response {
	return c.view
}

func (c *component) SetError(errorValue int, rerr respErr) bool {
	if !present(c.errors[errorValue]) {
		c.errors[errorValue] = rerr
		return true
	}
	log.Print("cannot override error value response")
	return false
}

func (c component) ErrorResponse(id int) Response {
	return c.errors[id].response
}

func (c component) RenderView(ctx *fiber.Ctx) error {
	switch res := c.internalRender(ctx); res {
	case nil:
		return c.RenderError(0, ctx)
	default:
		return res
	}
}

func (c component) internalRender(ctx *fiber.Ctx) error {

	// recover from error in component
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Render Fail occured on %s", c.name)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	//check if view assigned
	switch present(c.view) {
	case false:
		log.Printf("%s view not assigned", c.name)
		return c.RenderError(0, ctx)
	default:
		return renderHandler(c.view, ctx)
	}
}

func (c component) RenderError(errorValue int, ctx *fiber.Ctx) error {
	switch res := c.internalRenderError(errorValue, ctx); res {
	case nil:
		return renderHandler(buildError.response, ctx)
	default:
		return res
	}
}

func (c component) internalRenderError(errorValue int, ctx *fiber.Ctx) error {

	// recover from error in component
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error Fail occured on %s", c.name)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	//check if error assigned
	switch present(c.errors[errorValue].response) {
	default:
		return renderHandler(c.errors[errorValue].response, ctx)

	case false:
		log.Printf("%s 404 not assigned for error value %d", c.name, errorValue)
	}

	// if not default error, then display default
	if errorValue != 0 {
		return c.internalRenderError(0, ctx)
	}

	return renderHandler(buildError.response, ctx)
}
