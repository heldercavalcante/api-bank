package tools

import (
	"encoding/json"
	"net/http"
)

func WriteResponseJSON(w http.ResponseWriter, status int, r any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(r)
}