package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init get service", zap.String("Journey", "FindUserByID"))

	userDomain, err := service.userRepo.FindUserByID(id)

	if err != nil {
		logger.Error("error the find user id", err, zap.String("Journey", "FindUserByID"))
		return nil, err
	}

	return userDomain, nil
}
