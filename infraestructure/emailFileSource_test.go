package infraestructure

import (
	"io/fs"
	"testing"
)

type stubDir struct {
	name  string
	isDir bool
}

func (stub stubDir) Name() string {
	return stub.name
}
func (stub stubDir) IsDir() bool {
	return stub.isDir
}
func (stub stubDir) Type() fs.FileMode {
	return 0
}
func (stub stubDir) Info() (fs.FileInfo, error) {
	return nil, nil
}

func TestLoadFileNames(t *testing.T) {
	emailSource := new(EmailFileSource)
	emailSource.loadFileNames("testPath", mockDirReader)

	expected := "testPath\\secondLevel\\thirdLevel.exe"

	if emailSource.filePaths[0] != expected {
		t.Errorf("output %q, expected %q", emailSource.filePaths[0], expected)
	}

	if len(emailSource.filePaths) != 1 {
		t.Errorf("output %q, expected %q", len(emailSource.filePaths), 1)
	}

}

func mockDirReader(rootPath string) ([]fs.DirEntry, error) {
	if rootPath == "testPath\\secondLevel" {
		return []fs.DirEntry{stubDir{"thirdLevel.exe", false}}, nil
	}

	return []fs.DirEntry{stubDir{"secondLevel", true}}, nil
}

func TestParseEmailData(t *testing.T) {
	dataMap := parseEmailData([]emailField{messageId}, []byte("Message-ID: 123Id"))
	expectedMessageId := "123Id"
	outputMessageID := string(dataMap[messageId])

	if outputMessageID != expectedMessageId {
		t.Errorf("output %q, expected %q", outputMessageID, expectedMessageId)
	}

}

func TestParseEmailDataWihtBody(t *testing.T) {
	dataMap := parseEmailData([]emailField{messageId, date, body}, []byte(
		"Message-ID: 123Id\r\nDate: Mon, 14 May 2001 16:39:00 -0700 (PDT)\r\n\r\nHere is our forecast"))
	expectedMessageId := "123Id"
	outputMessageID := string(dataMap[messageId])

	if outputMessageID != expectedMessageId {
		t.Errorf("output %q, expected %q", outputMessageID, expectedMessageId)
	}

	expectedBody := "Here is our forecast"
	outputBody := string(dataMap[body])

	if outputBody != expectedBody {
		t.Errorf("output %q, expected %q", outputBody, expectedBody)
	}

}
