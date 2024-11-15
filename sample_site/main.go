package main

import (
	"test/fnet"
	"test/sample_site/host"
)

func main() {
	cfg := host.CreateConfig()
	cfg.CreateRoutes()
	fnet.Start()
}
