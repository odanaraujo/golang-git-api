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

func (uc *userControllerInterface) LoginUser(ctx *gin.Context) {
	logger.Info("init LoginUserService controller", zap.String("Journey", "LoginUserService"))

	var loginRequest request.UserLogin

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		logger.Error("error trying to validate user info", err,
			zap.String("Journey", "LoginUserService"))

		errRest := validation.ValidateUserRequest(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(loginRequest.Email, loginRequest.Password)

	userModel, err := uc.service.LoginUserService(domain)

	if err != nil {
		logger.Error("error trying validate email and password", err, zap.String(
			"Journey", "LoginUser"))

		ctx.JSON(err.Code, err)
		return
	}

	token, err := userModel.GenerateToken()

	if err != nil {
		logger.Error("error trying generate jwt token", err, zap.String(
			"Journey", "LoginUser"))

		ctx.JSON(err.Code, err)
		return
	}

	ctx.Header("Authorization", token)

	ctx.JSON(http.StatusOK, view.ConverterDomainToResponse(userModel))

}
