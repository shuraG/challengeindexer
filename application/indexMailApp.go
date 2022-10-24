package application

import (
	"example/indexer/domain"
	"fmt"
)

type IndexMailApp struct {
	EmailIndexer domain.EmailIndexerService
	EmailService domain.EmailService
}

func (im IndexMailApp) IndexEmails() {

	for im.EmailService.HasNextChunk() {
		emails := im.EmailService.NextChunk()
		im.EmailIndexer.SaveBulk(emails)
		fmt.Println("percentage: ", im.EmailService.GetPercentage())
	}

}

func (im IndexMailApp) FindByMatch(query string, start, limit int) []domain.Email {
	return im.EmailIndexer.FindByMatch(query, start, limit)
}
