package fnet

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func renderHandler(component Response, c *fiber.Ctx) error {
	handler := adaptor.HTTPHandler(&templ.ComponentHandler{
		Component:   component,
		ContentType: "text/html; charset=utf-8",
	})
	return handler(c)
}
