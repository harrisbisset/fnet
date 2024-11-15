package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var DBPage = fnet.NewComponent("DB Page").
	View(dbResponse).
	Error(0, generic.GenericBuildError).
	Build()

type DBPageWrapperStruct struct {
	DB string
}

var DBPageWrapper DBPageWrapperStruct

func (w *DBPageWrapperStruct) AddDBConn(s string) *DBPageWrapperStruct {
	w.DB = s
	return w
}

func (w *DBPageWrapperStruct) Handler(ctx *fiber.Ctx) error {

	// example
	ctx.Locals("data-base-result", w.DB)
	ctx.Next()

	return DBPage.RenderView(ctx)
}
