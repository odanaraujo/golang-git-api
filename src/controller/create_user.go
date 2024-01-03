package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/configuration/validation"
	"github.com/odanaraujo/golang/users-api/src/controller/model/request"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/view"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(ctx *gin.Context) {

	logger.Info("Init CreateUser Controller")

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err)
		excp := validation.ValidateUserRequest(err)
		ctx.JSON(excp.Code, excp)
		return
	}

	domain := model.NewUSerDomain(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		logger.Error("error trying create user domain", err)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully",
		zap.String("journey", "createduser"))

	ctx.JSON(http.StatusOK, view.ConverterDomainToResponse(domainResult))

}
