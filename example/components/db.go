package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/fnet"
)

var DBPage = fnet.NewComponent("DB Page").
	View(dbResponse).
	Error(0, buildError).
	Build()

type WrappedDBPageStruct struct {
	DB string
}

var WrappedDBPage WrappedDBPageStruct

func (w *WrappedDBPageStruct) AddDBConn(s string) *WrappedDBPageStruct {
	w.DB = s
	return w
}

func (w *WrappedDBPageStruct) DBPageShow(ctx *fiber.Ctx) error {

	// example
	ctx.Locals("data-base-result", w.DB)
	ctx.Next()

	return DBPage.RenderView(ctx)
}
