package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {

	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepo repository.UserRepository
}

type UserDomainService interface {
	CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception)
	UpdateUser(primitive.ObjectID, model.UserDomainInterface) *exception.Exception
	FindUserByID(id string) (model.UserDomainInterface, *exception.Exception)
	DeleteUser(id string) *exception.Exception
	FindUserByEmail(string) (model.UserDomainInterface, *exception.Exception)
}
