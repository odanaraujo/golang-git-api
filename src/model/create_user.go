package model

import (
	"fmt"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"go.uber.org/zap"
)

func (domain *UserDomain) CreateUser() *exception.Exception {

	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	domain.EncryptPassword()

	fmt.Println(domain.Password)
	return nil
}
