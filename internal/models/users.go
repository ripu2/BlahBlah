package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ripu2/blahblah/internal/config/db"
	"github.com/ripu2/blahblah/internal/utils"
)

type User struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserName    string    `json:"user_name" gorm:"unique;not null"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	DOB         time.Time `json:"dob"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"` // Ignore in JSON
	UpdatedAt   time.Time `json:"-" gorm:"autoUpdateTime"` // Ignore in JSON
}

type LoginRequest struct {
	UserName string `binding:"required" gorm:"unique;not null"`
	Password string `binding:"required" gorm:"not null"`
}

func (user *User) CreateUser() (string, error) {
	var exists *bool
	if db.DB == nil {
		return "", errors.New("database connection is nil or uninitialized")
	}
	// Check if the username already exists
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE user_name = $1)", user.UserName).Scan(&exists)
	if err != nil {
		return "", err
	}
	if *exists {
		return "", errors.New("username already exists")
	}

	// Insert new user
	query := `
		INSERT INTO users (user_name, first_name, last_name, phone_number, password, dob, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err = db.DB.QueryRow(query, user.UserName, user.FirstName, user.LastName, user.PhoneNumber, user.Password, user.DOB).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return "", err
	}

	return user.generateAuthToken()

}

func (user *User) generateAuthToken() (string, error) {
	authToken, err := utils.GenerateJWT(int64(user.ID), user.UserName, user.FirstName, user.LastName, user.CreatedAt)
	if err != nil {
		return " ", err
	}

	return authToken, nil
}

func GetUserByUserName(userName string) (User, error) {
	query := `SELECT * FROM users WHERE user_name = $1`
	response := db.DB.QueryRow(query, userName)
	var user User
	err := response.Scan(
		&user.ID,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("User not found")
		}
		return User{}, fmt.Errorf("error fetching user: %v", err)
	}
	return user, nil

}

func (loginRequest *LoginRequest) LoginUser() (string, error) {
	retrievedUser, err := GetUserByUserName(loginRequest.UserName)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(loginRequest.Password, retrievedUser.Password) {
		return "", errors.New("invalid Credentials")
	}
	return retrievedUser.generateAuthToken()

}
