package services

import (
	"errors"
	"time"

	"github.com/ripu2/blahblah/internal/models"
	"github.com/ripu2/blahblah/internal/utils"
)

func CreateUserService(user *models.User) (string, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", errors.New(err.Error())
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = hashedPassword
	token, err := user.CreateUser()
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil
}

func LoginUserService(loginRequest *models.LoginRequest) (string, error) {
	token, err := loginRequest.LoginUser()
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil
}
