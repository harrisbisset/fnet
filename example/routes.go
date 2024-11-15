package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet/example/components"
)

func (c cfg) CreateRoutes(app *fiber.App) {
	app.Get("/dumb", components.DumbPageShow) // should work
	app.Get("/bad", components.BadPageShow)   // should fail gracefully
	app.Get("/db", components.DBPageWrapper.AddDBConn("res").Handler)
}
