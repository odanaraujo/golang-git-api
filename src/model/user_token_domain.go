package model

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"os"
	"strings"
	"time"
)

const jwtUserKey = "JWT_USER_KEY"

func (ud *userDomain) GenerateToken() (string, *exception.Exception) {

	secret := os.Getenv(jwtUserKey)
	claims := jwt.MapClaims{
		"id":    ud.id,
		"name":  ud.name,
		"email": ud.email,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", exception.InternalServerException("" +
			fmt.Sprintf("error trying to generate jwt token, err = %s", err.Error()))
	}

	tokenStringBearer := removeBearePrefix(strings.TrimSpace(tokenString))

	return tokenStringBearer, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *exception.Exception) {
	secret := os.Getenv(jwtUserKey)
	token, err := jwt.Parse(removeBearePrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, exception.BadRequestException("invalid token")
	})

	if err != nil {
		return nil, exception.UnauthorizedRequestException("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, exception.UnauthorizedRequestException("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		name:  claims["name"].(string),
		email: claims["email"].(string),
		age:   uint8(claims["age"].(float64)),
	}, nil

}

func removeBearePrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
