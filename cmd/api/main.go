package main

import (
	"crud-api-golang/configs"
	"crud-api-golang/rest"
	"fmt"
	"log"
	"net/http"
)

func main() {

	configs.Load()

	port, host, nameDatabase, userDatabase, strConn, err := configs.GetConfigInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"\n\nPort:", port,
		"\nHost:", host,
		"\nDatabase name:", nameDatabase,
		"\nDatabase user:", userDatabase,
		"\nString connection nameDatabase:", strConn,
	)

	r := rest.NewServerRouter()

	log.Println("\n\nAPI is running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
