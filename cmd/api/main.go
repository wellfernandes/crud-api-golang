package main

import (
	"crud-api-golang/configs"
	"crud-api-golang/internal/constants"
	"crud-api-golang/rest"
	"fmt"
	"log"
	"net/http"
)

// main function for initializing the api
func main() {

	err := configs.Load()
	if err != nil {
		log.Println(constants.GENERIC_ERROR, err.Error())
	}

	port, host, nameDatabase, userDatabase, strConn, err := configs.GetConfigInfo()
	if err != nil {
		log.Fatal(constants.SETTING_PARAMETERS_ERROR, err.Error())
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
		log.Fatal(constants.GENERIC_ERROR, err.Error())
	}

	log.Println("\n\nAPI is running...")
}
