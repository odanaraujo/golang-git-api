package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception) {

	logger.Info("init login userDomain service", zap.String("journey", "LoginUserService"))

	userDomain.EncryptPassword()

	user, err := service.findUserByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, exception.InternalServerException(err.Error())
	}
	return user, nil
}
