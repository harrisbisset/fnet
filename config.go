package fnet

import (
	"os"
)

type (
	CFG struct {
		Mux  ServerMux
		Port string
	}
)

func CreateConfig() CFG {
	return CFG{
		Mux:  CreateMux(),
		Port: getPort(),
	}
}

func getPort() string {
	port := os.Getenv("PORT")

	switch len(port) {
	case 0:
		port = ":3000"
	default:
		port = "0.0.0.0:" + port
	}

	return port
}
