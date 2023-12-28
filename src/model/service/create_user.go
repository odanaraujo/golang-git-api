package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.uber.org/zap"
)

func (service *userDomainService) CreateUser(user model.UserDomainInterface) *exception.Exception {
	logger.Info("init create user service", zap.String("journey", "CreateUser"))

	user.EncryptPassword()

	return nil
}
