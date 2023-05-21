package main

import (
	"crud-api-golang/configs"
	"crud-api-golang/rest"
	"log"
	"net/http"
)

func main() {

	configs.Load()

	port, host, database, err := configs.GetConfigInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"\n\nAPI Port:", port,
		"\nHost:", host,
		"\nDatabase:", database,
	)

	r := rest.NewServerRouter()

	log.Println("running api...")
	log.Fatal(http.ListenAndServe(":5000", r))
}
