package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var BadPage = fnet.NewComponent("Bad Page").
	View(badResponse).
	Error(0, generic.GenericBuildError).
	Build()

func BadPageShow(ctx *fiber.Ctx) error {
	type tStruct struct {
		S string
	}
	str := tStruct{}

	ctx.Set("strpointer", str.S)
	return BadPage.RenderView(ctx)
}
