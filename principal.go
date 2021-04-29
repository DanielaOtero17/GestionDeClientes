package main

import (
	"net/http"

	"github.com/danielaotero17/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Get("/", services.CodeServer)
	r.Get("/{id}", services.GetUrlParam)
	r.Get("/import/", services.ImportWithoutDate)
	r.Get("/import/{date}", services.Import)
	r.Get("/buyers", services.Buyers)
	r.Get("/buyers/{id}", services.BuyersById)
	http.ListenAndServe(":8090", r)
}
