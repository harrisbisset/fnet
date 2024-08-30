package example

import (
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/example/components"
)

func (c cfg) CreateRoutes() {
	fnet.HandleComponent(fnet.GET, "/", components.IndexPageShow)
}
