package main

import (
	// "encoding/json"
	// "fmt"
	"example/indexer/application"
	"example/indexer/config"
	"example/indexer/infraestructure"
	"example/indexer/interface/restHttp"
)

func main() {

	httpZinc := infraestructure.NewHttpClient(config.HostZync, config.UserZync, config.PassZync, 5)

	emailZincService := infraestructure.EmailZincService{DefaultIndex: "indexEmail", HttpClient: httpZinc}

	indexMailApp := application.IndexMailApp{EmailIndexer: emailZincService}

	handler := restHttp.HandlerEmails{IndexMailApp: indexMailApp}

	restHttp.InitServer(":3000", handler)

}
