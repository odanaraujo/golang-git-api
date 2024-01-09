package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init create user service", zap.String("journey", "CreateUser"))

	user.EncryptPassword()

	userDomainRepository, err := service.userRepo.CreateUser(user)

	userBD, _ := service.FindUserByEmail(user.GetEmail())

	if userBD != nil {
		return nil, exception.BadRequestException("email is already registered in another account")
	}

	if err != nil {
		logger.Error("Unable to save user", err,
			zap.String("Journey", "CreateUserService"))
		return nil, exception.InternalServerException(err.Error())
	}

	return userDomainRepository, nil
}
