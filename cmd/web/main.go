package main

import (
	"go_module/routes"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/",routes.Home)
	mux.Get("/about",routes.About)
	mux.Get("/rooms/general", routes.General)
	mux.Get("/rooms/major",routes.Major)
	mux.Get("/contact",routes.Contact)
	mux.Get("/data",routes.Data)
	mux.Post("/data",routes.Booking)
	mux.Post("/contact",routes.PostQuery)

	http.ListenAndServe(":8080",mux)
}