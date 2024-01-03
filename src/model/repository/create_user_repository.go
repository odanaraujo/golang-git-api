package repository

import (
	"context"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

type params struct {
	name     string
	email    string
	password string
	age      uint8
}

func (ur *userRepository) CreateUser(userdomain model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init createUser repository", zap.String("Journey", "CreateUser"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	println(userdomain.GetName())
	result, err := collection.InsertOne(context.Background(),
		params{
			name:     userdomain.GetName(),
			email:    userdomain.GetEmail(),
			password: userdomain.GetPassword(),
			age:      userdomain.GetAge(),
		})

	if err != nil {
		return nil, exception.InternalServerException(err.Error())
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		strID := id.Hex()
		userdomain.SetID(strID)
	}

	return userdomain, nil
}
