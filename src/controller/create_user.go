package controller

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/controller/model/response"

	"github.com/odanaraujo/golang/users-api/src/configuration/validation"
	"github.com/odanaraujo/golang/users-api/src/controller/model/request"
)

func CreateUser(ctx *gin.Context) {

	logger.Info("Init CreateUser Controller")

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err)
		excp := validation.ValidateUserRequest(err)
		ctx.JSON(excp.Code, excp)
		return
	}

	userResponse := response.UserResponse{
		ID:    strconv.FormatInt(int64(rand.Intn(10)+1), 10),
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	ctx.JSON(http.StatusOK, userResponse)
}
