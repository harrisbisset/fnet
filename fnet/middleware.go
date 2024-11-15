package fnet

import (
	"net/http"
	"strings"
)

func NeuterFileSystem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// if path is empty, or if it ends in /
		// then it's a dir
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
