package repository

import (
	"context"
	"errors"
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

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *exception.Exception) {
	logger.Info("init FindUserByEmail repository", zap.String("Journey", "FindUserByEmail"))
	ctx := context.Background()
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(ctx, filter).Decode(userEntity)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errMessage := fmt.Sprintf("User not found with email %s", email)
			return nil, exception.NotFoundException(errMessage)
		}

		errMessage := fmt.Sprintf("Error trying to find user by email %s", email)
		return nil, exception.InternalServerException(errMessage)
	}

	return converter.ConverterEntitytoDomain(*userEntity), nil
}
