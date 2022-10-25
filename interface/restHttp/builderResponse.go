package restHttp

import (
	"net/http"

	"example/indexer/domain"
	"github.com/go-chi/render"
)

func NewEmailListResponse(emails []domain.Email) []render.Renderer {
	list := []render.Renderer{}
	for _, email := range emails {
		list = append(list, NewArticleResponse(email))
	}
	return list
}

type EmailResponse struct {
	domain.Email
}

func NewArticleResponse(email domain.Email) *EmailResponse {
	return &EmailResponse{Email: email}
}

func (rd *EmailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
