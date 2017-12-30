package services

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
