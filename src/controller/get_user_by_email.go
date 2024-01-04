package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/view"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByEmail(ctx *gin.Context) {
	logger.Info("init get user email", zap.String("Journey", "FindUserByEmail"))

	email := ctx.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String(
			"Journey", "FindUserByEmail"))
		errMessage := exception.BadRequestException("Email is not a valid email")
		ctx.JSON(http.StatusBadRequest, errMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)

	if err != nil {
		logger.Error("error get user email", err, zap.String("Journey", "FindUserByEmail"))
		err := exception.BadRequestException(err.Error())
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConverterDomainToResponse(userDomain))
}
