package infraestructure

import (
	"example/indexer/domain"
)

type emailBulkRequest struct {
	Index   string         `json:"index"`
	Records []domain.Email `json:"records"`
}

type EmailZincService struct {
	DefaultIndex string
	HttpClient   HttpClient
}

const saveBulkEndpoint = "api/_bulkv2"

func (s EmailZincService) SaveBulk(emails []domain.Email) bool {

	return s.HttpClient.SaveBulk(emails, s.DefaultIndex)
}

const findByMatch = "api/%s/_search"

func (s EmailZincService) FindByMatch(keyword string, start, limit int) []domain.Email {

	response := s.HttpClient.FindByMatch(s.DefaultIndex, keyword, start, limit)

	return buildEmailsFromZincRes(response)
}

func buildEmailsFromZincRes(data EmailZincResponse) []domain.Email {
	emails := make([]domain.Email, len(data.Hits.Hits))
	for i, hit := range data.Hits.Hits {
		emails[i] = domain.Email{
			MessageID:               hit.ID,
			Date:                    hit.Source.Date,
			From:                    hit.Source.From,
			To:                      hit.Source.To,
			Subject:                 hit.Source.Subject,
			MimeVersion:             hit.Source.MimeVersion,
			ContentType:             hit.Source.ContentType,
			ContentTransferEncoding: hit.Source.ContentTransferEncoding,
			XFrom:                   hit.Source.XFrom,
			XTo:                     hit.Source.XTo,
			Xcc:                     hit.Source.Xcc,
			Xbcc:                    hit.Source.Xbcc,
			XFolder:                 hit.Source.XFolder,
			XOrigin:                 hit.Source.XOrigin,
			XFileName:               hit.Source.XFileName,
			Body:                    hit.Source.Body,
		}
	}
	return emails
}
