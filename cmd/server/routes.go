package main

import (
	"github.com/mdhender/scheme/pkg/way"
	"net/http"
)

func routes(srv *server, cfg *Config) http.Handler {
	router := way.NewRouter()
	return router
}
