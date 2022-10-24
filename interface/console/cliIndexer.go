package main

import (
	"fmt"
	"os"

	"example/indexer/application"
	"example/indexer/config"
	"example/indexer/infraestructure"
)

func main() {
	rootFilePath := os.Args[1]
	// rootFilePath = "C:\\enron_mail_20110402\\maildir"

	fmt.Println("Reading for: ", rootFilePath)
	fileSource := infraestructure.NewEmailFileSource(rootFilePath)
	emailService := infraestructure.EmailIteratorService{Source: fileSource, Chunk: config.ChunkReader}

	httpZinc := infraestructure.NewHttpClient(config.HostZync, config.UserZync, config.PassZync, 5)

	emailZincService := infraestructure.EmailZincService{DefaultIndex: "indexEmail", HttpClient: httpZinc}

	indexMailApp := application.IndexMailApp{EmailService: &emailService, EmailIndexer: emailZincService}

	indexMailApp.IndexEmails()
}
