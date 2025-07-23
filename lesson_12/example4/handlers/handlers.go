package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/users", GetUsers)
}

// GetUsers returns users info.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	u := []struct {
		Name  string
		Email string
	}{
		{
			Name:  "x1",
			Email: "x1@gmail.com",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&u)
}
