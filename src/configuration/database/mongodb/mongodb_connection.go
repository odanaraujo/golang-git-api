package mongodb

import (
	"context"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
)

const (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	mongoUri := os.Getenv(MONGODB_URL)
	mongoDBDatabase := os.Getenv(MONGODB_USER_DB)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		logger.Error("Unable to connect the database", err, zap.String(
			"Journey", "NewMongoDBConnection"))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Unable to ping the database", err, zap.String(
			"Journey", "NewMongoDBConnection"))
	}

	return client.Database(mongoDBDatabase), nil
}
