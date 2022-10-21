package infraestructure

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
)

type dirReader func(rootPath string) ([]fs.DirEntry, error)

type EmailFileSource struct {
	filePaths []string
}

func NewEmailFileSource(rootPath string) *EmailFileSource {
	if rootPath == "" {
		return nil
	}
	fSource := new(EmailFileSource)
	fSource.loadFileNames(rootPath, os.ReadDir)
	return fSource
}

func (s *EmailFileSource) loadFileNames(rootPath string, read dirReader) {
	files, err := read(rootPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			s.loadFileNames(rootPath+"\\"+file.Name(), read)
		} else {
			s.filePaths = append(s.filePaths, rootPath+"\\"+file.Name())
		}
	}
}

func (s *EmailFileSource) GetEmails(start, finish int) []map[emailField][]byte {
	var emailsData []map[emailField][]byte

	for i := start; i < s.getLimitFiles(finish); i++ {
		email := getEmailData(s.filePaths[i])
		emailsData = append(emailsData, email)
	}

	return emailsData
}

func (s *EmailFileSource) GetNumberEmails() int {
	return len(s.filePaths)
}

func (s *EmailFileSource) getLimitFiles(j int) int {
	sizeFiles := len(s.filePaths)
	if j < sizeFiles {
		return j
	}
	return sizeFiles
}

func getEmailData(fileName string) map[emailField][]byte {
	data := getDataFromFile(fileName)
	return parseEmailData(defaultFields, data)
}

func getDataFromFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("unable to read file: %v", err)
	}
	return data
}

func parseEmailData(fields []emailField, data []byte) map[emailField][]byte {
	emailDataByField := make(map[emailField][]byte)
	for _, field := range fields {
		iField := bytes.Index(data, []byte(field))
		iData := bytes.Index(data[iField:], []byte("\n")) + iField
		if field == body || iData == -1 {
			iData = len(data)
		}
		emailDataByField[field] = bytes.TrimSpace(data[iField+len(field) : iData])
	}

	return emailDataByField
}

type emailField string

const (
	messageId               emailField = "Message-ID:"
	date                    emailField = "Date:"
	from                    emailField = "From:"
	to                      emailField = "To:"
	subject                 emailField = "Subject:"
	mimeVersion             emailField = "Mime-Version:"
	contentType             emailField = "Content-Type:"
	contentTransferEncoding emailField = "Content-Transfer-Encoding:"
	xFrom                   emailField = "X-From:"
	xTo                     emailField = "X-To:"
	xcc                     emailField = "X-cc:"
	xbcc                    emailField = "X-bcc:"
	xFolder                 emailField = "X-Folder:"
	xOrigin                 emailField = "X-Origin:"
	xFileName               emailField = "X-FileName:"
	body                    emailField = "\r\n\r\n"
)

var defaultFields = []emailField{messageId, date, from, to, subject, mimeVersion,
	contentType, contentTransferEncoding, xFrom, xTo, xcc, xbcc, xFolder, xOrigin, xFileName,
	body}
