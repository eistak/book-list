package tools

import (
	"books-list/models"
	"encoding/json"
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
