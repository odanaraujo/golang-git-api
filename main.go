package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/golang/users-api/src/configuration/database/mongodb"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/controller/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
		return
	}

	mongoDBConnection, err := mongodb.NewMongoDBConnection()

	if err != nil {
		logger.Error("Unable to connect to database", err)
		return
	}

	userController := initDependencies(mongoDBConnection)

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup, userController)

	if err := r.Run(); err != nil {
		logger.Error("error init server", err)
		return
	}
}
