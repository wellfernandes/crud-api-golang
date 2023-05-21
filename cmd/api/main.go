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

	//err := configs.Load()
	//if err != nil {
	//	panic(err)
	//}
	//
	//r := chi.NewRouter()
	//r.Post("/", handlers.Create)
	//r.Put("/{id}", handlers.Update)
	//r.Delete("/{id}", handlers.Delete)
	//r.Get("/", handlers.List)
	//r.Get("/{id}", handlers.Get)
	//
	//log.Println("running api...")
	//
	//http.ListenAndServe(fmt.Sprintf(": %s", configs.GetServerPort()), r)
}
