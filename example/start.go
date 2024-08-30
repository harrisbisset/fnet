package example

import (
	"github.com/harrisbisset/fnet"
)

func StartSite() {
	cidr := "0.0.0.0"
	addr := "80"

	config := NewConfig()
	config.CreateRoutes()
	fnet.Start(addr, cidr, nil)
}
