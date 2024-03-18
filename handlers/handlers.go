package handlers

import (
	"encoding/json"
	"net/http"
)

// JSON response struct
type Response struct {
	Message string `json:"message"`
}

// HomeHandler responds to the "/" route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to the terraform golang api!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
