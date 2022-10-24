package infraestructure

import "example/indexer/domain"

type EmailIteratorService struct {
	Source     emailSource
	Chunk      int
	currentIdx int
}

func (s *EmailIteratorService) NextChunk() []domain.Email {
	emailsData := s.Source.GetEmails(s.currentIdx, s.currentIdx+s.Chunk)
	emails := []domain.Email{}
	for _, emailData := range emailsData {
		emails = append(emails, buildEmail(emailData))
	}
	s.updEmailProcessed(s.Chunk)
	return emails
}

func (s *EmailIteratorService) HasNextChunk() bool {
	return s.currentIdx < s.Source.GetNumberEmails()
}

func (s *EmailIteratorService) updEmailProcessed(chunk int) {
	s.currentIdx += chunk
}

func (s *EmailIteratorService) GetPercentage() float32 {
	return float32(s.currentIdx+1) / float32(s.Source.GetNumberEmails()) * 100
}

func buildEmail(emailData map[emailField][]byte) domain.Email {
	return domain.Email{
		MessageID:               string(emailData[messageId]),
		Date:                    string(emailData[date]),
		From:                    string(emailData[from]),
		To:                      string(emailData[to]),
		Subject:                 string(emailData[subject]),
		MimeVersion:             string(emailData[mimeVersion]),
		ContentType:             string(emailData[contentType]),
		ContentTransferEncoding: string(emailData[contentTransferEncoding]),
		XFrom:                   string(emailData[xFrom]),
		XTo:                     string(emailData[xTo]),
		Xcc:                     string(emailData[xcc]),
		Xbcc:                    string(emailData[xbcc]),
		XFolder:                 string(emailData[xFolder]),
		XOrigin:                 string(emailData[xOrigin]),
		XFileName:               string(emailData[xFileName]),
		Body:                    string(emailData[body]),
	}
}
