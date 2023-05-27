package main

import (
	"crud-api-golang/configs"
	"crud-api-golang/rest"
	"fmt"
	"log"
	"net/http"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Printf("\nerror:", err.Error())
	}

	port, host, nameDatabase, userDatabase, strConn, err := configs.GetConfigInfo()
	if err != nil {
		log.Fatal("\n\nerror setting internal parameters: ", err.Error())
	}

	log.Println(
		"\n\nPort:", port,
		"\nHost:", host,
		"\nDatabase name:", nameDatabase,
		"\nDatabase user:", userDatabase,
		"\nString connection nameDatabase:", strConn,
	)

	r := rest.NewServerRouter()

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("\n\nerror: ", err.Error())
	}

	log.Println("\n\nAPI is running...")
}
