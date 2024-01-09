package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) UpdateUser(id string, user model.UserDomainInterface) *exception.Exception {
	logger.Info("init update user service", zap.String("journey", "UpdateUser"))

	user.EncryptPassword()

	if _, err := service.FindUserByID(id); err != nil {
		logger.Error("user id not found", err,
			zap.String("Journey", "UpdateUser"))
		return err
	}

	err := service.userRepo.UpdateUser(id, user)

	if err != nil {
		logger.Error("Unable to save user", err,
			zap.String("Journey", "UpdateUser"))
		return exception.InternalServerException(err.Error())
	}

	return nil
}
