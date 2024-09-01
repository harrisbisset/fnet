package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet/example/components"
)

func (c cfg) CreateRoutes(app *fiber.App) {

	app.Get("/", components.IndexPageShow)
}
