package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userID int64) (string, int64, error) {
	expiration_time := time.Now().Add(time.Hour *2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email": email,
		"exp": expiration_time,
	})

	tokenString, err := token.SignedString(([]byte(secretKey)))

	return tokenString, expiration_time, err
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, nil
	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		return 0, errors.New("invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid Claims")
	}

	userID := int64(claims["userID"].(float64))


	return userID, nil

}