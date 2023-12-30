package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/controller"
	"github.com/odanaraujo/golang/users-api/src/controller/routes"
	"github.com/odanaraujo/golang/users-api/src/model/service"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
	}

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)
	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup, userController)

	if err := r.Run(); err != nil {
		logger.Error("error init server", err)
	}
}
