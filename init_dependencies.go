package main

import (
	"github.com/odanaraujo/golang/users-api/src/controller"
	"github.com/odanaraujo/golang/users-api/src/model/repository"
	"github.com/odanaraujo/golang/users-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(mongoDBConnection *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(mongoDBConnection)
	userService := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(userService)
	return userController
}
