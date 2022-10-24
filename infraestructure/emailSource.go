package infraestructure

type emailSource interface {
	GetEmails(start, finish int) []map[emailField][]byte
	GetNumberEmails() int
}
