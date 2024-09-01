package fnet

import (
	"context"
	"fmt"
	"io"
	"log"

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
	//check if view assigned
	switch present(c.view) {
	default:
		err := c.view.Render(ctx.Context(), ctx.Response().BodyWriter())
		if err == nil {
			return nil
		}
		log.Printf("Render Fail occured on %s: %s", c.name, err)

	case false:
		err := fmt.Sprintf("%s view not assigned", c.name)
		log.Print(err)

	}
	return c.internalRenderError(0, ctx)
}

func (c component) RenderError(errorValue int, ctx *fiber.Ctx) error {
	return c.internalRenderError(errorValue, ctx)
}

func (c component) internalRenderError(errorValue int, ctx *fiber.Ctx) error {
	var err error

	//check if error assigned
	switch present(c.errors[errorValue].response) {
	default:
		err = c.errors[errorValue].response.Render(ctx.Context(), ctx.Response().BodyWriter())
		if err == nil {
			log.Printf("error rendered for %s", c.name)
			return nil
		}
		log.Printf("Error Fail occured on %s: %s", c.name, err)
	case false:
		err = fmt.Errorf("%s 404 not assigned", c.name)
		log.Print(err)
	}

	// if not default error, then display default
	if errorValue != 0 {
		return c.internalRenderError(0, ctx)
	}
	return err
}
