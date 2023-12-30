package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/configuration/validation"
	"github.com/odanaraujo/golang/users-api/src/controller/model/request"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/service"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUSerDomain(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Age)

	userService := service.NewUserDomainService()

	if err := userService.CreateUser(domain); err != nil {
		logger.Error("error trying create user domain", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.String(http.StatusOK, "")

}
