package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(badResponse).
	Error(0, buildError).
	Build()

func IndexPageShow(ctx *fiber.Ctx) error {
	type tStruct struct {
		S string
	}
	str := tStruct{}

	ctx.Set("strpointer", str.S)
	return IndexPage.Render(ctx)
}
