package restHttp

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/render"

	"example/indexer/application"
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

type HandlerEmails struct {
	IndexMailApp application.IndexMailApp
}

func (h HandlerEmails) listEmailsByMatch(w http.ResponseWriter, r *http.Request) {
	textToSearch := chi.URLParam(r, "searchText")
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))
	// limit := r.URL.Query().Get("limit")

	emails := h.IndexMailApp.FindByMatch(textToSearch, start, 10)
	render.RenderList(w, r, NewEmailListResponse(emails))
}
