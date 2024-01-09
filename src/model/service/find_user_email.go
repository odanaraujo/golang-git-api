package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init find user by email service", zap.String("Journey", "FindUserByEmail"))

	userDomain, err := service.userRepo.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return userDomain, nil
}
