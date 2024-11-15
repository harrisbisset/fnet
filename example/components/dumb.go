package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var DumbPage = fnet.NewComponent("Dumb Page").
	View(dumbResponse).
	Error(0, generic.GenericBuildError).
	Build()

func DumbPageShow(ctx *fiber.Ctx) error {
	return DumbPage.RenderView(ctx)
}
