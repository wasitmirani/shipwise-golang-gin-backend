package auth

import (
	"encoding/json"
	"net/http"
)

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Here you would typically add code to save the user to a database
	// For simplicity, we'll just return a success message

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Here you would typically add code to authenticate the user
	// For simplicity, we'll just return a success message

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged in successfully"})
}