package controller

import (
	"crud-api-golang/database"
	"crud-api-golang/internal/constants"
	"crud-api-golang/internal/repository"
	"crud-api-golang/models"
	"encoding/json"
	"errors"
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

// GetByID receives a request to get a user by ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		constants.Error(w, http.StatusBadRequest, errors.New(constants.ID_NOT_PROVIDED))
		return
	}

	db, err := database.NewConnection()
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetByID(id)
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		constants.Error(w, http.StatusNotFound, errors.New("Usuário não encontrado"))
		return
	}

	constants.JSON(w, http.StatusOK, user)
}

// Update receives a request to update a user
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		constants.Error(w, http.StatusBadRequest, errors.New(constants.ID_NOT_PROVIDED))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userToUpdate := &models.User{}
	err = json.Unmarshal(requestBody, userToUpdate)
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

	_, err = userRepository.GetByID(id)
	if err != nil {
		if err.Error() == constants.USER_NOT_FOUND {
			constants.Error(w, http.StatusNotFound, err)
		} else {
			constants.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	err = userRepository.Update(id, userToUpdate)
	if err != nil {
		constants.Error(w, http.StatusInternalServerError, err)
		return
	}

	constants.JSON(w, http.StatusOK, constants.USER_UPDATED)
}

// Delete receives a request to delete a user
func Delete(w http.ResponseWriter, r *http.Request) {

}
