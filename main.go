package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/golang/users-api/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
