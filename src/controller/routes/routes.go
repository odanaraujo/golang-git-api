package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/controller"
	"github.com/odanaraujo/golang/users-api/src/model"
)

// initialize the routes
func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {

	r.GET("/user/:id", model.VerifyTokenMiddleware, controller.FindUserByID)
	r.GET("/getUserByEmail/:email", model.VerifyTokenMiddleware, controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	r.POST("/login", controller.LoginUser)
}
