package services

import (
	"errors"
	"shipwise/internal/app/models"
	"shipwise/internal/app/repositories"
	"shipwise/internal/pkg/utils"
)

func Login(email, password string) (*models.User, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
