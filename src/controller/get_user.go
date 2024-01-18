package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

const authorization = "Authorization"

// get user by id
func (uc *userControllerInterface) FindUserByID(ctx *gin.Context) {
	logger.Info("init get user", zap.String("Journey", "FindUserByID"))

	id := ctx.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Error trying to validate id", err, zap.String(
			"Journey", "FindUserByID"))
		errMessage := exception.BadRequestException("user id is not a valid id")
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(id)

	if err != nil {
		logger.Error("Error when searching for user", err, zap.String(
			"Journey", "FindUserByID"))

		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConverterDomainToResponse(userDomain))

}
