package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

// User represents a user in the system
type User struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

// ValidationError represents a structured validation error response
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Mock database of users
var users = map[string]string{
	"admin": "password123",
	"user":  "password456",
}

// JWT secret key (should be stored securely in a real application)
var jwtSecret = []byte("your-256-bit-secret")

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate username
	if len(user.Username) < 3 {
		validationError := ValidationError{
			Field:   "username",
			Message: "Username must be at least 3 characters long",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationError)
		return
	}

	// Validate password
	if len(user.Password) < 6 {
		validationError := ValidationError{
			Field:   "password",
			Message: "Password must be at least 6 characters long",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationError)
		return
	}

	// Here you would typically add code to save the user to a database
	// For simplicity, we'll just return a success message

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login handles the login request
func Login(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		// Handle validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []ValidationError
			for _, e := range validationErrors {
				var message string
				switch e.Tag() {
				case "required":
					message = e.Field() + " is required"

				case "min":
					if e.Field() == "Username" {
						message = "Username must be at least 3 characters long"
					} else if e.Field() == "Password" {
						message = "Password must be at least 6 characters long"
					}
				}
				errors = append(errors, ValidationError{
					Field:   e.Field(),
					Message: message,
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the user exists and the password is correct
	if storedPassword, exists := users[user.Username]; exists {
		if storedPassword == user.Password {
			// Generate a JWT token
			token, err := generateToken(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			// Return the token and a success message
			c.JSON(http.StatusOK, gin.H{
				"message": "User logged in successfully",
				"token":   token,
			})
			return
		}
	}

	// If the user doesn't exist or the password is wrong
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// generateToken creates a JWT token for the given username
func generateToken(username string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
