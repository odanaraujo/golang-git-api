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

func (service *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init find user by email service", zap.String("Journey", "FindUserByEmail"))

	userDomain, err := service.userRepo.FindUserByEmail(email)

	if err != nil {
		return nil, exception.InternalServerException(err.Error())
	}
	return userDomain, nil
}

func (service *userDomainService) findUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init find user by email and password service", zap.String("Journey", "LoginUserService"))

	userDomain, err := service.userRepo.FindUserByEmailAndPassword(email, password)

	if err != nil {
		return nil, err
	}
	return userDomain, nil
}
