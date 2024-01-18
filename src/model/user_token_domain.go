package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"os"
	"strings"
	"time"
)

const (
	jwtUserKey   = "JWT_USER_KEY"
	authorizatio = "Authorization"
)

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

func VerifyTokenMiddleware(c *gin.Context) {

	secret := os.Getenv(jwtUserKey)
	tokenValue := removeBearePrefix(c.Request.Header.Get(authorizatio))

	token, err := jwt.Parse(removeBearePrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, exception.BadRequestException("invalid token")
	})

	if err != nil {
		excp := exception.UnauthorizedRequestException("invalid token")
		c.JSON(excp.Code, excp)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		excp := exception.UnauthorizedRequestException("invalid token")
		c.JSON(excp.Code, excp)
		c.Abort()
		return
	}

	userDomain := userDomain{
		id:    claims["id"].(string),
		name:  claims["name"].(string),
		email: claims["email"].(string),
		age:   uint8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("user authorization: %#v", userDomain))
}

func removeBearePrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
