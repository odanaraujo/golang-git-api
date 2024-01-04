package repository

import (
	"context"
	"fmt"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *exception.Exception) {

	logger.Info("init FindUserByID repository", zap.String("Journey", "FindUserByID"))

	ctx := context.Background()
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf("user not found with this id %s", id)
			return nil, exception.NotFoundException(errMessage)
		}
		errMessage := fmt.Sprintf("error try to find user by id")
		return nil, exception.InternalServerException(errMessage)
	}

	return converter.ConverterEntitytoDomain(*userEntity), nil
}
