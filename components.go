package fnet

import (
	"context"
	"errors"
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

type Result struct{}

func (c *component) View(v Response) bool {
	if present(v) {
		c.view = v
		return true
	}
	log.Print("view cannot be set to a nil value")
	return false
}

func (c *component) Error(errorValue int, rerr respErr) bool {
	if !present(c.errors[errorValue]) {
		c.errors[errorValue] = rerr
		return true
	}
	log.Print("cannot override error value response")
	return false
}

func (c component) Render(ctx *fiber.Ctx) error {
	err := c.internalRender(ctx)
	if err == nil {
		return c.RenderError(0, ctx)
	}
	return err

}

func (c component) internalRender(ctx *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Render Fail occured on %s", c.name)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	//check if view assigned
	switch present(c.view) {
	default:
		return c.view.Render(ctx.Context(), ctx.Response().BodyWriter())
	case false:
		log.Printf("%s view not assigned", c.name)
		return c.RenderError(0, ctx)
	}
}

func (c component) RenderError(errorValue int, ctx *fiber.Ctx) error {
	err := c.internalRenderError(errorValue, ctx)
	if err == nil {
		return errors.New("request failed")
	}
	return err
}

func (c component) internalRenderError(errorValue int, ctx *fiber.Ctx) error {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error Fail occured on %s", c.name)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	//check if error assigned
	switch present(c.errors[errorValue].response) {
	default:
		return c.errors[errorValue].response.Render(ctx.Context(), ctx.Response().BodyWriter())
	case false:
		log.Printf("%s 404 not assigned for %d", c.name, errorValue)
	}

	// if not default error, then display default
	if errorValue != 0 {
		return c.internalRenderError(0, ctx)
	}

	return errors.New("request failed")
}
