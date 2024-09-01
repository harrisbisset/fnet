package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(dumbResponse).
	Error(0, buildError).
	Build()

func IndexPageShow(ctx *fiber.Ctx) error {
	return IndexPage.Render(ctx)
}
