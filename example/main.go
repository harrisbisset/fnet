package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet/example/components"
)

func main() {
	app := fiber.New()

	app.Get("/dumb", components.DumbPageShow) // should work
	app.Get("/bad", components.BadPageShow)   // should fail gracefully
	app.Get("/db", components.DBPageWrapper.AddDBConn("res").Handler)

	app.Listen(":80")

}
