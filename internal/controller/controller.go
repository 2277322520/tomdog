package controller

import (
	"net/http"
	"tomdog/internal/config"
)

func RegisterRoutes() {
	registerIndexRoutes()
	registerWelcomeRoutes()
	registerLookRoutes()
	http.NotFoundHandler()
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(config.Config.Static))))
}
