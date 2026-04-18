package utils

import (
	"encoding/json"
	"net/http"
)

var Client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, message string, statusCode int) {
	WriteJSON(w, statusCode, ErrorResponse{Error: message})
}
