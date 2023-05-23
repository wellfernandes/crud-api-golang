package constants

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a json response to request
func JSON(w http.ResponseWriter, statusCode int, dataJson interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(dataJson)
	if err != nil {
		log.Fatal(err)
	}
}

// Error returns a json error
func Error(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}
