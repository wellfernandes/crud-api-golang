package controller

import "net/http"

// Create a new user
func Create(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create a new user"))
}

// GetList get all users
func GetList(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get all users"))
}

// GetByID get a user
func GetByID(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get a user"))
}

// Update update a user
func Update(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Update a user"))
}

// Delete update a user
func Delete(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Delete a user"))
}
