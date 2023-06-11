package controller

import (
	"crud-api-golang/database"
	"crud-api-golang/internal/constants"
	"crud-api-golang/internal/repository"
	"crud-api-golang/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Create receives a request to create a user
func Create(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := &models.User{}

	err = json.Unmarshal(requestBody, user)
	if err != nil {
		constants.Error(w, http.StatusBadRequest, err)
		return
	}

	err = user.PrepareUser()
	if err != nil {
		constants.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.NewConnection()
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userIDCreated, err := userRepository.Create(user)
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}
	constants.JSON(w, http.StatusCreated, userIDCreated)
}

// GetList receives a request to get all a users
func GetList(w http.ResponseWriter, r *http.Request) {

	requestKeyword := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.NewConnection()
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.GetList(requestKeyword)
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}

	constants.JSON(w, http.StatusOK, users)
	return
}

// GetByID receives a request to get a user
func GetByID(w http.ResponseWriter, r *http.Request) {

}

// Update receives a request to update a user
func Update(w http.ResponseWriter, r *http.Request) {

}

// Delete receives a request to delete a user
func Delete(w http.ResponseWriter, r *http.Request) {

}
