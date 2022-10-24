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
	fmt.Println("Reading for: ", config.ChunkReader)

	fmt.Println("Reading for: ", rootFilePath)
	fileSource := infraestructure.NewEmailFileSource(rootFilePath)
	emailService := infraestructure.EmailIteratorService{Source: fileSource, Chunk: config.ChunkReader}

	emailZincService := infraestructure.EmailZincService{
		Host:         config.HostZync,
		User:         config.UserZync,
		Pass:         config.PassZync,
		DefaultIndex: config.DefaultIndex,
	}

	indexMailApp := application.IndexMailApp{EmailService: &emailService, EmailIndexer: emailZincService}

	indexMailApp.IndexEmails()
}
