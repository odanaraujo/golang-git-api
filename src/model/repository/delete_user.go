package repository

import (
	"context"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) DeleteUser(id string) *exception.Exception {
	logger.Info("init FindUserByID repository", zap.String("Journey", "FindUserByID"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))
		return exception.InternalServerException(err.Error())
	}

	return nil
}
