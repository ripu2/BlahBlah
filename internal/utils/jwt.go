package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Load secret from env
func GenerateJWT(userId int64, userName, firstName, lastName string, createdAt time.Time) (string, error) {
	jwtExpires, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY_HOURS"), 10, 64)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     userId,
		"username":   userName,
		"first_name": firstName,
		"last_name":  lastName,
		"createdAt":  createdAt.Format(time.RFC3339),
		"exp":        time.Now().Add(time.Hour * time.Duration(jwtExpires)).Unix(),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return signedToken, err
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token signature")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, errors.New(err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	return int64(claims["id"].(float64)), nil
}

// var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Load secret from env

// func GenerateJWT(userId int64, userName, firstName, lastName string, createdAt time.Time) (string, error) {
// 	jwtExpires, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY_HOURS"), 10, 64)

// 	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// "userId":     userId,
// "username":   userName,
// "first_name": firstName,
// "last_name":  lastName,
// "createdAt":  createdAt.Format(time.RFC3339),
// "exp":        time.Now().Add(time.Hour * time.Duration(jwtExpires)).Unix(),
// 	})

// 	signedToken, err := authToken.SignedString([]byte(jwtSecret))
// 	if err != nil {
// 		return "", errors.New(err.Error())
// 	}
// 	return signedToken, nil
// }

// func VerifyToken(token string) (int64, error) {
// 	fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
// 	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)

// 		if !ok {
// 			return nil, errors.New("invalid token signature")
// 		}
// 		return []byte(jwtSecret), nil
// 	})

// 	if err != nil {
// 		return 0, errors.New(err.Error())
// 	}

// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)

// 	if !ok || !parsedToken.Valid {
// 		return 0, errors.New("invalid token")
// 	}
// 	return int64(claims["userId"].(float64)), nil
// }
