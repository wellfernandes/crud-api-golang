package handlers

import (
	"crud-api-golang/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id")) //convert string to integer
	if err != nil {
		log.Printf("Error when trying to parse the id %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error while decoding json file! %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error updating record! %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error! %d records were updated.", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Successfully updated!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
