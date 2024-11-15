package host

import (
	"test/fnet"
	"test/sample_site/host/server_components/pages"
)

func (cfg CFG) CreateRoutes() {
	//pages
	fnet.HandleComponent(fnet.GET, "/", pages.IndexHandler)
	// fnet.HandleTest(fnet.GET, "/about", pages.)

	//fnet.HandleComponent(fnet.GET, "/about", pages.AboutHandler)
	// fnet.HandleComponent(fnet.GET, "/about", func(w http.ResponseWriter, r *http.Request) {
	// 	p := cfg.DB
	// })

}
