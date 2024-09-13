package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// GenerateToken generates a JWT token with email and user ID claims.
func GenerateToken(email string, userId int64) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // Set expiration time
	}

	// Create the token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the token
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid{
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok{
	// 	return errors.New("invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
