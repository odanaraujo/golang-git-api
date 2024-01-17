package repository

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception)
	FindUserByEmail(id string) (model.UserDomainInterface, *exception.Exception)
	FindUserByID(id string) (model.UserDomainInterface, *exception.Exception)
	UpdateUser(id string, userModel model.UserDomainInterface) *exception.Exception
	DeleteUser(id string) *exception.Exception
	FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *exception.Exception)
}
