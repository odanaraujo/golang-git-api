package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/view"
	"go.uber.org/zap"
	"net/http"
)

// get user by id
func (uc *userControllerInterface) FindUserByID(ctx *gin.Context) {
	logger.Info("init get user", zap.String("Journey", "FindUserByID"))

	id := ctx.Param("id")

	userDomain, err := uc.service.FindUserByID(id)

	if err != nil {
		logger.Error("Error when searching for user", err, zap.String(
			"Journey", "FindUserByID"))
		excp := exception.InternalServerException(err.Error())
		ctx.JSON(http.StatusInternalServerError, excp)
		return
	}

	ctx.JSON(http.StatusOK, view.ConverterDomainToResponse(userDomain))

}
