package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(user model.UserDomainInterface) *exception.Exception
	UpdateUser(string, model.UserDomainInterface) *exception.Exception
	FindUserByID(string) (*model.UserDomainInterface, *exception.Exception)
	FindUserByEmail(string) (*model.UserDomainInterface, *exception.Exception)
	DeleteUser(string) *exception.Exception
}
