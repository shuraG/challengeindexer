package restHttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitServer(port string, handlerEmails HandlerEmails) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*", "*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}}))

	r.Route("/emails", func(r chi.Router) {
		r.Get("/search/{searchText}", handlerEmails.listEmailsByMatch)
	})

	http.ListenAndServe(":3000", r)
}
