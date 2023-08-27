package controller

import (
	"net/http"
	"tomdog/internal/templates"
)

func registerWelcomeRoutes() {
	http.HandleFunc("/welcome", handleWelcome)
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t := templates.Template.Lookup("welcome.tmpl")
		t.Execute(w, "Welcome to go web.")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}