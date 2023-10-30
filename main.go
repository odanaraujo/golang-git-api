package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"github.com/odanaraujo/golang/users-api/src/controller/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
	}

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup)

	if err := r.Run(); err != nil {
		logger.Error("error init server", err)
	}
}
