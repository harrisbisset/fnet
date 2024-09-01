package example

import (
	"github.com/gofiber/fiber/v2"
)

func StartSite() {
	app := fiber.New()
	config := NewConfig()
	config.CreateRoutes(app)

	app.Listen(":80")
}
