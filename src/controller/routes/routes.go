package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/controller"
)

// initialize the routes
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", controller.GetUserByID)
	r.GET("/getUserByEmail/:email", controller.GetUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}
