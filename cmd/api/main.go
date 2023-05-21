package main

import (
	"crud-api-golang/rest"
	"log"
	"net/http"
)

func main() {

	r := rest.NewServerRouter()

	log.Println("running api...")
	log.Fatal(http.ListenAndServe(":5000", r))
}
