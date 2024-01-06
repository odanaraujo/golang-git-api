package repository

import (
	"context"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) UpdateUser(id primitive.ObjectID, userDomain model.UserDomainInterface) *exception.Exception {
	logger.Info("init UpdateUser repository", zap.String("Journey", "UpdateUser"))

	ctx := context.Background()
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConverterDomainToEntity(userDomain)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return exception.InternalServerException(err.Error())
	}

	return nil
}
