package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "user_id": user_id, "exp": time.Now().Add(time.Hour * 2).Unix()})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	if token == "" {
		return 0, errors.New("invalid token")
	}

	if strings.Contains(token, " ") {
		token = strings.Split(token, " ")[1]
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // syntax for cecking the type of value

		if !ok {
			return nil, errors.New("invalid signin method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	user_id := claims["user_id"].(float64)
	// userId := int64(claims["userId"].(int64))

	return int64(user_id), nil
}
