package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/controller"
)

// initialize the routes
func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {

	r.GET("/user/:id", controller.FindUserByID)
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}
