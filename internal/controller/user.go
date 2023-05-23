package controller

import (
	"crud-api-golang/database"
	"crud-api-golang/internal/repository"
	"crud-api-golang/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Create receives a request to create a user
func Create(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err) // test
	}

	user := models.User{}

	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUserRepository(db)
	userIDCreated, err := repository.Create(user)
	if err != nil {
		log.Fatal(err) // test
	}

	w.Write([]byte(fmt.Sprintf("user created: %d", userIDCreated)))
}

// GetList receives a request to get all a users
func GetList(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get all users"))
}

// GetByID receives a request to get a user
func GetByID(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get a user"))
}

// Update receives a request to update a user
func Update(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Update a user"))
}

// Delete receives a request to delete a user
func Delete(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Delete a user"))
}
