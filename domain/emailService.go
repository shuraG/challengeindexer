package domain

type EmailService interface {
	NextChunk() []Email
	HasNextChunk() bool
	GetPercentage() float32
}
