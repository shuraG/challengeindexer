package domain

type EmailIndexerService interface {
	SaveBulk(emails []Email) bool
	FindByMatch(keyword string, start, limit int) []Email
}
