package main

import (
	"golang.org/x/exp/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello there"))
	})
	slog.Info("booting up server", "port", port)
	http.ListenAndServe(port, r)
}
