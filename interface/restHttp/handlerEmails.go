package restHttp

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"example/indexer/application"
)

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
