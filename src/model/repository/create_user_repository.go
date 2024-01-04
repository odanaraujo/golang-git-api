package repository

import (
	"context"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) CreateUser(userdomain model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init createUser repository", zap.String("Journey", "CreateUser"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConverterDomainToEntity(userdomain)
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, exception.InternalServerException(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConverterEntitytoDomain(*value), nil
}
