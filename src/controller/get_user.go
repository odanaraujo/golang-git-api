package controller

import "github.com/gin-gonic/gin"

// get user by id
func (uc *userControllerInterface) FindUserByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "getUserByID",
	})
}
