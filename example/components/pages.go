package components

import (
	fnet "github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/example/render/views/view_index"
)

var IndexPage = fnet.NewComponent(fnet.WithTempl(view_index.Show()))

var DatabasePage = fnet.NewComponent()
