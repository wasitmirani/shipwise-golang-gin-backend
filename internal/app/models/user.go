package models

import (
	"time"
	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	Name              string    `json:"name" gorm:"not null"`
	Email             string    `json:"email" gorm:"unique;not null"`
	EmailVerifiedAt   time.Time `json:"email_verified_at" gorm:"default:null"`
	Password          string    `json:"-" gorm:"not null"` // Password is hidden in JSON responses
	RememberToken     string    `json:"-"`
}