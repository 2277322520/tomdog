package middleware

import (
	"net/http"
	"tomdog/internal/log"
)

type LogMiddleWare struct {
	Next http.Handler
}

func (m LogMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Next == nil {
		m.Next = http.DefaultServeMux
	}
	log.Info.Println("[Request]\t- Method: " + r.Method + "\t\tURL: " + r.URL.Path)
	m.Next.ServeHTTP(w, r)
}
