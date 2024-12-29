package components

import (
	"net/http"

	fnet "github.com/harrisbisset/fnet"

	"github.com/harrisbisset/fnet/example/models"
	"github.com/harrisbisset/fnet/example/render/views/view_database"
	"github.com/harrisbisset/fnet/example/render/views/view_index"
)

var IndexPage = fnet.NewComponent(fnet.WithTempl(view_index.Show()))

var DatabasePage = fnet.NewDataComponent(fnet.SetRender(
	func(v models.DatabaseRecord) fnet.Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			fnet.RenderTempl(view_database.Hello(v), w, r)
		}
	},
))
