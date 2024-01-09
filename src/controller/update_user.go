package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/controller/model/request"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(ctx *gin.Context) {
	logger.Info("init update user", zap.String("Journey", "UpdateUser"))

	var userRequest request.UserUpdateRequest
	id := ctx.Param("id")

	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("error trying validate to id", err, zap.String(
			"Journey", "UpdateUser"))
		errMessage := exception.InternalServerException("user id is not a valid id")
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying validate to user", err, zap.String(
			"Journey", "UpdateUser"))
		errMessage := exception.InternalServerException(err.Error())
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	domain := model.NewUSerUpdateDomain(userRequest.Name, userRequest.Age)

	if err := uc.service.UpdateUser(id, domain); err != nil {
		logger.Error("error when update for user in the database", err, zap.String(
			"Journey", "UpdateUser"))

		ctx.JSON(err.Code, err)
		return
	}

	ctx.Status(http.StatusOK)

}
