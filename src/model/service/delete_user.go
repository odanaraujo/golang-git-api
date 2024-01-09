package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"go.uber.org/zap"
)

func (service *userDomainService) DeleteUser(id string) *exception.Exception {

	logger.Info("init get service", zap.String("Journey", "DeleteUser"))

	if _, err := service.FindUserByID(id); err != nil {
		logger.Error("user id not found", err,
			zap.String("Journey", "DeleteUser"))
		return err
	}

	if err := service.userRepo.DeleteUser(id); err != nil {
		logger.Error("error delete id", err, zap.String("Journey", "DeleteUser"))
		return err
	}

	return nil
}
